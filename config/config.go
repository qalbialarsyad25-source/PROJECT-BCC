package config

import (
	sql "bcc-geazy/pkg/mysql"
	"log"

	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

var DB *gorm.DB

func SambungDatabase() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Tidak Valid .env File")
	}

	DB = sql.StartMySQL()
}
