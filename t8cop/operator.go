// +build !mkopgo

package t8cop

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

func GetOperator() *Operator {
	return &operator
}

// Load - typically from "t8c-install/operator/deploy/crds/charts_v1alpha1_xl_cr.yaml"
func (op *Operator) Load(fileName string) error {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(b, op)
	if err != nil {
		return err
	}

	return nil
}

func (op *Operator) Dump() error {
	b, err := yaml.Marshal(op)
	if err != nil {
		return err
	}

	fmt.Print(string(b))

	return nil
}

func (op *Operator) Save(fileName string) error {
	b, err := yaml.Marshal(op)
	if err != nil {
		return err
	}

	err = RollBackup(fileName)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(fileName, b, 0666)
}

func (op *Operator) EnableEmbeddedReporting(adminPassword, password string) {
	_true := true

	spec := &op.Spec

	spec.Global.ExternalTimescaleDBIP = spec.Global.ExternalIP

	spec.Grafana.Enabled = &_true
	spec.Grafana.AdminPassword = &adminPassword
	if spec.Grafana.Grafana_ini == nil {
		spec.Grafana.Grafana_ini = make(GenericMap)
	}
	if spec.Grafana.Grafana_ini["database"] == nil {
		spec.Grafana.Grafana_ini["database"] = make(GenericMap)
	}
	spec.Grafana.Grafana_ini["database"].(GenericMap)["type"] = "postgres"
	spec.Grafana.Grafana_ini["database"].(GenericMap)["password"] = password

	if spec.Properties.Extractor == nil {
		spec.Properties.Extractor = make(GenericMap)
	}
	spec.Properties.Extractor["grafanaAdminPassword"] = adminPassword

	spec.Reporting.Enabled = &_true
	spec.Timescaledb.Enabled = &_true
	spec.Extractor.Enabled = &_true
}

func (op *Operator) EnableSingleSignon() {
	_true := true
	spec := &op.Spec
	spec.Properties.Api.SamlEnabled = &_true
	// spec.Properties.Api.SamlWebSsoEndpoint = ...
	// spec.Properties.Api.SamlEntityId =
	// spec.Properties.Api.SamlRegistrationId =
	// spec.Properties.Api.SamlIdpCertificate =
}

func (op *Operator) DisableSingleSignon() {
	_false := false
	spec := &op.Spec
	spec.Properties.Api.SamlEnabled = &_false
	spec.Properties.Api.SamlWebSsoEndpoint = nil
	spec.Properties.Api.SamlEntityId = nil
	spec.Properties.Api.SamlRegistrationId = nil
	spec.Properties.Api.SamlIdpCertificate = nil
}

func (op *Operator) EnableTurboOnTurbo() {

}

func (op *Operator) DisableTurboOnTurbo() {

}

func GetNeededResourceNames() []string {
	rtn := make([]string, 0, 16)

	for k, _ := range MemLimitMap {
		// find matching requirements condition
		cond := ""
		for _, dep := range requirements.Dependencies {
			if dep.Name == k {
				cond = dep.Condition
				break
			}
		}
		if cond == "" {
			continue
		}
		isNeeded := true
		for _, condVar := range strings.Split(cond, ",") {
			if strings.HasSuffix(condVar, ".enabled") {
				condVar = strings.TrimSuffix(condVar, ".enabled")
				if *EnabledMap[condVar] == nil || **EnabledMap[condVar] != true {
					isNeeded = false
				}
			}
		}
		if isNeeded {
			rtn = append(rtn, k)
		}
	}

	sort.Strings(rtn)
	return rtn
}
