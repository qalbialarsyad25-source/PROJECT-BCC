package main

import (
	"bcc-geazy/config"
	"bcc-geazy/internal/app"
)

func main() {

	config.SambungDatabase()
	app.Run()
}
