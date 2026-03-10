package main

import (
	"bcc-geazy/config"
	"fmt"
	"net/http"
)

func main() {
	config.SambungDatabase()
	config.InitDB()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Berjalan")
	})

	fmt.Println("Server berjalan 8080")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}
}
