package model

import (
	"bcc-geazy/internal/entity"

	"github.com/google/uuid"
)

type BuatUserDokter struct {
	Nama string `json:"nama"`
}

type DokterResponse struct {
	Id   uuid.UUID `json:"id"`
	Nama string    `json:"nama"`
}

type EditDokter struct {
	Nama string `json:"nama"`
}

func ToDokterResponse(Dokter entity.Dokter) DokterResponse {
	return DokterResponse{
		Id:   Dokter.Id,
		Nama: Dokter.Nama,
	}
}

func ToDokterResponses(Dokter []entity.Dokter) []DokterResponse {
	var responses []DokterResponse
	for _, Dokter := range Dokter {
		responses = append(responses, ToDokterResponse(Dokter))
	}

	return responses
}

func (e *EditDokter) ToMap() map[string]any {
	updates := map[string]any{}

	if e.Nama != "" {
		updates["nama"] = e.Nama
	}
	return updates
}
