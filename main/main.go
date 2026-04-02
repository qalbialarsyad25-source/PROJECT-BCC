package main

import (
	"bcc-geazy/config"
	"bcc-geazy/internal/app"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

func main() {
	config.NewConfig()
	app.Run()

	r := gin.Default()
	r.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "https://geazy.vercel.app/"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	handler := c.Handler(mux)
	http.ListenAndServe(":8080", handler)

}
