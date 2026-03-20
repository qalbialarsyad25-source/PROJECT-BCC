package entity

import (
	"github.com/google/uuid"
)

type Anak struct {
	Id            uuid.UUID `gorm:"type:char(36);primaryKey"`
	UserID        uuid.UUID `gorm:"type:char(36);not null;constraint;OnDelete:CASCADE"`
	LogMakanan    []Log     `gorm:"foreignKey:AnakID"`
	Nama          string    `gorm:"type:varchar(225); not null"`
	Tinggi        float64   `gorm:"type:decimal(10,2); not null"`
	BeratBadan    float64   `gorm:"type:decimal(10,2); not null"`
	Gender        string    `gorm:"type:varchar(10); not null"`
	LingkarLengan float64   `gorm:"type:decimal(10,2); not null"`
	LingkarKepala float64   `gorm:"type:decimal(10,2); not null"`
	BMI           float64   `gorm:"type:decimal(10,2); not null"`
	StatusGizi    string    `gorm:"type:varchar(10); not null"`
}
