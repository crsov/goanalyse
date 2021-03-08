package ui

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
)

func StartUi() {
	ui, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}

	defer ui.Close()

	ui.SetManagerFunc(layout)

	if err := ui.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := ui.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func layout(ui *gocui.Gui) error {
	maxX, maxY := ui.Size()

	if view, err := ui.SetView("hello", maxX, maxY, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(view, "ugashlgkjsdhglksjdghlaksdjfhasdlkjghaslfkgjasfhlkasfjghlskjghldsfkjadsfhlakdfhgasldjgasldkajladkdjfasglkfjgsdahlkfjgasglkjasflakjgal")
	}
	return nil
}

func quit(ui *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
