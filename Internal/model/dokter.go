package model

import (
	"bcc-geazy/internal/entity"

	"github.com/google/uuid"
)

type BuatUserDokter struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	Nama      string `json:"nama"`
	Spesialis string `json:"spesialis"`
	Profil    string `json:"profil"`
}

type DokterResponse struct {
	Id        uuid.UUID `json:"id"`
	Nama      string    `json:"nama"`
	Spesialis string    `json:"spesialis"`
	Profil    string    `json:"profil"`
}

type EditDokter struct {
	Nama      string `json:"nama"`
	Spesialis string `json:"spesialis"`
	Profil    string `json:"profil"`
}

func ToDokterResponse(Dokter entity.Dokter) DokterResponse {
	return DokterResponse{
		Id:        Dokter.Id,
		Nama:      Dokter.Nama,
		Spesialis: Dokter.Spesialis,
		Profil:    Dokter.Profil,
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

	if e.Spesialis != "" {
		updates["spesialis"] = e.Spesialis
	}

	if e.Profil != "" {
		updates["Profil"] = e.Profil
	}
	return updates
}
