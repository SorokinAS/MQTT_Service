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

type Mqtt struct {
	Address string `xml:"address"`
	Port    string `xml:"port"`
}

type Db struct {
	Address string `xml:"address"`
	Port    string `xml:"port"`
	Urn     string `xml:"urn"`
	Timeout string `xml:"timeout"`
}

type Server struct {
	Address string `xml:"address"`
	Port    string `xml:"port"`
}

type Root struct {
	Server Server       `xml:"server"`
	Tsdb   Db           `xml:"tsdb"`
	Broker Mqtt         `xml:"broker"`
	Mes    Measurements `xml:"measurements"`
	Com    Commands     `xml:"commands"`
}

func GetEquip(filename string) Root {
	var equip Root
	xmlFile, err := os.Open(filename)
	if err != nil {
		log.Fatal("Invalid XML file, ", err)
	}
	defer xmlFile.Close()
	byteValue, _ := io.ReadAll(xmlFile)
	xml.Unmarshal(byteValue, &equip)
	return equip
}
