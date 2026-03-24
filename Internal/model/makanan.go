package model

import (
	"bcc-geazy/internal/entity"

	"github.com/google/uuid"
)

type TambahMakanan struct {
	Nama    string  `json:"nama"`
	Energi  float64 `json:"energi"`
	Protein float64 `json:"protein"`
	Lemak   float64 `json:"lemak"`
	Karbo   float64 `json:"karbo"`
}

type MakananResponse struct {
	Id      uuid.UUID `json:"id"`
	Nama    string    `json:"nama"`
	Energi  float64   `json:"energi"`
	Protein float64   `json:"protein"`
	Lemak   float64   `json:"lemak"`
	Karbo   float64   `json:"karbo"`
}

type EditMakanan struct {
	Nama    string  `json:"nama"`
	Energi  float64 `json:"energi"`
	Protein float64 `json:"protein"`
	Lemak   float64 `json:"lemak"`
	Karbo   float64 `json:"karbo"`
}

func ToMakananResponse(makanan entity.Makanan) MakananResponse {
	return MakananResponse{
		Id:      makanan.Id,
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

func (p *EditMakanan) ToMap() map[string]any {
	perbaruan := map[string]any{}

	if p.Nama != "" {
		perbaruan["nama"] = p.Nama
	}

	if p.Energi != 0 {
		perbaruan["energi"] = p.Energi
	}

	if p.Protein != 0 {
		perbaruan["protein"] = p.Protein
	}

	if p.Lemak != 0 {
		perbaruan["lemak"] = p.Lemak
	}

	if p.Karbo != 0 {
		perbaruan["karbo"] = p.Karbo
	}

	return perbaruan
}
