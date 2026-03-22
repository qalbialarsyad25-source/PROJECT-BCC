package entity

import (
	"time"

	"github.com/google/uuid"
)

type Log struct {
	Id         uuid.UUID    `gorm:"type:char(36);primaryKey"`
	LogMakanan []LogMakanan `gorm:"foreignKey:LogId"`
	AnakID     uuid.UUID    `gorm:"type:char(36);not null;constraint;OnDelete:CASCADE"`
	WaktuMakan time.Time    `gorm:"type:timestamp; not null;autoCreateTime"`
}
