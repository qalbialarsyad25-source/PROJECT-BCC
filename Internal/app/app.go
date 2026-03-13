package app

import (
	"bcc-geazy/internal/controller/rest"
	"bcc-geazy/internal/repository"
	"bcc-geazy/internal/usecase"
	httpserver "bcc-geazy/pkg/gin"
	sql "bcc-geazy/pkg/mysql"
	"log"
	"os"
)

func Run() {
	db := sql.StartMySQL()
	app := httpserver.Start()

	repo := repository.NewRepository(db)
	uc := usecase.NewUsecase(repo)
	v1 := rest.NewV1(uc)

	rest.NewRouter(app, v1)

	if err := app.Run(":" + os.Getenv("APP_PORT")); err != nil {
		log.Fatalf("Gagal start server : %s", err.Error())
	}
}
