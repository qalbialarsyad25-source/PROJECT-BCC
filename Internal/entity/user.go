package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id             uuid.UUID    `gorm:"type:uuid;primaryKey"`
	GoogleID       string       `gorm:"uniqueIndex;type:varchar(225); not null"`
	Nama           string       `gorm:"type:varchar(225); not null"`
	Email          string       `gorm:"uniqueIndex;type:varchar(225); not null"`
	Anak           []Anak       `gorm:"foreignKey:UserID"`
	DibuatPada     time.Time    `gorm:"type:timestamp;not null;autoCreatedTime"`
	DiperbaruiPada time.Time    `gorm:"type:timestamp;not null;autoCreatedTime"`
	Konsultasi     []Konsultasi `gorm:"foreignKey:UserID"`
	LogMakanan     []LogMakanan `gorm:"foreignKey:UserID"`
	Informasi      []Informasi  `gorm:"foreignKey:UserID"`
}
