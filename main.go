package main

import (
	"encoding/json"
	"fmt"
	gateway "gateway/mqtt"
	"log"
	"net/http"
)

type Com struct {
	EquipmentType string      `json:"equipmentType"`
	CommandType   string      `json:"commandType"`
	Command       interface{} `json:"command"`
}

func main() {
	go gateway.MQTTSub()

	http.HandleFunc("/command", func(w http.ResponseWriter, r *http.Request) {
		var c Com
		err := json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			log.Fatal(err)
		}
		gateway.Pub("command/"+c.EquipmentType+"_"+c.CommandType, fmt.Sprint(c.Command))
	})

	log.Print("Connecting...")
	http.ListenAndServe(":8604", nil)
}
