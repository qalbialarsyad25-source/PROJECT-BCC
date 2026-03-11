package main

import (
	"bcc-geazy/config"
	mysql "bcc-geazy/pkg/mysql"
	"fmt"
	"net/http"
)

func main() {
	config.SambungDatabase()
	mysql.Mysql()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Berjalan dong, Qall selanjutnya apa?")
	})

	fmt.Println("Server berjalan 8080")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}
}
