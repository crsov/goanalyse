package main

import (
	"github.com/crsov/goanalyse/pkg/load_data"
)

func main() {
	roBaseName := "downloaded_renjuoffline_base.xml"
	rnBaseName := "downloaded_renjunet_base.rif"
	//gui.StartUi()
	load_data.DlRoBase(roBaseName)
	load_data.DlRnBase(rnBaseName)
	//load_data.LoadAccordingToCache(roBaseName)
}
