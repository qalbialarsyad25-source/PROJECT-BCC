package model

import (
	"bcc-geazy/internal/entity"

	"github.com/google/uuid"
)

type LogMakananResponse struct {
	Id      uuid.UUID `json:"id"`
	Nama    string    `json:"nama"`
	Gram    float64   `json:"gram"`
	Energi  float64   `json:"energi"`
	Protein float64   `json:"protein"`
	Lemak   float64   `json:"lemak"`
	Karbo   float64   `json:"karbo"`
}

func ToLogMakananResponse(logmakanan entity.LogMakanan) LogMakananResponse {
	m := logmakanan.Makanan

	return LogMakananResponse{
		Id:   m.Id,
		Nama: m.Nama,
		Gram: logmakanan.Gram,

		Energi:  (m.Energi / 100) * logmakanan.Gram,
		Protein: (m.Protein / 100) * logmakanan.Gram,
		Lemak:   (m.Lemak / 100) * logmakanan.Gram,
		Karbo:   (m.Karbo / 100) * logmakanan.Gram,
	}
}

func ToLogMakananResponses(logmakanan []entity.LogMakanan) []LogMakananResponse {
	var responses []LogMakananResponse
	for _, Logmakanan := range logmakanan {
		responses = append(responses, ToLogMakananResponse(Logmakanan))
	}

	return responses
}
