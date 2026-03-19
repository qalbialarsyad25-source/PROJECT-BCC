package app

import (
	"bcc-geazy/internal/controller/rest"
	"bcc-geazy/internal/repository"
	"bcc-geazy/internal/usecase"
	"bcc-geazy/pkg/bcrypt"
	httpserver "bcc-geazy/pkg/gin"
	"bcc-geazy/pkg/jwt"
	sql "bcc-geazy/pkg/mysql"
	"bcc-geazy/pkg/oauth"
	"log"
	"os"
)

func Run() {
	db := sql.StartMySQL()
	app := httpserver.Start()

	repo := repository.NewRepository(db)

	jwtService := *jwt.NewJWT()
	bcryptService := bcrypt.NewBcrypt()
	oauthConfig := oauth.GoogleOAuthConfig()

	uc := usecase.NewUsecase(jwtService, bcryptService, oauthConfig, repo)
	v1 := rest.NewV1(uc)
	rest.NewRouter(app, v1)

	if err := app.Run(":" + os.Getenv("APP_PORT")); err != nil {
		log.Fatalf("Gagal start server : %s", err.Error())
	}
}
