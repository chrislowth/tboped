package t8cop

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type CrdProperty struct {
	Description string 					`yaml:"description,omitempty"`
	Type string 						`yaml:"type,omitempty"`
	Properties map[string]CrdProperty	`yaml:"properties,omitempty"`
}

type Crd struct {
	ApiVersion string
	Kind string
	Metadata struct {
		Name string
		Annotations map[string]string
	}
	Spec struct {
		Group string
		Names struct {
			Kind string
			ListKind string 			`yaml:"listKind"`
			Plural string
			Singular string
		}
		Scope string
		// Subresources
		Version string
		Versions []struct {
			Name string
			Served bool
			Storage bool
		}
		Validation struct {
			OpenAPIV3Schema CrdProperty `yaml:"openAPIV3Schema"`
		}
	}
}

func (c *CrdProperty) Clone() *CrdProperty {
	rtn := CrdProperty{
		Description: c.Description,
		Type:        c.Type,
		Properties:  nil,
	}
	if (c.Properties != nil) {
		rtn.Properties = make(map[string]CrdProperty)
		for k, v := range c.Properties {
			rtn.Properties[k] = *v.Clone()
		}
	}
	return &rtn
}

// Load - typically from "t8c-install/operator/deploy/crds/charts_v1alpha1_xl_crd.yaml"
func (crd *Crd) Load(fileName string) error {
	b, err := ioutil.ReadFile(fileName)
	if err != nil { return err }

	err = yaml.Unmarshal(b, crd)
	if err != nil { return err }

	return nil
}

func (crd *Crd) Dump() error {
	b, err := yaml.Marshal(crd)
	if err != nil { return err }

	fmt.Print(string(b))

	return nil
}
