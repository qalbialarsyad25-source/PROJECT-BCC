package entity

import (
	"github.com/google/uuid"
)

type Makanan struct {
	ID         uuid.UUID    `gorm:"type:char(36);primaryKey"`
	Nama       string       `gorm:"type:varchar(100); not null"`
	LogMakanan []LogMakanan `gorm:"foreignKey:MakananID"`
	AnakID     uuid.UUID    `gorm:"type:char(36);not null; constraint;OnDelete;CASCADE"`
	Energi     float64      `gorm:"type:float"`
	Protein    float64      `gorm:"type:float"`
	Lemak      float64      `gorm:"type:float"`
	Karbo      float64      `gorm:"type:float"`
}
