package sql

import (
	"bcc-geazy/internal/entity"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func StartMySQL() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Gagal koneksi : %s", err.Error())
	}

	Migrate(db)

	return db
}

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&entity.User{},
		&entity.Anak{},
		&entity.Makanan{},
		&entity.Log{},
		&entity.LogMakanan{},
		&entity.Dokter{},
		&entity.Konsultasi{},
		&entity.Informasi{},
		&entity.LogInformasi{},
	)

	if err != nil {
		panic("Gagal" + err.Error())
	}
	fmt.Println("Database berhasil migration")
}
