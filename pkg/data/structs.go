package data

type roGame struct {
	id     int    `xml:"id"`
	black  string `xml:"black"`
	white  string `xml:"white"`
	winner string `xml:"winner"`
	moves  string `xml:"moves"`
	rule   string `xml:"rule"`
}

type rnGame struct {
}

type rnPlayer struct {
	id      int
	name    string
	country int
}

type rnRule struct {
	id   int
	name string
}

type rnCity struct {
	id      int
	country int
}

type rnOpenning struct {
	id   int
	abbr string
}

type rnCountry struct {
	id   int
	name string
	abbr string
}
