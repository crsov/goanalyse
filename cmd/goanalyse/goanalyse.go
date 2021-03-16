package main

import (
	"github.com/crsov/goanalyse/pkg/data"
	"github.com/crsov/goanalyse/pkg/gui"
)

func main() {
	gui.StartUi()
	data.LoadAccordingToCache("data.xml")
}
