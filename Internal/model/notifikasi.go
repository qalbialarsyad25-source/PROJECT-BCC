package model

import (
	"bcc-geazy/internal/entity"
	"time"

	"github.com/google/uuid"
)

type BuatNotifikasi struct {
	Pesan string `json:"pesan"`
	Judul string `json:"judul"`
}

type EditNotifikasi struct {
	Pesan string `json:"pesan"`
	Judul string `json:"judul"`
}

type NotifikasiResponse struct {
	Id         uuid.UUID `json:"id"`
	Judul      string    `json:"judul"`
	Pesan      string    `json:"pesan"`
	DibuatPada time.Time `json:"dibuat_pada"`
}

func ToNotifikasiResponse(notip entity.Notifikasi) NotifikasiResponse {
	return NotifikasiResponse{
		Id:         notip.Id,
		Judul:      notip.Judul,
		Pesan:      notip.Pesan,
		DibuatPada: notip.DibuatPada,
	}
}

func ToNotifikasiResponses(notifikasi []entity.Notifikasi) []NotifikasiResponse {
	var responses []NotifikasiResponse
	for _, Notip := range notifikasi {
		responses = append(responses, ToNotifikasiResponse(Notip))
	}

	return responses
}

func (e *EditNotifikasi) ToMap() map[string]any {
	updates := map[string]any{}

	if e.Pesan != "" {
		updates["pesan"] = e.Pesan
	}
	if e.Judul != "" {
		updates["judul"] = e.Judul
	}

	return updates
}
