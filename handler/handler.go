package handler

import (
	"fmt"
	gateway "gateway/mqtt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Com struct {
	EquipmentType string      `json:"equipmentType"`
	CommandType   string      `json:"commandType"`
	Command       interface{} `json:"command"`
}

type Server struct {
	Address string
	Port    string
}

func Run() {
	server := Server{
		Address: os.Getenv("SERVER_ADDRESS"),
		Port:    os.Getenv("SERVER_PORT"),
	}
	r := gin.Default()

	r.POST("/command", sendCommand)

	log.Fatal(r.Run(server.Address + ":" + server.Port))
}

func sendCommand(c *gin.Context) {
	var command Com
	if err := c.BindJSON(&command); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	gateway.Pub("command/"+command.EquipmentType+"_"+command.CommandType, fmt.Sprint(command.Command))
}
