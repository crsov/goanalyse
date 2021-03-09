package ui

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

func layout(ui *gocui.Gui) error {
	maxX, maxY := ui.Size()

	if view, err := ui.SetView("hello", maxX/2-30, maxY/2, maxX/2+30, maxY/2+2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(view, "fskfjsdhflkjh")
	}
	return nil
}
