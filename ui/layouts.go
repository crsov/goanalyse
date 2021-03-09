package ui

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	if v, err := g.SetView("hello", maxX/2-7, maxY/2-2, maxX/2+2, maxY/2+7); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "Helloooooooooooooooooooooooooooooooooooooooooooooooo!")
	}
	return nil
}
