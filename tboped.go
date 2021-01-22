package main

import (
	"./t8cop"
	"./ui"
	"fmt"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"io/ioutil"
	"os"
	"os/exec"
)

var comment *tview.TextView
var exitButton = ""

var True = true
var False = false

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

func edit() {
	fileName := "/opt/turbonomic/kubernetes/operator/deploy/crds/charts_v1alpha1_xl_cr.yaml"

	if len(os.Args) >= 2 {
		fileName = os.Args[1]
	}

	err := t8cop.GetOperator().Load(fileName)
	if err != nil {
		panic(err)
	}

	app := tview.NewApplication()

	header := tview.NewTextView().
		SetDynamicColors(true).
		SetText("[yellow]Turbonomic Operator - YAML Editor[-]\n\n(experimental, unsupported)").
		SetTextAlign(tview.AlignCenter)

	stringsForm := ui.InitConfigForm()
	stringsHelp := stringsForm.GetHelpText()

	componentsForm := ui.InitComponentsForm()
	componentsForm.SetComment(comment)
	componentsFormHelp := componentsForm.GetHelpText()

	memLimitsForm := ui.InitMemLimitsForm()
	memLimitsFormHelp := memLimitsForm.GetHelpText()

	reportingForm := ui.InitReportingForm(componentsForm, memLimitsForm)
	reportingHelp := reportingForm.GetHelpText()

	comment = tview.NewTextView().
		SetTextColor(tcell.ColorWhite).
		SetText("Helpful comment here").
		SetTextAlign(tview.AlignCenter)

	componentsForm.Populate()
	memLimitsForm.Populate()

	footer := tview.NewTextView().
		SetDynamicColors(true).
		SetText(componentsFormHelp).
		SetTextColor(tcell.ColorBlue).
		SetTextAlign(tview.AlignCenter)

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

	grid := tview.NewGrid()
	grid.SetRows(
		3,
		stringsForm.GetFormItemCount()+4,
		reportingForm.GetFormItemCount()+4,
		0,
		1,
		1,
	).SetColumns(0)
	grid.SetBorders(true).SetBordersColor(tcell.ColorBlue)

	focusStringsForm := func() {
		app.SetFocus(stringsForm)
		footer.SetText(stringsHelp)
		focussedForm(&stringsForm.Form)
		unfocussedForm(&componentsForm.Form)
		unfocussedForm(&memLimitsForm.Form)
		unfocussedForm(&reportingForm.Form)
		comment.SetText("Edit some basic configuration variables")
	}

	focusReportingForm := func() {
		app.SetFocus(reportingForm)
		footer.SetText(reportingHelp)
		unfocussedForm(&stringsForm.Form)
		unfocussedForm(&componentsForm.Form)
		unfocussedForm(&memLimitsForm.Form)
		focussedForm(&reportingForm.Form)
	}

	focusComponentsForm := func() {
		app.SetRoot(grid, true)
		app.SetFocus(componentsForm)
		footer.SetText(componentsFormHelp)
		unfocussedForm(&stringsForm.Form)
		focussedForm(&componentsForm.Form)
		unfocussedForm(&memLimitsForm.Form)
		unfocussedForm(&reportingForm.Form)
	}

	focusMemLimitsForm := func() {
		app.SetFocus(memLimitsForm)
		footer.SetText(memLimitsFormHelp)
		unfocussedForm(&stringsForm.Form)
		unfocussedForm(&componentsForm.Form)
		focussedForm(&memLimitsForm.Form)
		unfocussedForm(&reportingForm.Form)
		comment.SetText("Edit memory limits for the different PODs")
	}

	var modal *tview.Modal = nil

	escapePressed := func() {
		modal = confirmSave(app)
		app.SetRoot(modal, true)
		app.SetFocus(modal)
	}

	stringsForm.OnTab(focusReportingForm)
	stringsForm.OnEsc(escapePressed)

	reportingForm.OnTab(focusComponentsForm)
	reportingForm.OnEsc(escapePressed)

	componentsForm.SetApp(app)
	componentsForm.OnTab(focusMemLimitsForm)
	componentsForm.OnEsc(escapePressed)
	componentsForm.SetLimitsForm(memLimitsForm)
	componentsForm.SetFocusFunc(focusComponentsForm)

	memLimitsForm.OnTab(focusStringsForm)
	memLimitsForm.OnEsc(escapePressed)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if stringsForm.HandleEvent(event) ||
			reportingForm.HandleEvent(event) ||
			componentsForm.HandleEvent(event) ||
			memLimitsForm.HandleEvent(event) {
			return nil
		}

		switch event.Key() {
		case tcell.KeyEscape:
			switch {
			case modal != nil && modal.HasFocus():
				app.SetRoot(grid, true)
				modal = nil
				focusStringsForm()
				return nil
			}
		}

		return event
	})

	grid.AddItem(header, 0, 0, 1, 2, 0, 0, false)
	grid.AddItem(stringsForm, 1, 0, 1, 1, 0, 0, true)
	grid.AddItem(reportingForm, 2, 0, 1, 1, 0, 0, false)
	grid.AddItem(componentsForm, 3, 0, 1, 1, 0, 0, true)
	grid.AddItem(memLimitsForm, 1, 1, 3, 1, 0, 0, false)
	grid.AddItem(comment, 4, 0, 1, 2, 0, 0, false)
	grid.AddItem(footer, 5, 0, 1, 2, 0, 0, false)

	focusStringsForm()

	if err := app.SetRoot(grid, true).SetFocus(grid).Run(); err != nil {
		panic(err)
	}

	if exitButton == "Yes" {
		f, err := ioutil.TempFile(os.TempDir(), "tboped-")
		if err != nil {
			panic(err)
		}
		memLimitsForm.EmitChanges(f)
		componentsForm.EmitChanges(f)
		reportingForm.EmitChanges(f)
		stringsForm.EmitChanges(f)
		f.Close()

		out, err := exec.Command("yq", "w", "-i", "-s", f.Name(), fileName).CombinedOutput()
		fmt.Print(string(out))
		if err != nil {
			panic(err)
		}
		os.Remove(f.Name())
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
