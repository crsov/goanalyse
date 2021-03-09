package cui

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	boardView, err := g.SetView("board", 0, 0, int(0.6*float32(maxX)), maxY-1)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
	}
	gamesView, err := g.SetView("games", int(0.6*float32(maxX)), 0, int(0.4*float32(maxX)), maxY-1)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
	}

	fmt.Fprintln(gamesView, maxX, maxY)
	fmt.Fprintln(boardView, maxX, maxY)

	return nil
}
