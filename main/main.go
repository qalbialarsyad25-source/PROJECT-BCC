package main

import (
	"bcc-geazy/config"
	"bcc-geazy/internal/app"
	"bcc-geazy/internal/usecase"
	"net/http"
)

func main() {

	config.SambungDatabase()
	app.Run()

	http.HandleFunc("/auth/google", usecase.GoogleLogin)
	http.HandleFunc("/auth/google/callback", usecase.GoogleCallback)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}
}
