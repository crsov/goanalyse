package cui

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	board_view, err := g.SetView("board", 0, 0, int(0.6*float32(maxX)), maxY-1)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
	}
	games_view, err := g.SetView("games", int(0.6*float32(maxX)), 0, int(0.4*float32(maxX)), maxY-1)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
	}

	fmt.Fprintln(board_view, maxX, maxY)
	fmt.Fprintln(games_view, maxX, maxY)

	return nil
}
