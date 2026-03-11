package sql

import (
	"bcc-geazy/Internal/entity"
	"fmt"

	"gorm.io/gorm"
)

var DB *gorm.DB

func Mysql() {
	err := DB.AutoMigrate(
		&entity.User{},
		&entity.Anak{},
		&entity.Makanan{},
		&entity.LogMakanan{},
		&entity.Dokter{},
		&entity.Konsultasi{},
		&entity.Informasi{},
	)

	if err != nil {
		panic("Gagal" + err.Error())
	}
	fmt.Println("Database berhasil migration")
}
