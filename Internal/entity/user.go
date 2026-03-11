package entity

import (
	"time"
)

type User struct {
	ID         uint   `gorm:"primaryKey"`
	GoogleID   string `gorm:"uniqueIndex;type:varchar(225); not null"`
	Nama       string `gorm:"type:varchar(225); not null"`
	Email      string `gorm:"uniqueIndex;type:varchar(225); not null"`
	Anak       []Anak `gorm:"foreignKey:UserID"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Userkonsul []Konsultasi `gorm:"foreignKey:UserID"`
	LogMakanan []LogMakanan `gorm:"foreignKey:UserID"`
	Informasi  []Informasi  `gorm:"foreignKey:UserID"`
}
