package entity

import (
	"time"

	"github.com/google/uuid"
)

type Anak struct {
	Id              uuid.UUID `gorm:"type:char(36);primaryKey"`
	UserID          uuid.UUID `gorm:"type:char(36);not null;constraint;OnDelete:CASCADE"`
	LogMakanan      []Log     `gorm:"foreignKey:AnakID"`
	Nama            string    `gorm:"type:varchar(225); not null"`
	TanggalLahir    int       `gorm:"type:int;not null"`
	BulanLahir      string    `gorm:"type:varchar(100); not null"`
	TahunLahir      int       `gorm:"type:int;not null"`
	Umur            string
	Tinggi          float64   `gorm:"type:decimal(10,2); not null"`
	BeratBadan      float64   `gorm:"type:decimal(10,2); not null"`
	Gender          string    `gorm:"type:varchar(10); not null"`
	AnakKe          int       `gorm:"type:int; not null"`
	LingkarLengan   float64   `gorm:"type:decimal(10,2); not null"`
	LingkarKepala   float64   `gorm:"type:decimal(10,2); not null"`
	GolonganDarah   string    `gorm:"type:varchar(10); not null"`
	Alergi          string    `gorm:"type:varchar(100); not null"`
	RiwayatPenyakit string    `gorm:"type:varchar(225); not null"`
	DibuatPada      time.Time `gorm:"type:timestamp; not null; autoCreateTime"`
	BMI             float64
	StatusGizi      string
}
