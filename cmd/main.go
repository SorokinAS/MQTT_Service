package main

import (
	"gateway/handler"
	"gateway/mqtt"
	"log"
	"path/filepath"

	"github.com/joho/godotenv"
)

func init() {
	configPath := filepath.Join("./config", ".env")
	if err := godotenv.Load(configPath); err != nil {
		log.Fatal("Failed to read .env", err)
	}
}

func main() {
	go mqtt.Run()
	handler.Run()
}
