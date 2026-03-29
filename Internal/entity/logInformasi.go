package entity

import (
	"github.com/google/uuid"
)

type LogInformasi struct {
	Id          uuid.UUID `gorm:"type:char(36);primaryKey"`
	Informasi   Informasi `gorm:"foreignKey:InformasiId"`
	InformasiId uuid.UUID `gorm:"type:char(36); not null"`
	User        User      `gorm:"foreignKey:UserID"`
	UserID      uuid.UUID `gorm:"type:char(36); not null; constraint;OnDelete:CASCADE"`
}
