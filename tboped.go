package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"./t8cop"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

var ticked = " ✓ "
var notTicked = " - "
var noDescription = "(no description available)"
var comment *tview.TextView
var exitButton = ""

type StringField struct {
	Label string
	Value **string
}

var StringFields = []StringField{
	{"Version tag:", &t8cop.GetOperator().Spec.Global.Tag},
	{"Exernal IP:", &t8cop.GetOperator().Spec.Global.ExternalIP},
	{"Exernal DB IP:", &t8cop.GetOperator().Spec.Global.ExternalDbIP},
	//	{"External Timescale DB IP:", &t8cop.GetOperator().Spec.Global.ExternalTimescaleDBIP },
}

// --------------------------------------------------------------------------------------------------------------------

func safeGet(str *string) string {
	if str == nil {
		return ""
	}
	return *str
}

func safeSet(str **string, value string) {
	if value == "" {
		*str = nil
	} else {
		*str = &value
	}
}

// --------------------------------------------------------------------------------------------------------------------

func confirmSave(app *tview.Application) *tview.Modal {
	modal := tview.NewModal()

	modal.
		SetText("Save changes?").
		AddButtons([]string{"Yes", "No"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Yes" || buttonLabel == "No" {
				exitButton = buttonLabel
				app.Stop()
			}
		})

	return modal
}

// --------------------------------------------------------------------------------------------------------------------

func populateLimitsForm(form *tview.Form) {
	form.Clear(false)
	limitNames := t8cop.GetNeededResourceNames()
	for _, k := range limitNames {
		v := t8cop.MemLimitMap[k]

		makeCallback := func(v **string) func(string) {
			return func(text string) {
				str := text
				*v = &str
			}
		}

		value := ""
		if *v != nil {
			value = **v
		}
		form.AddInputField(k, value, 10, nil, makeCallback(v))
	}
}

// --------------------------------------------------------------------------------------------------------------------

type CheckboxPlus struct {
	tview.Checkbox
}

// change the checkbox title colour depending on whether it is ticked or not
func (cb *CheckboxPlus) Draw(screen tcell.Screen) {
	if cb.IsChecked() {
		cb.SetLabelColor(tcell.ColorLightGreen)
	} else {
		cb.SetLabelColor(tcell.ColorRed)
	}
	cb.Checkbox.Draw(screen)
}

// update the text in the comment line when the checkbox gets focus to contains
// its description
func (cb *CheckboxPlus) Focus(delegate func(tview.Primitive)) {
	cb.Checkbox.Focus(delegate)
	mesg := t8cop.EnableFlagDescriptions[cb.GetLabel()]
	comment.SetText(mesg)
}

// --------------------------------------------------------------------------------------------------------------------

func populateComponentsForm(form, memLimitsForm *tview.Form) {
	form.Clear(true)

	componentNames := make([]string, 0, 10)

	for k, v := range t8cop.EnabledMap {
		if *v != nil {
			componentNames = append(componentNames, k)
		}
	}

	sort.Strings(componentNames)
	for _, k := range componentNames {
		v := t8cop.EnabledMap[k]
		makeCallback := func(v **bool) func(bool) {
			return func(checked bool) {
				b := checked
				*v = &b
				populateLimitsForm(memLimitsForm)
			}
		}
		cb := CheckboxPlus{*tview.NewCheckbox()}
		cb.SetLabel(k)
		cb.SetChecked(**v)
		cb.SetChangedFunc(makeCallback(v))
		form.AddFormItem(&cb)
	}
}

// --------------------------------------------------------------------------------------------------------------------

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

// --------------------------------------------------------------------------------------------------------------------

func addComponent(form *tview.Form, name string) {
	if *t8cop.EnabledMap[name] == nil {
		b := true
		*t8cop.EnabledMap[name] = &b
	}
}

// --------------------------------------------------------------------------------------------------------------------

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

// --------------------------------------------------------------------------------------------------------------------

func edit() {
	fileName := "/opt/turbonomic/kubernetes/operator/deploy/crds/charts_v1alpha1_xl_cr.yaml"

	if len(os.Args) >= 2 {
		fileName = os.Args[1]
	}

	err := t8cop.GetOperator().Load(fileName)
	if err != nil {
		panic(err)
	}

	// t8cop.GetRequirements().Load("t8c-install/operator/helm-charts/xl/requirements.yaml")

	app := tview.NewApplication()

	grid := tview.NewGrid()
	grid.SetRows(3, len(StringFields)+4, 0, 1, 1).SetColumns(0)
	grid.SetBorders(true).SetBordersColor(tcell.ColorBlue)

	header := tview.NewTextView().
		SetDynamicColors(true).
		SetText("[yellow]Turbonomic Operator - YAML Editor[-]\n\n(experimental, unsupported)").
		SetTextAlign(tview.AlignCenter)

	stringsHelp := strings.Join([]string{
		"<[white]UP[-]>/<[white]DOWN[-]> Navigate",
		"<[white]TAB[-]> Change Window",
		"<[white]ESC[-]> Save/Exit",
	}, ", ")

	componentsFormHelp := strings.Join([]string{
		"<[white]UP[-]>/<[white]DOWN[-]> Navigate",
		"<[white]SPACE[-]> Enable/Disable",
		"<[white]INS[-]> Add component",
		"<[white]DEL[-]> Delete",
		"<[white]TAB[-]> Change Window",
		"<[white]ESC[-]> Save/Exit",
	}, ", ")

	memLimitsFormHelp := strings.Join([]string{
		"<[white]UP[-]>/<[white]DOWN[-]> Navigate",
		"<[white]TAB[-]> Change Window",
		"<[white]ESC[-]> Save/Exit",
	}, ", ")

	footer := tview.NewTextView().
		SetDynamicColors(true).
		SetText(componentsFormHelp).
		SetTextColor(tcell.ColorBlue).
		SetTextAlign(tview.AlignCenter)

	stringsForm := tview.NewForm()
	stringsForm.
		SetItemPadding(0).
		SetFieldBackgroundColor(tcell.ColorBlack).
		SetBorder(true).
		SetTitle("  Configuration Params  ").
		SetBorderPadding(1, 1, 2, 2)

	for _, sf := range StringFields {
		stringsForm.AddInputField(sf.Label, safeGet(*sf.Value), 40, nil, func(text string) {
			safeSet(sf.Value, text)
		})
	}

	componentsForm := tview.NewForm()
	componentsForm.
		SetItemPadding(0).
		SetFieldBackgroundColor(tcell.ColorBlack).
		SetBorder(true).
		SetTitle("  Enabled Components  ").
		SetBorderPadding(1, 1, 2, 2)

	memLimitsForm := tview.NewForm()
	memLimitsForm.
		SetItemPadding(0).
		SetFieldBackgroundColor(tcell.ColorBlack).
		SetBorder(true).
		SetTitle("  POD Memory Resource Limits  ").
		SetBorderPadding(1, 1, 2, 2)

	comment = tview.NewTextView().
		SetTextColor(tcell.ColorWhite).
		SetText("Helpful comment here").
		SetTextAlign(tview.AlignCenter)

	populateComponentsForm(componentsForm, memLimitsForm)
	populateLimitsForm(memLimitsForm)

	focussedForm := func(form *tview.Form) {
		form.SetBackgroundColor(tcell.NewRGBColor(10, 10, 10))
		form.SetBorderColor(tcell.ColorWhite)
		form.SetBorderAttributes(tcell.AttrBold)
	}

	unfocussedForm := func(form *tview.Form) {
		form.SetBackgroundColor(tcell.NewRGBColor(10, 10, 10))
		form.SetBorderColor(tcell.ColorWhite)
		form.SetBorderAttributes(tcell.AttrDim)
	}

	focusStringsForm := func() {
		app.SetFocus(stringsForm)
		footer.SetText(stringsHelp)
		focussedForm(stringsForm)
		unfocussedForm(componentsForm)
		unfocussedForm(memLimitsForm)
		comment.SetText("Edit some basic configuration variables")
	}

	focusComponentsForm := func() {
		app.SetFocus(componentsForm)
		footer.SetText(componentsFormHelp)
		unfocussedForm(stringsForm)
		focussedForm(componentsForm)
		unfocussedForm(memLimitsForm)
	}

	focusMemLimitsForm := func() {
		app.SetFocus(memLimitsForm)
		footer.SetText(memLimitsFormHelp)
		unfocussedForm(stringsForm)
		unfocussedForm(componentsForm)
		focussedForm(memLimitsForm)
		comment.SetText("Edit memory limits for the different PODs")
	}

	var modal *tview.Modal = nil
	var confirmDeleteDialog *tview.Modal = nil
	var confirmItem = ""
	var componentSelector *tview.List = nil

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEscape:
			switch {
			case componentSelector != nil && componentSelector.HasFocus():
				app.SetRoot(grid, true)
				componentSelector = nil
				focusComponentsForm()
				return nil
			case componentsForm.HasFocus() || memLimitsForm.HasFocus() || stringsForm.HasFocus():
				modal = confirmSave(app)
				app.SetRoot(modal, true)
				app.SetFocus(modal)
				return nil
			case modal != nil && modal.HasFocus():
				app.SetRoot(grid, true)
				modal = nil
				focusStringsForm()
				return nil
			}

		case tcell.KeyTAB:
			switch {
			case stringsForm.HasFocus():
				focusComponentsForm()
				return nil
			case componentsForm.HasFocus():
				focusMemLimitsForm()
				return nil
			case memLimitsForm.HasFocus():
				focusStringsForm()
				return nil
			}

		case tcell.KeyEnter:
			switch {
			case componentSelector != nil && componentSelector.HasFocus():
				n := componentSelector.GetCurrentItem()
				text, _ := componentSelector.GetItemText(n)
				addComponent(componentsForm, strings.TrimSpace(text))
				populateComponentsForm(componentsForm, memLimitsForm)
				populateLimitsForm(memLimitsForm)
				app.SetRoot(grid, true)
				focusComponentsForm()
				componentSelector = nil
			}

		case tcell.KeyInsert:
			switch {
			case componentsForm.HasFocus():
				componentSelector = selectComponentToAdd(app)
				return nil
			}

		case tcell.KeyDelete:
			switch {
			case componentsForm.HasFocus():
				n, _ := componentsForm.GetFocusedItemIndex()
				confirmItem = strings.TrimSpace(componentsForm.GetFormItem(n).(*CheckboxPlus).GetLabel())
				confirmDeleteDialog = confirmDeleteComponent(app, confirmItem)
				return nil
			}

		case tcell.KeyRune:
			switch {
			case confirmDeleteDialog != nil && confirmDeleteDialog.HasFocus():
				switch event.Rune() {
				case rune('Y'), rune('y'):
					*t8cop.EnabledMap[confirmItem] = nil
					app.SetRoot(grid, true)
					populateComponentsForm(componentsForm, memLimitsForm)
					populateLimitsForm(memLimitsForm)
					focusComponentsForm()
					confirmDeleteDialog = nil
				case rune('N'), rune('n'):
					app.SetRoot(grid, true)
					focusComponentsForm()
					confirmDeleteDialog = nil
				}
			case componentsForm.HasFocus():
				switch event.Rune() {
				case 'D', 'd', '-':
					app.QueueEvent(tcell.NewEventKey(tcell.KeyDelete, 0, tcell.ModNone))
				case 'I', 'i', '+':
					app.QueueEvent(tcell.NewEventKey(tcell.KeyInsert, 0, tcell.ModNone))
				}
			}
		}
		return event
	})

	grid.AddItem(header, 0, 0, 1, 2, 0, 0, false)
	grid.AddItem(stringsForm, 1, 0, 1, 1, 0, 0, true)
	grid.AddItem(componentsForm, 2, 0, 1, 1, 0, 0, true)
	grid.AddItem(memLimitsForm, 1, 1, 2, 1, 0, 0, false)
	grid.AddItem(comment, 3, 0, 1, 2, 0, 0, false)
	grid.AddItem(footer, 4, 0, 1, 2, 0, 0, false)

	focusStringsForm()

	if err := app.SetRoot(grid, true).SetFocus(grid).Run(); err != nil {
		panic(err)
	}

	if exitButton == "Yes" {
		err = t8cop.GetOperator().Save(fileName)
		if err != nil {
			panic(err)
		}
	}
}

func enableEmbeddedReports() {
	fileName := "/opt/turbonomic/kubernetes/operator/deploy/crds/charts_v1alpha1_xl_cr.yaml"

	if len(os.Args) >= 3 {
		fileName = os.Args[2]
	}

	op := t8cop.GetOperator()
	err := op.Load(fileName)
	if err != nil {
		panic(err)
	}

	op.EnableEmbeddedReporting("grafanaAdmnPassword", "grafanaPassword")

	err = op.Save(fileName)
	if err != nil {
		panic(err)
	}
}

func main() {
	if len(os.Args) >= 2 && os.Args[1] == "--enable-embedded-reporting" {
		enableEmbeddedReports()
	} else {
		edit()
	}
}