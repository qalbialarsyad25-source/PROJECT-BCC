package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id                  uuid.UUID      `gorm:"type:char(36);primaryKey"`
	GoogleID            *string        `gorm:"uniqueIndex;type:varchar(225)"`
	Nama                string         `gorm:"type:varchar(225); not null"`
	UserName            string         `gorm:"type:varchar(225); not null"`
	Email               string         `gorm:"uniqueIndex;type:varchar(225); not null"`
	Anak                []Anak         `gorm:"foreignKey:UserID"`
	Password            string         `gorm:"type:varchar(255); not null"`
	Role                string         `gorm:"type:varchar(20);not null"`
	Konsultasi          []Konsultasi   `gorm:"foreignKey:UserID"`
	Informasi           []Informasi    `gorm:"foreignKey:UserID"`
	LogInformasi        []LogInformasi `gorm:"foreignKey:UserID"`
	ResetToken          string
	ResetTokenExpiredAt *time.Time
	Profil              string `gorm:"type:text"`
}
