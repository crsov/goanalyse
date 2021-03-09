package data

import (
	"encoding/xml"
)

type roRoot struct {
	XMLName xml.Name `xml:"root"`
	root    []roGame
}

type roGame struct {
	XMLName         xml.Name `xml:"game"`
	id              int      `xml:"id"`
	creation_time   int      `xml:"creation_time"`
	black           string   `xml:"black"`
	white           string   `xml:"white"`
	turn_color      string   `xml:"turn_color"`
	board           string   `xml:"board"`
	moves           string   `xml:"moves"`
	alt5            string   `xml:"alt_5"`
	proposition     string   `xml:"proposition"`
	proposer        string   `xml:"proposer"`
	status          string   `xml:"status"`
	rule            string   `xml:"rule"`
	time            int      `xml:"time"`
	tpm             int      `xml:"tpm"`
	start_time      int      `xml:"start_time"`
	end_time        int      `xml:"end_time"`
	winner          string   `xml:"winner"`
	winby           string   `xml:"winby"`
	time_left_black int      `xml:"time_left_black"`
	time_left_white int      `xml:"time_left_white"`
	tid             string   `xml:"tid"`
}
