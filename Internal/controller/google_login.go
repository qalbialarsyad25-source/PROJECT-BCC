package controller

import (
	"bcc-geazy/config"
	"bcc-geazy/internal/model"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func GoogleLogin(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Client ID:", config.GoogleOAuthConfig.ClientID)
	url := config.GoogleOAuthConfig.AuthCodeURL("state-token")

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func GoogleCallback(w http.ResponseWriter, r *http.Request) {

	code := r.URL.Query().Get("code")

	fmt.Println("Code dari Google:", code)

	token, err := config.GoogleOAuthConfig.Exchange(context.Background(), code)
	if err != nil {
		http.Error(w, "Token error", http.StatusInternalServerError)
		return
	}

	fmt.Println("Access Token:", token.AccessToken)

	client := config.GoogleOAuthConfig.Client(context.Background(), token)

	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		http.Error(w, "Gagal ambil user info", http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	var googleUser model.GoogleUser

	json.NewDecoder(resp.Body).Decode(&googleUser)

	fmt.Println("User ID:", googleUser.Id)
	fmt.Println("Email:", googleUser.Email)
	fmt.Println("Nama:", googleUser.Nama)
	http.Redirect(w, r, "http://localhost:8080", http.StatusTemporaryRedirect)
}
