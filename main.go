package main

import (
	"encoding/json"
	"fmt"
	gateway "gateway/mqtt"
	"log"
	"net/http"
	"os"
)

var (
	c    Com
	port = os.Getenv("SERVER_PORT")
)

type Com struct {
	EquipmentType string      `json:"equipmentType"`
	CommandType   string      `json:"commandType"`
	Command       interface{} `json:"command"`
}

func main() {
	go gateway.MQTTSub()

	http.HandleFunc("/command", func(w http.ResponseWriter, r *http.Request) {
		err := json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			log.Fatal(err)
		}
		gateway.Pub("command/"+c.EquipmentType+"_"+c.CommandType, fmt.Sprint(c.Command))
	})

	log.Print("Connecting...")
	if port == "" {
		log.Fatal("Empty ENV variable 'SERVER_PORT'")
	}
	http.ListenAndServe(":"+port, nil)
}
