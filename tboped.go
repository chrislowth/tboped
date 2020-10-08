package main

import (
	"errors"
	"fmt"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"io/ioutil"
	"os"
	"os/exec"
	"sort"
	"strings"
)

var ticked = " ✓ "
var notTicked = " - "
var noDescription = "(no description available)"

// --------------------------------------------------------------------------------------------------------------------

var exitButton = ""

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

type YamlFile struct {
	Lines         []string
	LineNumByPath map[string]int
	EnabledFlags  map[string]int
	OptionNames   []string
	Options       map[string]string
}

func newYamlFile() *YamlFile {
	return &YamlFile{
		Lines:         make([]string, 0, 32),
		LineNumByPath: make(map[string]int),
		EnabledFlags:  make(map[string]int),
		OptionNames:   make([]string, 0, 32),
		Options:       make(map[string]string),
	}
}

func loadFile(fileName string, pathPrefix *string) (*YamlFile, error) {
	rtn := newYamlFile()

	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	path := make([]string, 0, 4)
	indents := make([]int, 0, 4)

	if pathPrefix == nil {
		path = append(path, ".")
		indents = append(indents, -1, 0)
	} else {
		path = append(path, "spec", ".")
		indents = append(indents, -1, -1, 0)

	}

	rtn.Lines = strings.Split(strings.TrimSpace(string(b)), "\n")

	for lineNum, line := range rtn.Lines {
		if strings.HasPrefix(strings.TrimSpace(line), "#") || strings.TrimSpace(line) == "" {
			continue
		}
		trimmed := strings.TrimLeft(line, " \t")
		indent := len(line) - len(trimmed)
		kv := strings.SplitN(trimmed, ":", 2)
		if len(kv) != 2 {
			continue
		}

		lastIndent := indents[len(indents)-1]

		if indent == lastIndent {
			path[len(path)-1] = kv[0]
		} else if indent > lastIndent {
			indents = append(indents, indent)
			path = append(path, kv[0])
		} else {
			for indents[len(indents)-1] >= indent {
				indents = indents[0:len(indents)-1]
				path = path[0:len(path)-1]
			}
			indents = append(indents, indent)
			path = append(path, kv[0])
		}
		rtn.LineNumByPath[strings.Join(path, "/")] = lineNum
		if len(path) == 3 && path[0] == "spec" && path[2] == "enabled" {
			rtn.EnabledFlags[path[1]] = lineNum
		}
	}

	for k, lineNum := range rtn.EnabledFlags {
		rtn.OptionNames = append(rtn.OptionNames, k)
		f := strings.Split(rtn.Lines[lineNum], ":")
		rtn.Options[k] = strings.TrimSpace(f[1])
	}

	sort.Strings(rtn.OptionNames)

	return rtn, nil
}


func (y *YamlFile) GetField(path string) (string, bool) {
	lineNum, ok := y.LineNumByPath[path]
	if !ok {
		return "", false
	}
	f := strings.SplitN(y.Lines[lineNum], ":", 2)
	return strings.TrimSpace(f[1]), true
}

// --------------------------------------------------------------------------------------------------------------------

func getOperatorPodName() (string, error) {
	podBytes, err := exec.Command("kubectl", "get", "pods", "-n", "turbonomic").CombinedOutput()
	if err != nil { return "", err }
	for _, line := range strings.Split(string(podBytes), "\n") {
		if strings.HasPrefix(line, "t8c-operator-") {
			f := strings.Split(line, " ")
			return f[0], nil
		}
	}
	return "", errors.New("Pod not found")
}

// --------------------------------------------------------------------------------------------------------------------

func loadOperatorYaml(podName, fileName string) (*YamlFile, error) {
	yamlBytes, err := exec.Command("kubectl", "exec", podName, "-n", "turbonomic", "--", "cat", fileName).CombinedOutput()
	if err != nil { return nil, err }

	f, err := ioutil.TempFile("/tmp", "tboped-")
	if err != nil { return nil, err }

	f.Write(yamlBytes)
	f.Close()

	spec := "spec"
	rtn, err := loadFile(f.Name(), &spec)
	os.Remove(f.Name())
	return rtn, err
}

// --------------------------------------------------------------------------------------------------------------------

func listOperatorDirectory(podName, dirName string) ([]string, error) {
	b, err := exec.Command("kubectl", "exec", podName, "-n", "turbonomic", "--", "ls", dirName ).CombinedOutput()
	if err != nil { return nil, err }

	return strings.Split(strings.TrimSpace(string(b)), "\n"), nil
}

// --------------------------------------------------------------------------------------------------------------------

func loadValuesYaml(podName string) (*YamlFile, error) {
	fileName := "/opt/helm/helm-charts/xl/values.yaml"
	return loadOperatorYaml(podName, fileName)
}

// --------------------------------------------------------------------------------------------------------------------

func (crd *YamlFile) getDescription(name string) string {
	descr, _ := crd.GetField("spec/validation/openAPIV3Schema/properties/spec/properties/"+name+"/properties/enabled/description")
	if descr == "" {
		descr = name
	}
	return descr
}

// --------------------------------------------------------------------------------------------------------------------

func safeToAdd(feature string) bool {
	return true
}

// --------------------------------------------------------------------------------------------------------------------

func main() {
	fileName := "/opt/turbonomic/kubernetes/operator/deploy/crds/charts_v1alpha1_xl_cr.yaml"
	crdName := "/opt/turbonomic/kubernetes/operator/deploy/crds/charts_v1alpha1_xl_crd.yaml"

	if len(os.Args) >= 2 {
		fileName = os.Args[1]
	}
	if len(os.Args) >= 3 {
		crdName = os.Args[2]
	}

	fmt.Printf("Loading '%s' .. ", fileName)
	file, err := loadFile(fileName, nil)
	if err != nil {
		fmt.Println("")
		panic(err)
	}
	fmt.Println("ok")

	fmt.Printf("Loading '%s' .. ", crdName)
	crd, err := loadFile(crdName, nil)
	if err != nil {
		fmt.Println("")
		panic(err)
	}
	fmt.Println("ok")

	fmt.Print("Finding t8c-operator pod .. ")
	podName, err := getOperatorPodName()
	if err != nil { panic(err) }
	fmt.Printf("%s\n", podName)

	fmt.Print("Load values.yaml from pod .. ")
	values, err := loadValuesYaml(podName)
	if err != nil { panic(err) }
	fmt.Println("ok")

	// addable contains the list of options that can be enabled but are not
	// listed in the maim file
	addable := make([]string, 0, 10)
	for _, o := range values.OptionNames {
		_, found := file.EnabledFlags[o]
		if !found && safeToAdd(o) {
			line := values.Lines[values.LineNumByPath["spec/"+o+"/enabled"]]
			if strings.HasSuffix(strings.TrimSpace(line), "false") {
				addable = append(addable, o)
			}
		}
	}
	app := tview.NewApplication()

	grid := tview.NewGrid()
	grid.SetRows(3, 0, 1, 1).SetColumns(0)
	grid.SetBorders(true).SetBordersColor(tcell.ColorBlue)

	header := tview.NewTextView().
		SetText("\nTurbonomic Operator - Edit Tool").
		SetTextColor(tcell.ColorBlack).
		SetTextAlign(tview.AlignCenter).
		SetBackgroundColor(tcell.ColorWhite)

	formHelp := "<[white]UP[-]>/<[white]DOWN[-]> Navigate, <[white]SPACE[-]> Select/Deselect, <[white]TAB[-]> Change Window, <[white]ESC[-]> Save and/or Exit"
	listHelp := "<[white]UP[-]>/<[white]DOWN[-]> Navigate, <[white]RETURN[-]> Select, <[white]TAB[-]> Change Window, <[white]ESC[-]> Save and/or Exit"

	footer := tview.NewTextView().
		SetDynamicColors(true).
		SetText(formHelp).
		SetTextColor(tcell.ColorBlue).
		SetTextAlign(tview.AlignCenter)

	version, _ := file.GetField("spec/global/tag")

	form := tview.NewForm()
	form.
		SetItemPadding(0).
		SetFieldBackgroundColor(tcell.ColorBlack).
		SetBorder(true).
		SetTitle("  Edit Configuration  ").
		SetBorderPadding(1,1,2,2)
	form.AddInputField("Version:", version, 40, nil, nil)

	list := tview.NewList()
	list.
		ShowSecondaryText(false).
		SetWrapAround(false).
		SetBorder(true).
		SetTitle("  Select features to add  ").
		SetBorderPadding(1,1,2,2)

	comment := tview.NewTextView().
		SetTextColor(tcell.ColorWhite).
		SetText("Helpful comment here").
		SetTextAlign(tview.AlignCenter)

	// create a closure to call on checkbox redraw - to allow us to set the comment
	newDrawFunc := func(p *tview.Box, name string) func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
		return func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
			if form.HasFocus() {
				item, _ := form.GetFocusedItemIndex()

				newLabel := noDescription
				oldLabel := strings.TrimSpace(comment.GetText(false))

				if item >= 1 && item <= len(file.OptionNames) {
					optionName := strings.TrimPrefix(file.OptionNames[item-1], "+ ")
					newLabel = crd.getDescription(optionName)
				} else {
					newLabel = noDescription
				}

				if oldLabel != newLabel {
					comment.SetText(newLabel)
					app.QueueEvent(tcell.NewEventKey(tcell.KeyCtrlA, 0, tcell.ModNone))
				}
			}
			return p.GetInnerRect()
		}
	}

	optionValues := make(map[string]bool)

	changedFunc := func(name string) func(bool) {
		return func(checked bool) {
			optionValues[strings.TrimPrefix(name, "+ ")] = checked
		}
	}
	// Set up the checkBoxes in the left hand form
	for i, name := range file.OptionNames {
		value, ok := file.GetField("spec/"+name+"/enabled")
		if ok {
			optionValues[name] = value == "true"

			form.AddCheckbox(name, value == "true", changedFunc(name))
			cb := form.GetFormItem(i+1).(*tview.Checkbox)
			cb.SetDrawFunc(newDrawFunc(cb.Box, name))
		}
	}

	focusForm := func() {
		app.SetFocus(form)
		footer.SetText(formHelp)
		form.SetBackgroundColor(tcell.ColorDarkBlue)
		list.SetBackgroundColor(tcell.ColorBlack)
		list.SetSelectedBackgroundColor(tcell.ColorBlack)
		list.SetSelectedTextColor(tcell.ColorWhite)
	}

	focusList := func() {
		app.SetFocus(list)
		footer.SetText(listHelp)
		list.SetBackgroundColor(tcell.ColorDarkBlue)
		list.SetSelectedBackgroundColor(tcell.ColorWhite)
		list.SetSelectedTextColor(tcell.ColorBlack)
		form.SetBackgroundColor(tcell.ColorBlack)
	}

	populateList := func() { }

	addFeature := func(feature string, row int) {
		form.AddCheckbox("+ "+feature, true, changedFunc(feature))
		optionValues[feature] = true
		file.OptionNames = append(file.OptionNames, feature)
		newAddable := make([]string, 0, 16)
		for _, a := range addable {
			if a != feature {
				newAddable = append(newAddable, a)
			}
		}
		addable = newAddable
		populateList()
		list.SetCurrentItem(row)
	}

	listChangedFunc := func(index int, mainText string) {
		name := strings.TrimSpace(mainText)
		descr := crd.getDescription(name)
		if descr == name { descr = noDescription }
		comment.SetText(descr)
	}

	populateList = func() {
		list.Clear()
		for i, a := range addable {
			added := func(feature string,row int) func() {
				return func() {
					addFeature(feature, row)
				}
			}
			list.AddItem(fmt.Sprintf("  %-50s  ", a), "", 0, added(a, i))
		}

		list.SetChangedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
			listChangedFunc(index, mainText)
		})
	}

	populateList()

	var modal *tview.Modal

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			if form.HasFocus() || list.HasFocus() {
				modal = confirmSave(app)
				app.SetRoot(modal, true)
				app.SetFocus(modal)
			} else if modal != nil && modal.HasFocus() {
				app.SetRoot(grid, true)
				app.SetFocus(grid)
				focusForm()
			}

			return nil

		} else if event.Key() == tcell.KeyTAB && form.HasFocus() {
			focusList()
			i := list.GetCurrentItem()
			text, _ := list.GetItemText(i)
			listChangedFunc(i, text)
			return nil

		} else if event.Key() == tcell.KeyTAB && list.HasFocus() {
			focusForm()
			return nil
		}

		return event
	})

	grid.AddItem(header, 0, 0, 1, 2, 0, 0, false)
	grid.AddItem(form, 1, 0, 1, 1, 0, 0, true)
	grid.AddItem(list, 1, 1, 1, 1, 0, 0, false)
	grid.AddItem(comment, 2, 0, 1, 2, 0, 0, false)
	grid.AddItem(footer, 3, 0, 1, 2, 0, 0, false)

	focusForm()

	if err := app.SetRoot(grid, true).SetFocus(grid).Run(); err != nil {
		panic(err)
	}

	if exitButton == "Yes" {
		for label, checked := range optionValues {
			lineNum, ok := file.EnabledFlags[label]
			if ok {
				f := strings.Split(file.Lines[lineNum], ":")
				file.Lines[lineNum] = f[0] + ": " + fmt.Sprintf("%v", checked)
			} else {
				file.Lines = append(file.Lines, "")
				file.Lines = append(file.Lines, "  "+strings.TrimPrefix(label, "+ ")+":")
				file.Lines = append(file.Lines, fmt.Sprintf("    enabled: %v", checked))
			}
		}

		lineNo, ok := file.LineNumByPath["spec/global/tag"]
		if ok {
			f := strings.SplitN(file.Lines[lineNo], ":", 2)
			version := strings.TrimSpace(form.GetFormItem(0).(*tview.InputField).GetText())
			file.Lines[lineNo] = f[0] + ": " + version
		}

		content := strings.TrimSpace(strings.Join(file.Lines, "\n")) + "\n"
		err = ioutil.WriteFile(fileName+".new", []byte(content), 0666)
		if err != nil {
			panic(err)
		}

		os.Remove(fileName + ".bak9")
		os.Rename(fileName+".bak8", fileName+".bak9")
		os.Rename(fileName+".bak7", fileName+".bak8")
		os.Rename(fileName+".bak6", fileName+".bak7")
		os.Rename(fileName+".bak5", fileName+".bak6")
		os.Rename(fileName+".bak4", fileName+".bak5")
		os.Rename(fileName+".bak3", fileName+".bak4")
		os.Rename(fileName+".bak2", fileName+".bak3")
		os.Rename(fileName+".bak", fileName+".bak2")
		os.Rename(fileName, fileName+".bak")
		os.Rename(fileName+".new", fileName)
	}
}
