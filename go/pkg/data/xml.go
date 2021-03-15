package data

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

func ParseXml(file string) ro {

	xmlFile, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully Opened xml")

	defer xmlFile.Close()

	inBytes, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		panic(err)
	}
	var ro ro
	err = xml.Unmarshal(inBytes, &ro)
	if err != nil {
		panic(err)
	}

	return ro
}
