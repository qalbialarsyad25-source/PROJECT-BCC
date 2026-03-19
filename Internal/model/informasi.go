package model

import (
	"bcc-geazy/internal/entity"
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
	Ringkasan string `json:"ringkasan"`
	Judul     string `json:"judul"`
}

func toInformasiResponse(informasi entity.Informasi) InformasiResponse {
	return InformasiResponse{
		Ringkasan: informasi.Ringkasan,
		Judul:     informasi.Judul,
	}
}

func toInformasiResponses(informasi []entity.Informasi) []InformasiResponse {
	var responses []InformasiResponse
	for _, Informasi := range informasi {
		responses = append(responses, toInformasiResponse(Informasi))
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
