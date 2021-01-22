package ui

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"../t8cop"
)

type CheckboxPlus struct {
	tview.Checkbox
	comment *tview.TextView
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
	if cb.comment != nil {
		cb.comment.SetText(mesg)
	}
}

func (cb *CheckboxPlus) SetCommentLine(comment *tview.TextView) {
	cb.comment = comment
}
