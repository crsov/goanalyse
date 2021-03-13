package main

import (
	"fmt"
	"github.com/crsov/goanalyse/pkg/data"
)

func main() {
	fmt.Println((data.LoadAccordingToCache("data.xml")))
	//cui.StartUi()
}
