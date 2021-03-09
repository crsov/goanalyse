package ui

import (
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

func quit(ui *gocui.Gui, view *gocui.View) error {
	return gocui.ErrQuit
}
