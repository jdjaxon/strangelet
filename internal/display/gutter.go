package display

import (
	"github.com/Kyrasuum/cview"
	"github.com/gdamore/tcell/v2"
)

var (
	gutterBGColor tcell.Color
	gutterFGColor tcell.Color
)

// Gutter display struct
type gutter struct {
	*cview.TextView
}

func NewGutter(subFlex *cview.Flex) (gutt *gutter) {
	gutt = &gutter{}

	if gutterBGColor == 0 {
		gutterBGColor = tcell.NewRGBColor(40, 40, 40)
	}
	if gutterFGColor == 0 {
		gutterFGColor = tcell.NewRGBColor(170, 170, 170)
	}
	gutt.TextView = cview.NewTextView()
	gutt.TextView.SetTextAlign(cview.AlignRight)
	gutt.TextView.SetRegions(true)
	gutt.TextView.Box.SetBackgroundColor(gutterBGColor)
	gutt.TextView.SetTextColor(gutterFGColor)
	gutt.TextView.SetHighlightBackgroundColor(bufferBGColor)
	gutt.TextView.SetHighlightForegroundColor(gutterFGColor)
	gutt.TextView.SetScrollBarVisibility(cview.ScrollBarNever)

	subFlex.AddItem(gutt, 2, 1, false)

	return gutt
}
