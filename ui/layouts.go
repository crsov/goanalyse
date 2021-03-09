package ui

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	if v, err := g.SetView("hello", 10, maxY-10, maxX-10, 10); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "Helloooooooooooooooooooooooooooooooooooooooooooooooo!")
	}
	return nil
}
