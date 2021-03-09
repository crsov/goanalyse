package data

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

func Xml() {

	xmlFile, err := os.Open("data.xml")
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully Opened xml")

	defer xmlFile.Close()

	inBytes, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		panic(err)
	}
	var roGames roGame
	xml.Unmarshal(inBytes, &roGames)
	fmt.Println(roGames)

}
