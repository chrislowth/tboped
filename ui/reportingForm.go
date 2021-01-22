package ui

import (
	"fmt"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"io"
	"strings"
	"../t8cop"
)

var grafanaDbPass *string

var ReportingFields = []StringField{
	{"spec.global.externalTimescaleDBIP", "Database IP:", &t8cop.GetOperator().Spec.Global.ExternalTimescaleDBIP },
	{ "spec.grafana.adminPassword", "DB Admin Password:", &t8cop.GetOperator().Spec.Grafana.AdminPassword },
	{ "spec.grafana.\"grafana.ini\".database.password", "DB User Password:", &grafanaDbPass },
}

type ReportingForm struct {
	tview.Form
	changes map[string]string
	onTab func()
	onEsc func()
}

func InitReportingForm(cf *ComponentsForm, mf *MemLimitsForm) *ReportingForm {
	blankPass := ""
	grafanaDbPass = &blankPass

	dbStruct, ok := t8cop.GetOperator().Spec.Grafana.Grafana_ini["database"].(map[string]interface{})
	if ok {
		pass, ok := dbStruct["password"].(string)
		if ok {
			grafanaDbPass = &pass
		}
	}

	rtn := ReportingForm{}
	rtn.Form = *tview.NewForm()
	rtn.changes = make(map[string]string)

	rtn.
		SetItemPadding(0).
		SetFieldBackgroundColor(tcell.ColorBlack).
		SetBorder(true).
		SetTitle("  Embedded Reporting  ").
		SetBorderPadding(1, 1, 2, 2)

	enabled := false
	if t8cop.GetOperator().Spec.Reporting.Enabled != nil {
		enabled = *t8cop.GetOperator().Spec.Reporting.Enabled
	}

	cb := CheckboxPlus{}
	cb.Checkbox = *tview.NewCheckbox()
	cb.SetLabel("Enabled?")
	cb.SetChecked(enabled)
	cb.SetChangedFunc(func (checked bool) {
		rtn.changes["enabled"] = fmt.Sprintf("%v", checked)
		cf.Set("grafana", checked)
		cf.Set("reporting", checked)
		cf.Set("timescaledb", checked)
		cf.Set("extractor", checked)

		g := t8cop.GetOperator().Spec.Grafana
		if g.Grafana_ini == nil {
			g.Grafana_ini = make(t8cop.GenericMap)
		}
		if g.Grafana_ini["database"] == nil {
			g.Grafana_ini["database"] = make(map[string]interface{})
		}
		if g.Grafana_ini["database"].(map[string]interface{})["type"] != "postgres" {
			g.Grafana_ini["database"].(map[string]interface{})["type"] = "postgress"
			rtn.changes["spec.grafana.\"grafana.ini\".database.type"] = "postgres"
		}

		cf.Populate()
		mf.Populate()
	})
	rtn.AddFormItem(&cb)
//	cb.SetCommentLine(f.comment)

	for _, sf := range ReportingFields {
		makeChangedFunc := func(sf StringField) func(string) {
			return func(text string) {
				safeSet(sf.Value, text)
				rtn.changes[sf.Path] = text
			}
		}
		rtn.AddInputField(sf.Label, safeGet(*sf.Value), 40, nil, makeChangedFunc(sf))
	}

//	rtn.AddInputField("DB IP", safeGet(t8cop.GetOperator().Spec.Global.ExternalTimescaleDBIP), 40, nil, func (text string) {
//		safeSet(&t8cop.GetOperator().Spec.Global.ExternalTimescaleDBIP, text)
//		rtn.changes["spec.global.externalTimescaleDBIP"] = text
//	})

	return &rtn
}

func (f *ReportingForm) HandleEvent(event *tcell.EventKey) bool {
	if !f.HasFocus() {
		return false
	}

	switch event.Key() {
	case tcell.KeyTAB:
		if f.onTab != nil {
			f.onTab()
			return true
		}
	case tcell.KeyEscape:
		if f.onEsc != nil {
			f.onEsc()
			return true
		}
	}

	return false
}

func (f *ReportingForm) OnTab(fn func()) {
	f.onTab = fn
}

func (f *ReportingForm) OnEsc(fn func()) {
	f.onEsc = fn
}

func (f *ReportingForm) EmitChanges(file io.Writer) {
	enabled := f.changes["enabled"]
	if enabled != "" {
		emitScript(file, "update", "spec.reporting.enabled", enabled == "true")
	}
	if enabled == "true" {
		ap := *t8cop.GetOperator().Spec.Grafana.AdminPassword

		if t8cop.GetOperator().Spec.Properties.Extractor == nil {
			t8cop.GetOperator().Spec.Properties.Extractor = make(t8cop.GenericMap)
		}
		if t8cop.GetOperator().Spec.Properties.Extractor["grafanaAdminPassword"] != ap {
			emitScript(file, "update", "spec.properties.extractor.grafanaAdminPassword", ap)
		}
	}
	for k, v := range f.changes {
		if k == "enabled" {
			continue
		}
		if v == "" {
			emitScript(file,"delete", k, nil)
		} else {
			emitScript(file,"update", k, v)
		}
	}
}

func (f *ReportingForm) GetHelpText() string {
	return strings.Join([]string{
		"<[white]UP[-]>/<[white]DOWN[-]> Navigate",
		"<[white]TAB[-]> Change Window",
		"<[white]ESC[-]> Save/Exit",
	}, ", ")
}
