package model

import (
	"bcc-geazy/internal/entity"
)

type MakananResponse struct {
	Nama    string  `json:"nama"`
	Energi  float64 `json:"energi"`
	Protein float64 `json:"Protein"`
	Lemak   float64 `json:"lemak"`
	Karbo   float64 `json:"karbo"`
}

func ToMakananResponse(makanan entity.Makanan) MakananResponse {
	return MakananResponse{
		Nama:    makanan.Nama,
		Energi:  makanan.Energi,
		Protein: makanan.Protein,
		Lemak:   makanan.Lemak,
		Karbo:   makanan.Karbo,
	}
}

func ToMakananResponses(makanan []entity.Makanan) []MakananResponse {
	var responses []MakananResponse
	for _, Makanan := range makanan {
		responses = append(responses, ToMakananResponse(Makanan))
	}

	return responses
}
