package app

import (
	"bcc-geazy/internal/controller/rest"
	"bcc-geazy/internal/repository"
	"bcc-geazy/internal/usecase"
	"bcc-geazy/pkg/bcrypt"
	httpserver "bcc-geazy/pkg/gin"
	"bcc-geazy/pkg/jwt"
	"bcc-geazy/pkg/middleware"
	sql "bcc-geazy/pkg/mysql"
	"bcc-geazy/pkg/oauth"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/go-playground/validator/v10"
)

func Run() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error : %v", err)
	}

	db := sql.StartMySQL()
	app := httpserver.Start()

	jwtInit := jwt.NewJWT()
	bcryptService := bcrypt.NewBcrypt()
	oauthConfig := oauth.GoogleOAuthConfig()
	mw := middleware.NewMiddleware(jwtInit)
	validator := validator.New()

	repo := repository.NewRepository(db)

	uc := usecase.NewUsecase(jwtInit, bcryptService, oauthConfig, repo)
	v1 := rest.NewV1(mw, validator, uc)
	rest.NewRouter(app, v1)

	if err := app.Run(":" + os.Getenv("APP_PORT")); err != nil {
		log.Fatalf("Gagal start server : %s", err.Error())
	}
}
