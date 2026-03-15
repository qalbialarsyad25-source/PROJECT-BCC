package model

import (
	"bcc-geazy/internal/entity"
)

type BuatInformasi struct {
	Ringkasan string `json:"ringkasan"`
	Judul     string `json:"judul"`
}

type InformasiResponse struct {
	Ringkasan string `json:"ringkasan"`
	Judul     string `json:"judul"`
}

func toInformasiResponse(informasi entity.Informasi)
