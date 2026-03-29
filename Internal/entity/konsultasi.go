package entity

import (
	"time"

	"github.com/google/uuid"
)

type Konsultasi struct {
	Id         uuid.UUID `gorm:"type:char(36);primaryKey"`
	UserID     uuid.UUID `gorm:"type:char(36);not null;constraint;OnDelete:CASCADE"`
	DokterID   uuid.UUID `gorm:"type:char(36);not null;constraint;OnDelete:CASCADE"`
	SenderID   uuid.UUID `gorm:"type:char(36);not null"`
	Pesan      string    `gorm:"type:longtext"`
	Dibaca     bool      `gorm:"default:false"`
	WaktuPesan time.Time `gorm:"type:timestamp;not null;autoCreateTime"`
}
