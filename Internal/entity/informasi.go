package entity

import (
	"github.com/google/uuid"
)

type Informasi struct {
	Id        uuid.UUID `gorm:"type:char(36);primaryKey"`
	Ringkasan string    `gorm:"type:longtext; not null"`
	Judul     string    `gorm:"type:varchar(225); not null"`
	UserID    uuid.UUID `gorm:"type:char(36);not null;constraint;OnDelete:CASCADE"`
}
