package mqtt

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var (
	values    []string
	mes       Measurement
	equipment string
	mag       float64
	ang       float64
	url       = "http://127.0.0.1:8081/save"
	res       map[string]interface{}
)

type Measurement struct {
	Topic     string  `json:"topic"`
	Equipment string  `json:"equipment"`
	Mag       float64 `json:"mag"`
	Ang       float64 `json:"ang"`
	TimeStamp string  `json:"timestamp"`
}

func (mes *Measurement) newMeasurement(topic *string, equip *string, mag *float64, ang *float64) {
	(*mes).Topic = *topic
	(*mes).Equipment = *equip
	(*mes).Mag = *mag
	(*mes).Ang = *ang
	(*mes).TimeStamp = time.Now().Format("2006-01-02T15:04:05Z07:00")

}

func sendMeasurements(msg *string, topic string) {
	values = strings.Split(*msg, ",")
	mag, _ = strconv.ParseFloat(values[0], 32)
	if len(values) > 1 {
		ang, _ = strconv.ParseFloat(values[1], 32)
	} else {
		ang = 0.0
	}

	equipment = getEquipment(&topic)
	mes.newMeasurement(&topic, &equipment, &mag, &ang)
	message, err := json.Marshal(mes)
	if err != nil {
		log.Fatal("Invalid Marshal message with measurement: ", err)
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(message))
	if err != nil {
		log.Fatal("Invalid sending POST request with measurement: ", err)
	}
	log.Println("Sended measurement: ", string(message))

	json.NewDecoder(resp.Body).Decode(&res)
	//log.Println(res["json"]) uncomment if you want to see response
}

func getEquipment(topic *string) string {
	return strings.Split(*topic, "/")[1]
}
