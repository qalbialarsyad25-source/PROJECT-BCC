package models

import (
	"time"
)

type User struct {
	ID         uint   `gorm:"primaryKey"`
	GoogleID   string `gorm:"uniqueIndex;type:varchar(225); not null"`
	Nama       string `gorm:"type:varchar(225); not null"`
	Email      string `gorm:"uniqueIndex;type:varchar(225); not null"`
	Anak       []Anak `gorm:"foreignKey:UserID"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Userkonsul []Konsultasi `gorm:"foreignKey:UserID"`
	LogMakanan []LogMakanan `gorm:"foreignKey:UserID"`
}

type Anak struct {
	ID         uint `gorm:"primaryKey"`
	UserID     uint
	Loganak    []LogMakanan `gorm:"foreignKey:AnakID"`
	Makanan    []Makanan    `gorm:"foreignKey:AnakID"`
	NamaAnak   string       `gorm:"type:varchar(225); not null"`
	TinggiAnak float64      `gorm:"type:decimal(10,2); not null"`
	BBAnak     float64      `gorm:"type:float; not null"`
	GenderAnak string       `gorm:"type:varchar(100); not null"`
}

type Makanan struct {
	ID           uint   `gorm:"primaryKey"`
	NamaMakanan  string `gorm:"type:varchar(100); not null"`
	AnakID       uint
	LogMakananID uint
	EnergiFood   float64 `gorm:"type:float"`
	Protein      float64 `gorm:"type:float"`
	Lemak        float64 `gorm:"type:float"`
	Karbo        float64 `gorm:"type:float"`
}

type Dokter struct {
	ID           uint         `gorm:"primaryKey"`
	NamaDokter   string       `gorm:"type:varchar(225); not null"`
	DokterKonsul []Konsultasi `gorm:"foreignKey:DokterID"`
}

type LogMakanan struct {
	ID          uint `gorm:"primaryKey"`
	UserID      uint
	Makanan     []Makanan `gorm:"foreignKey:LogMakananID"`
	AnakID      uint
	Gram        float64 `gorm:"type:float"`
	TotalKalori float64 `gorm:"type:float"`
	WaktuMakan  time.Time
}

type Konsultasi struct {
	ID         uint `gorm:"primaryKey"`
	UserID     uint
	DokterID   uint
	Pesan      string `gorm:"type:longtext"`
	WaktuPesan time.Time
}

type Informasi struct {
	ID        uint   `gorm:"primaryKey"`
	Ringkasan string `gorm:"type:longtext; not null"`
	Judul     string `gorm:"type:varchar(225); not null"`
}
