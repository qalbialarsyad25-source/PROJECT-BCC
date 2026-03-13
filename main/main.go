package main

import (
	"bcc-geazy/config"
	"bcc-geazy/internal/app"
	controller "bcc-geazy/internal/auth"
	"net/http"
)

func main() {

	config.SambungDatabase()
	app.Run()

	http.HandleFunc("/auth/google", controller.GoogleLogin)
	http.HandleFunc("/auth/google/callback", controller.GoogleCallback)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}
}
