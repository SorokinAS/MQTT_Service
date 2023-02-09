package main

import (
	"fmt"
	syscfg "gateway/configuration"
	gateway "gateway/mqtt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	command   Com
	sysconfig = syscfg.GetEquip("equipment.xml")
)

type Com struct {
	EquipmentType string      `json:"equipmentType"`
	CommandType   string      `json:"commandType"`
	Command       interface{} `json:"command"`
}

func main() {
	go gateway.MQTT()

	r := gin.Default()

	r.POST("/command", sendCommand)

	r.Run(sysconfig.Server.Address + ":" + sysconfig.Server.Port)
}

func sendCommand(c *gin.Context) {

	if err := c.BindJSON(&command); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	gateway.Pub("command/"+command.EquipmentType+"_"+command.CommandType, fmt.Sprint(command.Command))
}
