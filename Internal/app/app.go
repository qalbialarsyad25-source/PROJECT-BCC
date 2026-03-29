package app

import (
	websocket "bcc-geazy/internal/controller/delivery"
	"bcc-geazy/internal/controller/rest"
	"bcc-geazy/internal/repository"
	"bcc-geazy/internal/seeder"
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

	repo := repository.NewRepository(db)

	seeder.InfoMakanan(db)
	app := httpserver.Start()

	wsManager := websocket.NewWSManager()

	jwtInit := jwt.NewJWT()
	bcryptService := bcrypt.NewBcrypt()
	oauthConfig := oauth.GoogleOAuthConfig()
	mw := middleware.NewMiddleware(jwtInit)
	validator := validator.New()

	uc := usecase.NewUsecase(jwtInit, bcryptService, oauthConfig, repo, wsManager)
	v1 := rest.NewV1(mw, validator, uc, wsManager)
	rest.NewRouter(app, v1, wsManager)

	if err := app.Run(":" + os.Getenv("APP_PORT")); err != nil {
		log.Fatalf("Gagal start server : %s", err.Error())
	}
}
