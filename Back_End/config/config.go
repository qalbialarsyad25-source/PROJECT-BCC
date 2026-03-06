package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func sambungdatabase() {
	dsn := "root:@tcp(127.0.0.1:3306)/db_gizi?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Gagal koneksi ke database!")
	}

	fmt.Println("Koneksi Database Berhasil!")
	DB = database
}
