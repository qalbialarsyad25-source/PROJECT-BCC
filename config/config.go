package config

import (
	"log"

	"github.com/joho/godotenv"
)

func SambungDatabase() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Tidak Valid .env File")
	}
}
