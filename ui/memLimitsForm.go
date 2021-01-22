package ui

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"io"
	"strings"
	"../t8cop"
)

type MemLimitsForm struct {
	tview.Form
	changes map[string]string
	onTab func()
	onEsc func()
}

func InitMemLimitsForm() *MemLimitsForm {
	memLimitsForm := tview.NewForm()
	memLimitsForm.
		SetItemPadding(0).
		SetFieldBackgroundColor(tcell.ColorBlack).
		SetBorder(true).
		SetTitle("  POD Memory Resource Limits  ").
		SetBorderPadding(1, 1, 2, 2)
	rtn := &MemLimitsForm{}
	rtn.Form = *memLimitsForm
	rtn.changes = make(map[string]string)

	return rtn
}

func (f *MemLimitsForm) Populate() {
	f.Clear(false)
	limitNames := t8cop.GetNeededResourceNames()
	for _, k := range limitNames {
		v := t8cop.MemLimitMap[k]

		makeCallback := func(name string, v **string) func(string) {
			return func(text string) {
				f.changes[name] = text
				str := text
				*v = &str
			}
		}

		value := ""
		if *v != nil {
			value = **v
		}
		f.AddInputField(k, value, 10, nil, makeCallback(k, v))
	}
}

func (f *MemLimitsForm) GetHelpText() string {
	return strings.Join([]string{
		"<[white]UP[-]>/<[white]DOWN[-]> Navigate",
		"<[white]TAB[-]> Change Window",
		"<[white]ESC[-]> Save/Exit",
	}, ", ")
}

func (f *MemLimitsForm) OnTab(fn func()) {
	f.onTab = fn
}

func (f *MemLimitsForm) OnEsc(fn func()) {
	f.onEsc = fn
}

func (f *MemLimitsForm) HandleEvent(event *tcell.EventKey) bool {
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

func (f *MemLimitsForm) EmitChanges(file io.Writer) {
	for k, v := range f.changes {
		if v == "" {
			emitScript(file,"update", "spec."+k+".resources.limits.memory", nil )
		} else {
			emitScript(file,"update", "spec."+k+".resources.limits.memory", v)
		}
	}
}