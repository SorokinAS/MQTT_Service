package main

import (
	"gateway/handler"
	"gateway/mqtt"
	"log"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load("./config/.env"); err != nil {
		log.Fatal("Failed to read .env", err)
	}
}

func main() {
	go mqtt.Run()
	handler.Run()
}
