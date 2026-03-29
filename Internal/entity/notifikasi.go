package entity

import (
	"time"

	"github.com/google/uuid"
)

type Notifikasi struct {
	Id         uuid.UUID `gorm:"type:char(36);primaryKey"`
	UserId     uuid.UUID `gorm:"type:char(36); not null"`
	Judul      string    `gorm:"type:varchar(225); not null"`
	Pesan      string    `gorm:"type:varchar(225); not null"`
	Dibaca     bool      `gorm:"default:false"`
	DibuatPada time.Time `gorm:"type:timestamp; not null; autoCreateTime"`
}
