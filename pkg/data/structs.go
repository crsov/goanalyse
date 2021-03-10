package data

type roRoot struct {
	RoGames []roGame `xml:"root>game"`
}
type roGame struct {
	Id     int    `xml:"game>id"`
	Black  string `xml:"game>black"`
	White  string `xml:"game>white"`
	Moves  string `xml:"game>moves"`
	Alt5   string `xml:"game>alt_5"`
	Rule   string `xml:"game>rule"`
	Winner string `xml:"game>winner"`
}
