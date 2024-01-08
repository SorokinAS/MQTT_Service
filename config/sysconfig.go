package configuration

import (
	"encoding/xml"
	"io"
	"log"
	"os"
)

type Measurements struct {
	Value []Value `xml:"value"`
}

type Commands struct {
	Actions []Action `xml:"action"`
}

type Value struct {
	Topic     string  `xml:"topic,attr"`
	Equipment string  `xml:"equipment,attr"`
	Mag       float64 `xml:"mag,attr"`
	Ang       float32 `xml:"ang,attr"`
}

type Action struct {
	Topic     string `xml:"topic,attr"`
	Equipment string `xml:"equipment,attr"`
}

type Root struct {
	Mes Measurements `xml:"measurements"`
	Com Commands     `xml:"commands"`
}

func GetEquip(filename string) Root {
	var equip Root
	xmlFile, err := os.Open(filename)
	if err != nil {
		log.Fatal("Invalid XML file, ", err)
	}
	defer xmlFile.Close()
	byteValue, _ := io.ReadAll(xmlFile)
	err = xml.Unmarshal(byteValue, &equip)
	if err != nil {
		log.Fatal("Failed unmarshall ", err)
	}
	return equip
}
