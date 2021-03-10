package data

import "encoding/xml"

type roRoot struct {
	XMLNme  xml.Name `xml:"root"`
	RoGames []roGame
}
type roGame struct {
	Id     int    `xml:"root>game>id"`
	Black  string `xml:"root>game>black"`
	White  string `xml:"root>game>white"`
	Moves  string `xml:"root>game>moves"`
	Alt5   string `xml:"root>game>alt_5"`
	Rule   string `xml:"root>game>rule"`
	Winner string `xml:"root>game>winner"`
}
