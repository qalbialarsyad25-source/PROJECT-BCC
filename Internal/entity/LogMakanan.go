package entity

import (
	"github.com/google/uuid"
)

type LogMakanan struct {
	Id        uuid.UUID `gorm:"type:char(36);primaryKey"`
	LogId     uuid.UUID `gorm:"type:char(36);not null; constraint:OnDelete:CASCADE"`
	MakananId uuid.UUID `gorm:"type:char(36);not null; constraint:OnDelete:CASCADE"`
	Gram      float64   `gorm:"type:float"`
	Makanan   Makanan   `gorm:"foreignKey:MakananId"`
}
