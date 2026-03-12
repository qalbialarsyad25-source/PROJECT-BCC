package main

import (
	"bcc-geazy/config"
	"bcc-geazy/internal/controller"
	mysql "bcc-geazy/pkg/mysql"
	"fmt"
	"net/http"
)

func main() {
	config.SambungDatabase()
	mysql.StartMySQL()
	config.GoogleOuath()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Berjalan dong, Qall selanjutnya apa?")
	})

	http.HandleFunc("/auth/google", controller.GoogleLogin)
	http.HandleFunc("/auth/google/callback", controller.GoogleCallback)

	fmt.Println("Server berjalan 8080")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}
}
