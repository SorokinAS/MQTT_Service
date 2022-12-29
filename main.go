package main

import (
	"encoding/json"
	"fmt"
	gateway "gateway/mqtt"
	"log"
	"net/http"
)

var (
	c    Com
	port = "8604"
)

type Com struct {
	EquipmentType string      `json:"equipmentType"`
	CommandType   string      `json:"commandType"`
	Command       interface{} `json:"command"`
}

func main() {
	go gateway.MQTT()

	http.HandleFunc("/command", func(w http.ResponseWriter, r *http.Request) {
		err := json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			log.Fatal(err)
		}
		gateway.Pub("command/"+c.EquipmentType+"_"+c.CommandType, fmt.Sprint(c.Command))
	})

	log.Print("Connecting to 127.0.0.1:8604")
	http.ListenAndServe(":"+port, nil)
}
