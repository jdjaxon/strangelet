package display

import (
	"github.com/Kyrasuum/cview"
	"github.com/gdamore/tcell/v2"
)

var (
	statusBarBGColor tcell.Color
	statusBarFGColor tcell.Color
)

type statusBar struct {
	statusbar *cview.TextView
}

func NewStatusBar(flex *cview.Flex) (stat *statusBar) {
	stat = &statusBar{}

	if statusBarBGColor == 0 {
		statusBarBGColor = tcell.NewRGBColor(160, 160, 160)
	}
	if statusBarFGColor == 0 {
		statusBarFGColor = tcell.NewRGBColor(20, 20, 20)
	}

	stat.statusbar = cview.NewTextView()
	stat.statusbar.SetTextAlign(cview.AlignLeft)
	stat.statusbar.SetText("Status Bar")
	stat.statusbar.Box.SetBackgroundColor(statusBarBGColor)
	stat.statusbar.SetTextColor(statusBarFGColor)

	flex.AddItem(stat.statusbar, 1, 1, false)

	return stat
}
