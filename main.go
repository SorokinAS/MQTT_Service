package main

import (
	"encoding/json"
	"log"

	gateway "gateway/mqtt"
	"net/http"
)

type Com struct {
	EquipmentType string      `json:"equipmentType"`
	CommandType   string      `json:"commandType"`
	Command       interface{} `json:"command"`
}

func main() {
	gateway.MQTTSub()

	http.HandleFunc("/command", func(w http.ResponseWriter, r *http.Request) {
		var c Com
		err := json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			log.Fatal(err)
		}
		log.Print("Получено сообщение: ")
		log.Print("command/"+c.EquipmentType+"_"+c.CommandType, c.Command.(string))
		// gateway.Pub("command/"+c.EquipmentType+"_"+c.CommandType, c.Command.(string))
	})

	log.Print("Connecting...")
	http.ListenAndServe(":8604", nil)
}
