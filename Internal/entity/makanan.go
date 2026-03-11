package entity

import (
	"github.com/google/uuid"
)

type Makanan struct {
	ID         uuid.UUID    `gorm:"type:chat(36);primaryKey"`
	Nama       string       `gorm:"type:varchar(100); not null"`
	LogMakanan []LogMakanan `gorm:"foreignKey:MakananID"`
	Energi     float64      `gorm:"type:float"`
	Protein    float64      `gorm:"type:float"`
	Lemak      float64      `gorm:"type:float"`
	Karbo      float64      `gorm:"type:float"`
}
