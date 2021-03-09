package data

import (
	"encoding/xml"
	"fmt"
	"os"
)

func Xml() {

	f, err := os.Open("data.xml")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	decoder := xml.NewDecoder(f)

	roGames := make([]roGame, 0)
	for {
		tok, err := decoder.Token()
		if err != nil {
			panic(err)
		}
		if tok == nil {
			break
		}
		switch tokType := tok.(type) {
		case xml.StartElement:
			if tokType.Name.Local == "game" {
				// Декодирование элемента в структуру
				var g roGame
				decoder.DecodeElement(&g, &tokType)
				roGames = append(roGames, g)
			}
		}
	}
	fmt.Println(roGames)
}
