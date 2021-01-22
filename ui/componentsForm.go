package ui

import (
	"fmt"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"io"
	"sort"
	"../t8cop"
	"strings"
)

type ComponentsForm struct {
	tview.Form
	changes map[string]*bool
	onTab func()
	onEsc func()
	focusFunc func()
	confirmItem string
	confirmDeleteDialog *tview.Modal
	componentSelector *tview.List
	app *tview.Application
	limitsForm *MemLimitsForm
	comment *tview.TextView
}

func InitComponentsForm() *ComponentsForm {
	rtn := ComponentsForm{}
	rtn.Form = *tview.NewForm()
	rtn.changes = make(map[string]*bool)

	rtn.Form.
		SetItemPadding(0).
		SetFieldBackgroundColor(tcell.ColorBlack).
		SetBorder(true).
		SetTitle("  Enabled Components  ").
		SetBorderPadding(1, 1, 2, 2)

	return &rtn
}

func (f *ComponentsForm) Set(name string, checked bool) {
	b := checked
	f.changes[name] = &b
	*t8cop.EnabledMap[name] = &b
}

func (f *ComponentsForm) Populate() {
	f.Clear(true)

	componentNames := make([]string, 0, 10)

	for k, v := range t8cop.EnabledMap {
		if *v != nil {
			componentNames = append(componentNames, k)
		}
	}

	sort.Strings(componentNames)
	for _, k := range componentNames {
		v := t8cop.EnabledMap[k]
		makeCallback := func(name string, v **bool) func(bool) {
			return func(checked bool) {
				b := checked
				*v = &b
				f.changes[name] = &b
				f.limitsForm.Populate()
			}
		}
		cb := CheckboxPlus{}
		cb.Checkbox = *tview.NewCheckbox()
		cb.SetLabel(k)
		cb.SetChecked(**v)
		cb.SetChangedFunc(makeCallback(k, v))
		cb.SetCommentLine(f.comment)
		f.AddFormItem(&cb)
	}
}

func (f *ComponentsForm) GetHelpText() string {
	return 	strings.Join([]string{
		"<[white]UP[-]>/<[white]DOWN[-]> Navigate",
		"<[white]SPACE[-]> Enable/Disable",
		"<[white]INS[-]> Add component",
		"<[white]DEL[-]> Delete",
		"<[white]TAB[-]> Change Window",
		"<[white]ESC[-]> Save/Exit",
	}, ", ")
}

func confirmDeleteComponent(app *tview.Application, name string) *tview.Modal {
	modal := tview.NewModal()

	modal.
		SetText("Confirm: delete component '" + name + "' ?").
		AddButtons([]string{"Yes", "No"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Yes" || buttonLabel == "No" {
				app.QueueEvent(tcell.NewEventKey(tcell.KeyRune, rune(buttonLabel[0]), tcell.ModNone))
			}
		})

	app.SetRoot(modal, true)
	app.SetFocus(modal)

	return modal
}


func selectComponentToAdd(app *tview.Application) *tview.List {
	list := tview.NewList()
	list.SetRect(10, 2, 50, 25)
	list.SetBorder(true)
	list.ShowSecondaryText(false)
	list.SetWrapAround(false)
	list.SetTitle("  Select the component to add  ")

	var names = make([]string, 0, 10)

	for k, v := range t8cop.EnabledMap {
		if *v == nil {
			names = append(names, k)
		}
	}
	sort.Strings(names)

	for _, name := range names {
		list.AddItem(fmt.Sprintf("  %-48.48s", name), "", 0, func() {})
	}

	app.SetRoot(list, false).SetFocus(list)

	return list
}


func (f *ComponentsForm) HandleDeleteDialogEvent(event *tcell.EventKey) bool {
	if f.confirmDeleteDialog == nil || !f.confirmDeleteDialog.HasFocus() {
		return false
	}
	switch event.Key() {
	case tcell.KeyRune:
		switch event.Rune() {
		case rune('Y'), rune('y'):
			*t8cop.EnabledMap[f.confirmItem] = nil
			f.changes[f.confirmItem] = nil
			f.Populate()
			f.focusFunc()
			f.confirmDeleteDialog = nil
			return true
		case rune('N'), rune('n'):
			f.focusFunc()
			f.confirmDeleteDialog = nil
			return true
		}
	}
	return false
}

func (f *ComponentsForm) addComponent(name string) {
	if *t8cop.EnabledMap[name] == nil {
		b1 := true
		b2 := true
		f.changes[name] = &b1
		*t8cop.EnabledMap[name] = &b2
	}
}

func (f *ComponentsForm) HandleSelectorDialogEvent(event *tcell.EventKey) bool {
	if f.componentSelector == nil || !f.componentSelector.HasFocus() {
		return false
	}
	switch event.Key() {

	case tcell.KeyEscape:
		f.componentSelector = nil
		f.focusFunc()
		return true

	case tcell.KeyEnter:
		n := f.componentSelector.GetCurrentItem()
		text, _ := f.componentSelector.GetItemText(n)
		f.addComponent(strings.TrimSpace(text))
		f.Populate()
		f.focusFunc()
		f.componentSelector = nil

	case tcell.KeyRune:
		switch event.Rune() {

		}
	}
	return false
}

func (f *ComponentsForm) HandleEvent(event *tcell.EventKey) bool {
	if f.HandleDeleteDialogEvent(event) || f.HandleSelectorDialogEvent(event) {
		return true
	}

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

	case tcell.KeyDelete:
		n, _ := f.GetFocusedItemIndex()
		f.confirmItem = strings.TrimSpace(f.GetFormItem(n).(*CheckboxPlus).GetLabel())
		f.confirmDeleteDialog = confirmDeleteComponent(f.app, f.confirmItem)
		return true

	case tcell.KeyInsert:
		f.componentSelector = selectComponentToAdd(f.app)
		return true

	case tcell.KeyRune:
		switch event.Rune() {
		case 'D', 'd', '-':
			return f.HandleEvent(tcell.NewEventKey(tcell.KeyDelete, 0, tcell.ModNone))
		case 'I', 'i', '+':
			return f.HandleEvent(tcell.NewEventKey(tcell.KeyInsert, 0, tcell.ModNone))
		}
	}

	return false
}

func (f *ComponentsForm) OnTab(fn func()) {
	f.onTab = fn
}

func (f *ComponentsForm) OnEsc(fn func()) {
	f.onEsc = fn
}

func (f *ComponentsForm) SetFocusFunc(fn func()) {
	f.focusFunc = fn
}

func (f *ComponentsForm) SetLimitsForm(fm *MemLimitsForm) {
	f.limitsForm = fm
}

func (f *ComponentsForm) SetApp(app *tview.Application) {
	f.app = app
}

func (f *ComponentsForm) SetComment(comment *tview.TextView) {
	f.comment = comment
}

func (f *ComponentsForm) EmitChanges(file io.Writer) {
	for k, v := range f.changes {
		if v == nil {
			emitScript(file,"delete", "spec."+k, nil)
		} else if *v == true {
			emitScript(file,"update", "spec."+k+".enabled", true)
		} else {
			emitScript(file,"update", "spec."+k+".enabled", false)
		}
	}
}
