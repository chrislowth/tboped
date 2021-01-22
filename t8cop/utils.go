package t8cop

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

var crd Crd
var requirements Requirements

var resourcesObject = CrdProperty{
	Description: "(resources object)",
	Type:        "object",
	Properties: map[string]CrdProperty{
		"limits": {
			Type: "object",
			Properties: map[string]CrdProperty{
				"memory": {
					Type: "string",
				},
				"cpu": {
					Type: "string",
				},
			},
		},
		"requests": {
			Type: "object",
			Properties: map[string]CrdProperty{
				"memory": {
					Type: "string",
				},
				"cpu": {
					Type: "string",
				},
			},
		},
	},
}

var envObject = CrdProperty{
	Description: "(env object)",
	Type:        "object",
}

func yamlField(y interface{}, keys ...string) interface{} {
	rtn := y
	for _, key := range keys {
		switch r := rtn.(type) {
		case map[interface{}]interface{}:
			rtn = r[key]
		case map[string]interface{}:
			rtn = r[key]
		default:
			fmt.Fprintf(os.Stderr, "Cant find branch %s\n", strings.Join(keys, "."))
			return nil
		}
	}
	return rtn
}

// Find the names of feature branches missing from the CRD structure but used as conditions in the REQUIREMENTS
func findMissingConditions() []string {
	missing := make(map[string]bool)

	for _, r := range requirements.Dependencies {
		for _, conf := range strings.Split(r.Condition, ",") {
			ok := true
			condParts := strings.Split(conf, ".")
			prop := crd.Spec.Validation.OpenAPIV3Schema.Properties["spec"]
			for i, p := range condParts {
				prop, ok = prop.Properties[p]
				if !ok {
					if i == 0 && len(condParts) == 2 && condParts[1] == "enabled" {
						missing[condParts[0]] = true
					}
				}
			}
		}
	}

	rtn := make([]string, 0, 8)
	for k, _ := range missing {
		rtn = append(rtn, k)
	}
	sort.Strings(rtn)

	return rtn
}

// Fix missing conditions in the CRD, assuming that the feature they refer to has just one property: a boolean
// called "enabled".
func fixMissingConditions() {
	for _, cond := range findMissingConditions() {
		crd.Spec.Validation.OpenAPIV3Schema.Properties["spec"].Properties[cond] = CrdProperty{
			Description: "(missing description)",
			Type:        "object",
			Properties: map[string]CrdProperty{
				"enabled": {
					Description: "Enable " + cond + " feature",
					Type:        "boolean",
				}},
		}
	}
}

func fixMissingResourceLimits(dirName string) {
	type ValuesYaml struct {
		Resources *struct {
			Limits *struct {
				Memory *string
			}
			Requests *struct {
				Memory *string
			}
		}
	}

	for _, dir := range []string{"base", "probes", "services"} {
		files, err := ioutil.ReadDir(dirName + "/helm-charts/" + dir)
		if err != nil {
			panic(err)
		}
		for _, fileInfo := range files {
			values := ValuesYaml{}

			fileName := dirName + "/helm-charts/" + dir + "/" + fileInfo.Name() + "/values.yaml"

			b, err := ioutil.ReadFile(fileName)
			if err != nil {
				panic(err)
			}

			err = yaml.Unmarshal(b, &values)
			if err != nil {
				panic(err)
			}

			if values.Resources == nil {
				continue
			}

			spec := crd.Spec.Validation.OpenAPIV3Schema.Properties["spec"]
			prop1, ok1 := spec.Properties[fileInfo.Name()]
			if ok1 {
				_, ok2 := prop1.Properties["resources"]
				if !ok2 {
					prop1.Properties["resources"] = *resourcesObject.Clone()
				}
			}
		}
	}
}

func patchInto(crdProp CrdProperty, propName string, yamlProp interface{}) {
	if yamlProp == nil {
		return
	}
	if crdProp.Properties == nil {
		crdProp.Properties = make(map[string]CrdProperty)
	}
	switch y := yamlProp.(type) {
	case map[string]interface{}:
		if crdProp.Properties[propName].Type == "" {
			crdProp.Properties[propName] = CrdProperty{
				// Description: "MISSING!",
				Type:       "object",
				Properties: make(map[string]CrdProperty),
			}
		}
		for k, v := range y {
			patchInto(crdProp.Properties[propName], k, v)
			// patchInto(crdProp, k.(string), v)
		}

	case map[interface{}]interface{}:
		if crdProp.Properties[propName].Type == "" {
			crdProp.Properties[propName] = CrdProperty{
				// Description: "MISSING!",
				Type:       "object",
				Properties: make(map[string]CrdProperty),
			}
		}
		for k, v := range y {
			patchInto(crdProp.Properties[propName], k.(string), v)
			// patchInto(crdProp, k.(string), v)
		}
	case []interface{}:
		// TO DO
	case string:
		crdProp.Properties[propName] = CrdProperty{
			Type: "string",
		}
	case bool:
		crdProp.Properties[propName] = CrdProperty{
			Type: "boolean",
		}
	case int:
		crdProp.Properties[propName] = CrdProperty{
			Type: "integer",
		}
	case float64:
		// TO DO
	default:
		panic(fmt.Sprintf("Bad yamlProp type: %T", yamlProp))
	}
}

// Process the "values.yaml" helm-chart files and patch any values found there into our in-memory
// CRD manifest
func fixRemainingValues(dirName string) {
	for _, dir := range []string{"base", "probes", "services"} {
		files, err := ioutil.ReadDir(dirName + "/helm-charts/" + dir)
		if err != nil {
			panic(err)
		}
		for _, fileInfo := range files {
			var values interface{}

			fileName := dirName + "/helm-charts/" + dir + "/" + fileInfo.Name() + "/values.yaml"

			b, err := ioutil.ReadFile(fileName)
			if err != nil {
				panic(err)
			}

			err = yaml.Unmarshal(b, &values)
			if err != nil {
				panic(err)
			}

			patchInto(crd.Spec.Validation.OpenAPIV3Schema.Properties["spec"], fileInfo.Name(), values)
		}
	}
}

// Update the in-memory CRD manifest to have a genericObject at the specified location
func setGenericObject(path ...string) {
	obj := crd.Spec.Validation.OpenAPIV3Schema

	for _, step := range path[0 : len(path)-1] {
		obj = obj.Properties[step]
	}

	obj.Properties[path[len(path)-1]] = CrdProperty{
		Description: "(forced generic object)",
		Type:        "object",
	}
}

var valueVarRe = regexp.MustCompile(`{{-?\s*(.*?)\s*-?}}`)
var re1 = regexp.MustCompile(`\.Values\.[a-zA-Z0-9_\.]+`)
var re2 = regexp.MustCompile(`^\s*(with|range)\s+(.*?:=\s*)?(\.Values\.[a-zA-Z0-9_\.]+)\s*$`)
var re3 = regexp.MustCompile(`^\s*toYaml\s+(\.Values\.[a-zA-Z0-9_\.]+)\s*\|`)

var ranges = make(map[string]bool)
var vars = make(map[string]bool)

func expandVar(v, chartName string) string {
	if !strings.HasPrefix(v, ".Values.global.") {
		v = ".Values." + chartName + "." + strings.TrimPrefix(v, ".Values.")
	}
	return v
}

func fixUsingTemplates(chartDir, chartName string) {
	filepath.Walk(chartDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if strings.HasSuffix(path, ".yaml") || strings.HasSuffix(path, ".tpl") {
			b, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			for _, found := range valueVarRe.FindAllStringSubmatch(string(b), -1) {
				for _, f := range found[1:] {
					if m := re2.FindStringSubmatch(f); m != nil {
						ranges[expandVar(m[3], chartName)] = true
						vars[expandVar(m[3], chartName)] = true
					} else if m := re3.FindStringSubmatch(f); m != nil {
						ranges[expandVar(m[1], chartName)] = true
						vars[expandVar(m[1], chartName)] = true
					} else if m := re1.FindAllString(f, -1); m != nil {
						for _, mm := range m {
							vars[expandVar(mm, chartName)] = true
						}
					}
				}
			}
		}

		return nil
	})
}

func fixVar(v string) {
	currentPath := make([]string, 0, 10)
	steps := strings.Split(strings.TrimPrefix(v, ".Values."), ".")

	obj := crd.Spec.Validation.OpenAPIV3Schema.Properties["spec"]
	for _, step := range steps {
		currentPath = append(currentPath, step)
		obj2, ok := obj.Properties[step]
		if !ok || obj2.Type == "" {
			if strings.Join(currentPath, ".") != v || ranges[".Values."+v] {
				obj.Properties[step] = CrdProperty{
					Description: "(added by fixVar)",
					Type:        "object",
					Properties:  make(map[string]CrdProperty),
				}
				obj2 = obj.Properties[step]
			} else {
				if obj.Type == "object" {
					obj.Properties[step] = CrdProperty{
						Description: "(added by fixVar)",
						Type:        "string",
						Properties:  nil,
					}
				}
				break
			}
		}
		obj = obj2
	}
}

func fixUsingAllCharts(dirName string) {
	filepath.Walk(dirName, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.Mode().IsRegular() && info.Name() == "Chart.yaml" {
			var chart interface{}

			b, err := ioutil.ReadFile(path)
			if err == nil {
				err = yaml.Unmarshal(b, &chart)
			}
			if err != nil {
				panic(err)
			}
			name := yamlField(chart, "name")
			if name != nil {
				chartDir := strings.TrimSuffix(path, "/Chart.yaml")
				fixUsingTemplates(chartDir+"/templates", name.(string))
			}
		}

		return nil
	})

	varsList := make([]string, 0, 32)
	for v, _ := range vars {
		varsList = append(varsList, v)
	}
	sort.Strings(varsList)
	for _, v := range varsList {
		fixVar(strings.TrimPrefix(v, ".Values."))
	}
}

func fixByName(obj *CrdProperty) {
	if obj.Type != "object" {
		return
	}
	for name, _ := range obj.Properties {
		if name == "env" {
			obj.Properties["env"] = envObject
		} else if name == "resources" {
			obj.Properties[name] = *resourcesObject.Clone()
		} else {
			prop := obj.Properties[name]
			fixByName(&prop)
		}
	}
}

func LoadConfig(dirName string) error {
	err := crd.Load(dirName + "/deploy/crds/charts_v1alpha1_xl_crd.yaml")
	if err == nil {
		err = requirements.Load(dirName + "/helm-charts/xl/requirements.yaml")
	}

	fixMissingConditions()
	fixMissingResourceLimits(dirName)
	fixRemainingValues(dirName)
	fixUsingAllCharts(dirName)

	setGenericObject("spec", "grafana", "grafana.ini")
	setGenericObject("spec", "grafana", "dashboards")
	setGenericObject("spec", "grafana", "env")
	setGenericObject("spec", "properties", "extractor")
	setGenericObject("spec", "elasticsearch", "cluster", "env")

	fixByName(&crd.Spec.Validation.OpenAPIV3Schema)

	return err
}

func RollBackup(fileName string) error {
	exists := func(name string) bool {
		_, err := os.Stat(name)
		return err == nil
	}

	rename := func(from, to string) error {
		if exists(from) {
			err := os.Rename(from, to)
			if err != nil {
				return err
			}
		}
		return nil
	}

	var err error

	if exists(fileName + ".bak8") {
		err = os.Remove(fileName + ".bak8")
	}
	if err == nil {
		err = rename(fileName+".bak7", fileName+".bak8")
	}
	if err == nil {
		err = rename(fileName+".bak6", fileName+".bak7")
	}
	if err == nil {
		err = rename(fileName+".bak5", fileName+".bak6")
	}
	if err == nil {
		err = rename(fileName+".bak4", fileName+".bak5")
	}
	if err == nil {
		err = rename(fileName+".bak3", fileName+".bak4")
	}
	if err == nil {
		err = rename(fileName+".bak2", fileName+".bak3")
	}
	if err == nil {
		err = rename(fileName+".bak", fileName+".bak2")
	}
	if err == nil {
		err = rename(fileName, fileName+".bak")
	}

	return err
}

func GetCrd() *Crd {
	return &crd
}

func GetRequirements() *Requirements {
	return &requirements
}
