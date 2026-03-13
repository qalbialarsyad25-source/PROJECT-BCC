package controller

import (
	"bcc-geazy/config"
	"bcc-geazy/internal/entity"
	"bcc-geazy/internal/model"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
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

	var user entity.User
	var userresponse model.UserResponse

	err = config.DB.Where("google_id = ?", googleUser.Id).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			user = entity.User{
				Id:             uuid.New(),
				GoogleID:       googleUser.Id,
				Nama:           googleUser.Nama,
				Email:          googleUser.Email,
				DibuatPada:     time.Now(),
				DiperbaruiPada: time.Now(),
			}
			config.DB.Create(&user)
		}
	}

	fmt.Println("User ID:", googleUser.Id)
	fmt.Println("Email:", googleUser.Email)
	fmt.Println("Nama:", googleUser.Nama)
	fmt.Println("Dibuat pada:", userresponse.DibuatPada)
	fmt.Println("DIperbarui pada:", userresponse.DiperbaruiPada)

	tokenJWT, err := GenerateJWT(user.Id.String(), user.Email)
	if err != nil {
		http.Error(w, "Gagal membuat token", http.StatusInternalServerError)
		return
	}

	response := model.LoginResponse{
		Token: tokenJWT,
		User: model.UserResponse{
			Id:             user.Id.String(),
			Nama:           user.Nama,
			Email:          user.Email,
			DibuatPada:     user.DibuatPada.Format("2006-01-02 15:04:05"),
			DiperbaruiPada: user.DiperbaruiPada.Format("2006-01-02 15:04:05"),
		},
	}

	json.NewEncoder(w).Encode(response)
}
