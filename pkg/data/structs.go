package data

import "encoding/xml"

type roRoot struct {
	XMLNme  xml.Name `xml:"root"`
	RoGames []roGame
}
type roGame struct {
	Id     int    `xml:"id"`
	Black  string `xml:"black"`
	White  string `xml:"white"`
	Moves  string `xml:"moves"`
	Alt5   string `xml:"alt_5"`
	Rule   string `xml:"rule"`
	Winner string `xml:"winner"`
}
