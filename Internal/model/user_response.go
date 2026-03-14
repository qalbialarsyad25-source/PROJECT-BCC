package model

import "github.com/google/uuid"

type UserResponse struct {
	Id             uuid.UUID `json:"id" binding:"required"`
	Nama           string    `json:"nama" binding:"required"`
	Email          string    `json:"email" binding:"required"`
	DibuatPada     string    `json:"dibuat_pada" binding:"required"`
	DiperbaruiPada string    `json:"diperbarui_pada" binding:"required`
}
