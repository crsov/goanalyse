package main

import (
	"github.com/crsov/goanalyse/pkg/gui"
	"github.com/crsov/goanalyse/pkg/data"
)

func main() {
	gui.StartUi()
	data.LoadAccordingToCache("data.xml")
}
