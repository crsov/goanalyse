package load_data

import "encoding/xml"

// Main Format
type mainf struct {
	id     int
	black  string
	white  string
	moves  string
	rule   string
	winner string
	alt5   string //?
}

type ro struct {
	RoGames []roGame `xml:"game"`
}
type roGame struct {
	Id     string `xml:"id"`
	Black  string `xml:"black"`
	White  string `xml:"white"`
	Moves  string `xml:"moves"`
	Alt5   string `xml:"alt_5"`
	Rule   string `xml:"rule"`
	Winner string `xml:"winner"`
}

type rn struct {
	XMLName   xml.Name `xml:"database"`
	Version   string   `xml:"version,attr"`
	Date      string   `xml:"date,attr"`
	Countries struct {
		Country []struct {
			ID   string `xml:"id,attr"`
			Name string `xml:"name,attr"`
			Abbr string `xml:"abbr,attr"`
		} `xml:"country"`
	} `xml:"countries"`
	Cities struct {
		City []struct {
			ID      string `xml:"id,attr"`
			Country string `xml:"country,attr"`
			Name    string `xml:"name,attr"`
		} `xml:"city"`
	} `xml:"cities"`
	Months struct {
		Month []struct {
			ID   string `xml:"id,attr"`
			Name string `xml:"name,attr"`
		} `xml:"month"`
	} `xml:"months"`
	Rules struct {
		Rule []struct {
			ID   string `xml:"id,attr"`
			Name string `xml:"name,attr"`
			Info string `xml:"info"`
		} `xml:"rule"`
	} `xml:"rules"`
	Openings struct {
		Opening []struct {
			ID   string `xml:"id,attr"`
			Abbr string `xml:"abbr,attr"`
			Name string `xml:"name,attr"`
		} `xml:"opening"`
	} `xml:"openings"`
	Players struct {
		Player []struct {
			ID      string `xml:"id,attr"`
			Name    string `xml:"name,attr"`
			Surname string `xml:"surname,attr"`
			Country string `xml:"country,attr"`
			City    string `xml:"city,attr"`
		} `xml:"player"`
	} `xml:"players"`
	Tournaments struct {
		Tournament []struct {
			ID      string `xml:"id,attr"`
			Name    string `xml:"name,attr"`
			Country string `xml:"country,attr"`
			City    string `xml:"city,attr"`
			Year    string `xml:"year,attr"`
			Month   string `xml:"month,attr"`
			Start   string `xml:"start,attr"`
			End     string `xml:"end,attr"`
			Rule    string `xml:"rule,attr"`
			Rated   string `xml:"rated,attr"`
		} `xml:"tournament"`
	} `xml:"tournaments"`
	Games struct {
		Game []struct {
			ID         string `xml:"id,attr"`
			Publisher  string `xml:"publisher,attr"`
			Tournament string `xml:"tournament,attr"`
			Round      string `xml:"round,attr"`
			Rule       string `xml:"rule,attr"`
			Black      string `xml:"black,attr"`
			White      string `xml:"white,attr"`
			Bresult    string `xml:"bresult,attr"`
			Opening    string `xml:"opening,attr"`
			Alt        string `xml:"alt,attr"`
			Swap       string `xml:"swap,attr"`
			Move       string `xml:"move"`
		} `xml:"game"`
	} `xml:"games"`
}
