package t8cop

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Dependency struct {
	Name       string
	Repository string
	Version    string
	Condition  string
}

type Requirements struct {
	Dependencies []Dependency
}

// Load from t8c-install/operator/helm-charts/xl/requirements.yaml
func (r *Requirements) Load(fileName string) error {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(b, r)
	if err != nil {
		return err
	}

	return nil
}
