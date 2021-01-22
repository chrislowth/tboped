package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"../t8cop"
)

var sf *os.File
var stdout = os.Stdout

func goName(name string) string {
	rtn := strings.ToUpper(name[0:1]) + name[1:]
	rtn = strings.Replace(rtn, "-", "_", -1)
	rtn = strings.Replace(rtn, ".", "_", -1)
	return strings.Replace(rtn, "/", "_", -1)
}

func isEmptyObject(objDef *t8cop.CrdProperty) bool {
	n := 0
	for _, _ = range objDef.Properties {
		n += 1
	}
	return n == 0
}

func generateGoStruct(objDefn *t8cop.CrdProperty, indent string) {
//	usedGoNames := make(map[string]bool)
	propNames := make([]string, 0, 16)
	for name, _ := range objDefn.Properties {
		propNames = append(propNames, name)
	}
	sort.Strings(propNames)
	for _, name := range propNames {
		if name == "PullPolicy" { // wrong leading cap
			continue;
		}
		gn := goName(name)
//		if usedGoNames[gn] {
//			continue;
//		}
//		usedGoNames[gn] = true

		def := objDefn.Properties[name]
		switch def.Type {
		case "string":
			fmt.Printf("%s%s *string `yaml:\"%s,omitempty\"`\n", indent+"\t", gn, name)
		case "array":
			fmt.Printf("%s%s []string `yaml:\"%s,omitempty\"`\n", indent+"\t", gn, name)
		case "boolean":
			fmt.Printf("%s%s *bool `yaml:\"%s,omitempty\"`\n", indent+"\t", gn, name)
		case "integer":
			fmt.Printf("%s%s *int64 `yaml:\"%s,omitempty\"`\n", indent+"\t", gn, name)
		case "object":
			if isEmptyObject(&def) || name == "annotations" || name == "podAnnotations" {
				fmt.Println(indent + "\t" + gn + " GenericMap `yaml:\"" + name + ",omitempty\"`")
			} else if len(propNames) > 80 {
				fmt.Fprintln(sf, "\ntype "+gn+"_t struct {")
				os.Stdout = sf
				generateGoStruct(&def, "")
				os.Stdout = stdout
				fmt.Fprintf(sf, "}\n")
				fmt.Println(indent + "\t" + gn + " " + gn + "_t  `yaml:\"" + name + ",omitempty\"`")
			} else {
				fmt.Println(indent + "\t" + gn + " struct {")
				generateGoStruct(&def, indent+"\t")
				fmt.Printf(indent+"\t} `yaml:\"%s,omitempty\"`\n", name)
			}
		default:
			panic("bad type: " + def.Type)
		}
	}
}

func generateEditorMapping() {
	names := make([]string, 0, 16)
	specProps := t8cop.GetCrd().Spec.Validation.OpenAPIV3Schema.Properties["spec"].Properties
	for name, _ := range specProps {
		names = append(names, name)
	}
	sort.Strings(names)

	enableFlagDescriptions := make(map[string]string)
	enableFlagNames := make([]string, 0, 16)
	fmt.Println("\nvar EnabledMap = map[string]**bool {")
	for _, name := range names {
		defn := specProps[name]
		prop, hasEnabled := defn.Properties["enabled"]
		if hasEnabled {
			enableFlagNames = append(enableFlagNames, name)
			fmt.Printf("\t\"%s\": &operator.Spec.%s.Enabled,\n", name, goName(name))
			if prop.Description != "" {
				enableFlagDescriptions[name] = prop.Description
			}
		}
	}
	fmt.Println("}")

	fmt.Println("\nvar EnableFlagDescriptions = map[string]string {")
	sort.Strings(enableFlagNames)
	for _, name := range enableFlagNames {
		if enableFlagDescriptions[name] != "" {
			fmt.Printf("\t\"%s\": \"%s\",\n", name, enableFlagDescriptions[name])
		}
	}
	fmt.Println("}")

	memLimitDescriptions := make(map[string]string)
	memLimitNames := make([]string, 0, 16)

	fmt.Println("\nvar MemLimitMap = map[string]**string {")
	for _, name := range names {
		defn := specProps[name]
		res, ok := defn.Properties["resources"]
		if ok {
			lim, ok2 := res.Properties["limits"]
			if ok2 {
				prop, ok3 := lim.Properties["memory"]
				if ok3 {
					memLimitNames = append(memLimitNames, name)
					fmt.Printf("\t\"%s\": &operator.Spec.%s.Resources.Limits.Memory,\n", name, goName(name))
					if prop.Description != "" {
						memLimitDescriptions[name] = prop.Description
					}
				}
			}
		}
	}
	fmt.Println("}")

	fmt.Println("\nvar MemLimitDescriptions = map[string]string {")
	sort.Strings(memLimitNames)
	for _, name := range memLimitNames {
		if memLimitDescriptions[name] != "" {
			fmt.Printf("\t\"%s\": \"%s\",\n", name, memLimitDescriptions[name])
		}
	}
	fmt.Println("}")
}

func main() {
	sf, _ = os.Create("structs.go")

	fmt.Fprintf(sf, "package t8cop\n")

	t8cop.LoadConfig(os.Args[1])

	fmt.Println("package t8cop")
	fmt.Println("")
	fmt.Println("type GenericMap map[interface{}]interface{}")
	fmt.Println("")
	fmt.Println("type Operator struct {")
	schema := t8cop.GetCrd().Spec.Validation.OpenAPIV3Schema
	generateGoStruct(&schema, "")
	fmt.Println("}")
	fmt.Println("")
	fmt.Println("var operator Operator")

	generateEditorMapping()

	fmt.Println("\nfunc init() {")
	fmt.Println("\trequirements = Requirements{")
	fmt.Println("\t\tDependencies: []Dependency{")
	for _, d := range t8cop.GetRequirements().Dependencies {
		fmt.Printf("\t\t\t{\"%s\", \"%s\", \"%s\", \"%s\"},\n", d.Name, d.Repository, d.Version, d.Condition)
	}
	fmt.Println("\t\t},")
	fmt.Println("\t}")
	fmt.Print("}")
}
