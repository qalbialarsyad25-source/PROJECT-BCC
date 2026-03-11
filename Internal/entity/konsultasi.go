package entity

import (
	"time"

	"github.com/google/uuid"
)

type Konsultasi struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID     uuid.UUID `gorm:"type:uuid;not null;constraint;OnDelete:CASCADE"`
	DokterID   uuid.UUID `gorm:"type:uuid;not null;constraint;OnDelete:CASCADE"`
	Pesan      string    `gorm:"type:longtext"`
	WaktuPesan time.Time `gorm:"type:timestamp;not null;autoCreateTime"`
}
