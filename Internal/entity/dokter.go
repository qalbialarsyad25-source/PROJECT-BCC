package entity

import (
	"github.com/google/uuid"
)

type Dokter struct {
	Id           uuid.UUID    `gorm:"type:char(36);primaryKey"`
	Nama         string       `gorm:"type:varchar(225); not null"`
	KonsultasiID []Konsultasi `gorm:"foreignKey:DokterID"`
}
