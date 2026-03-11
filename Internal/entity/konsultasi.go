package entity

import (
	"time"
)

type Konsultasi struct {
	ID         uint `gorm:"primaryKey"`
	UserID     uint
	DokterID   uint
	Pesan      string `gorm:"type:longtext"`
	WaktuPesan time.Time
}
