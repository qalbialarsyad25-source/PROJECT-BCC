package entity

import (
	"github.com/google/uuid"
)

type Informasi struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Ringkasan string    `gorm:"type:longtext; not null"`
	Judul     string    `gorm:"type:varchar(225); not null"`
	UserID    uuid.UUID `gorm:"type:uuid;not null;constraint;OnDelete:CASCADE"`
}
