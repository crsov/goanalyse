package data

import (
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

	//decoder := charmap.Windows1252.NewDecoder()
	//reader := decoder.Reader(xmlFile)
	inBytes, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(inBytes))
	// var roGames roGame
	// xml.Unmarshal(inBytes, &roGames)
	// fmt.Println(roGames)
}
