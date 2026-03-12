package entity

import (
	"time"

	"github.com/google/uuid"
)

type LogMakanan struct {
	ID          uuid.UUID `gorm:"type:char(36);primaryKey"`
	MakananID   uuid.UUID `gorm:"type:char(36);not null;constraint;OnDelete:CASCADE"`
	AnakID      uuid.UUID `gorm:"type:char(36);not null;constraint;OnDelete:CASCADE"`
	Gram        float64   `gorm:"type:float"`
	TotalKalori float64   `gorm:"type:float"`
	WaktuMakan  time.Time `gorm:"type:timestamp; not null;autoCreateTime"`
}
