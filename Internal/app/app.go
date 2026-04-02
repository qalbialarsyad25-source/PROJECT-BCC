package app

import (
	websocket "bcc-geazy/internal/controller/delivery"
	"bcc-geazy/internal/controller/rest"
	"bcc-geazy/internal/jadwal"
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
	"time"

	"github.com/gin-contrib/cors"

	"github.com/go-playground/validator/v10"
)

func Run() {
	db := sql.StartMySQL()

	repo := repository.NewRepository(db)

	seeder.InfoMakanan(db)
	seeder.SeedAdmin(repo.UserRepository)
	app := httpserver.Start()

	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
			"https://geazy.vercel.app",
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	wsManager := websocket.NewWSManager()

	jwtInit := jwt.NewJWT()
	bcryptService := bcrypt.NewBcrypt()
	oauthConfig := oauth.GoogleOAuthConfig()
	mw := middleware.NewMiddleware(jwtInit)
	validator := validator.New()

	uc := usecase.NewUsecase(jwtInit, bcryptService, oauthConfig, repo, wsManager)
	v1 := rest.NewV1(mw, validator, uc, wsManager)
	rest.NewRouter(app, v1, wsManager)

	jadwal.StartCron(
		uc.NotifikasiUsecase,
		repo.UserRepository,
		repo.AnakRepository,
	)

	if err := app.Run(":" + os.Getenv("APP_PORT")); err != nil {
		log.Fatalf("Gagal start server : %s", err.Error())
	}
}
