package controller

import (
	"bcc-geazy/config"
	"fmt"
	"net/http"
)

func GoogleLogin(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Client ID:", config.GoogleOAuthConfig.ClientID)
	url := config.GoogleOAuthConfig.AuthCodeURL("state-token")

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func GoogleCallback(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "http://localhost:8080", http.StatusTemporaryRedirect)
}
