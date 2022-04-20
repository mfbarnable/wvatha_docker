package gui

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func GetHomePage() ui.Drawable {
	l := widgets.NewList()
	l.Rows = []string{
		"[1] List running instances",
		"[2] Stop",
	}
	l.TextStyle = ui.NewStyle(ui.ColorYellow)
	l.WrapText = false
	l.SetRect(0, 0, 25, 8)
	return l

}
