package ui

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"../t8cop"
	"io"
	"strings"
)

type StringField struct {
	Path string
	Label string
	Value **string
}

var StringFields = []StringField{
	{"spec.global.tag", "Version tag:", &t8cop.GetOperator().Spec.Global.Tag},
	{"spec.global.externalIP", "Exernal IP:", &t8cop.GetOperator().Spec.Global.ExternalIP},
	{"spec.global.externalDbIP", "Exernal DB IP:", &t8cop.GetOperator().Spec.Global.ExternalDbIP},
	//	{"External Timescale DB IP:", &t8cop.GetOperator().Spec.Global.ExternalTimescaleDBIP },
}

type ConfigForm struct {
	tview.Form
	changes map[string]string
	onTab func()
	onEsc func()
}

func InitConfigForm() *ConfigForm {
	rtn := ConfigForm{}
	rtn.Form = *tview.NewForm()
	rtn.changes = make(map[string]string)

	rtn.
		SetItemPadding(0).
		SetFieldBackgroundColor(tcell.ColorBlack).
		SetBorder(true).
		SetTitle("  Configuration Params  ").
		SetBorderPadding(1, 1, 2, 2)

	for _, sf := range StringFields {
		makeChangedFunc := func(sf StringField) func(string) {
			return func(text string) {
				safeSet(sf.Value, text)
				rtn.changes[sf.Path] = text
			}
		}
		rtn.AddInputField(sf.Label, safeGet(*sf.Value), 40, nil, makeChangedFunc(sf))
	}

	return &rtn
}

func (f *ConfigForm) HandleEvent(event *tcell.EventKey) bool {
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

func (f *ConfigForm) OnTab(fn func()) {
	f.onTab = fn
}

func (f *ConfigForm) OnEsc(fn func()) {
	f.onEsc = fn
}

func (f *ConfigForm) EmitChanges(file io.Writer) {
	for k, v := range f.changes {
		if v == "" {
			emitScript(file,"delete", k, nil)
		} else {
			emitScript(file,"update", k, v)
		}
	}
}

func (f *ConfigForm) GetHelpText() string {
	return strings.Join([]string{
		"<[white]UP[-]>/<[white]DOWN[-]> Navigate",
		"<[white]TAB[-]> Change Window",
		"<[white]ESC[-]> Save/Exit",
	}, ", ")
}
