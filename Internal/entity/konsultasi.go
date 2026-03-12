package entity

import (
	"time"

	"github.com/google/uuid"
)

type Konsultasi struct {
	ID         uuid.UUID `gorm:"type:char(36);primaryKey"`
	UserID     uuid.UUID `gorm:"type:char(36);not null;constraint;OnDelete:CASCADE"`
	DokterID   uuid.UUID `gorm:"type:char(36);not null;constraint;OnDelete:CASCADE"`
	Pesan      string    `gorm:"type:longtext"`
	WaktuPesan time.Time `gorm:"type:timestamp;not null;autoCreateTime"`
}
