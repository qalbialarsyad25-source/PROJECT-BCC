package config

import (
	"bcc-geazy/models"
	"fmt"
)

func InitDB() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Anak{},
		&models.Makanan{},
		&models.LogMakanan{},
		&models.Dokter{},
		&models.Konsultasi{},
		&models.Informasi{},
	)

	if err != nil {
		panic("Gagal" + err.Error())
	}
	fmt.Println("Database berhasil migration")
}
