package entity

import (
	"github.com/google/uuid"
)

type Dokter struct {
	Id           uuid.UUID    `gorm:"type:char(36);primaryKey"`
	UserId       uuid.UUID    `gorm:"type:char(36);uniqueIndex;not null"`
	Nama         string       `gorm:"type:varchar(225); not null"`
	Spesialis    string       `gorm:"type:varchar(225); not null"`
	Profil       string       `gorm:"type:text"`
	KonsultasiID []Konsultasi `gorm:"foreignKey:DokterID"`
}
