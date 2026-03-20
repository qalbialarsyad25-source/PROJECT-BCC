package main

import (
	"bcc-geazy/config"
	"bcc-geazy/internal/app"
)

func main() {

	app.Run()
	config.SambungDatabase()

}
