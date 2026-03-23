package model

import (
	"bcc-geazy/internal/entity"
	"time"

	"github.com/google/uuid"
)

type BuatInformasi struct {
	Ringkasan string `json:"ringkasan"`
	Judul     string `json:"judul"`
}

type EditInformasi struct {
	Ringkasan string `json:"ringkasan"`
	Judul     string `json:"judul"`
}

type InformasiResponse struct {
	Id         uuid.UUID `json:"id"`
	Judul      string    `json:"judul"`
	Ringkasan  string    `json:"ringkasan"`
	DibuatPada time.Time `json:"dibuat_pada"`
}

func ToInformasiResponse(informasi entity.Informasi) InformasiResponse {
	return InformasiResponse{
		Id:         informasi.Id,
		Ringkasan:  informasi.Ringkasan,
		Judul:      informasi.Judul,
		DibuatPada: informasi.DibuatPada,
	}
}

func ToInformasiResponses(informasi []entity.Informasi) []InformasiResponse {
	var responses []InformasiResponse
	for _, Informasi := range informasi {
		responses = append(responses, ToInformasiResponse(Informasi))
	}

	return responses
}

func (e *EditInformasi) ToMap() map[string]any {
	updates := map[string]any{}

	if e.Ringkasan != "" {
		updates["ringkasan"] = e.Ringkasan
	}
	if e.Judul != "" {
		updates["judul"] = e.Judul
	}

	return updates
}
