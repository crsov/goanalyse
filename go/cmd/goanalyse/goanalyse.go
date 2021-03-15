package main

import (
	"github.com/crsov/goanalyse/go/pkg/data"
	"github.com/crsov/goanalyse/go/pkg/gui"
)

func main() {
	gui.StartUi()
	data.LoadAccordingToCache("data.xml")
}
