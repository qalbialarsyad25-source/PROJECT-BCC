package entity

import (
	"time"
)

type LogMakanan struct {
	ID          uint `gorm:"primaryKey"`
	UserID      uint
	Makanan     []Makanan `gorm:"foreignKey:LogMakananID"`
	AnakID      uint
	Gram        float64 `gorm:"type:float"`
	TotalKalori float64 `gorm:"type:float"`
	WaktuMakan  time.Time
}
