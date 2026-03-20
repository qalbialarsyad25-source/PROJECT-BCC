package model

import (
	"bcc-geazy/internal/entity"

	"github.com/google/uuid"
)

type TambahDataAnak struct {
	Nama          string  `json:"nama"`
	Tinggi        float64 `json:"tinggi"`
	BeratBadan    float64 `json:"berat_badan"`
	Gender        string  `json:"gender"`
	LingkarKepala float64 `json:"lingkar_kepala"`
	LingkarLengan float64 `json:"lingkar_lengan"`
}

type EditDataAnak struct {
	Nama          string   `json:"nama"`
	Tinggi        *float64 `json:"tinggi"`
	BeratBadan    *float64 `json:"berat_badan"`
	Gender        string   `json:"gender"`
	LingkarKepala *float64 `json:"lingkar_kepala"`
	LingkarLengan *float64 `json:"lingkar_lengan"`
}

type AnakResponse struct {
	Id            uuid.UUID `json:"id"`
	Nama          string    `json:"nama"`
	Tinggi        float64   `json:"tinggi"`
	BeratBadan    float64   `json:"berat_badan"`
	Gender        string    `json:"gender"`
	LingkarKepala float64   `json:"lingkar_kepala"`
	LingkarLengan float64   `json:"lingkar_lengan"`
	BMI           float64   `json:"bmi"`
	StatusGizi    string    `json:"status"`
}

func (p *EditDataAnak) ToMap() map[string]any {
	perbaruan := map[string]any{}

	if p.Nama != "" {
		perbaruan["nama"] = p.Nama
	}

	if p.Tinggi != nil {
		perbaruan["tinggi"] = p.Tinggi
	}

	if p.BeratBadan != nil {
		perbaruan["berat_badan"] = p.BeratBadan
	}

	if p.Gender != "" {
		perbaruan["gender"] = p.Gender
	}

	if p.LingkarKepala != nil {
		perbaruan["lingkar_kepala"] = p.LingkarKepala
	}

	if p.LingkarLengan != nil {
		perbaruan["lingkar_lengan"] = p.LingkarLengan
	}

	return perbaruan
}

func ToAnakResponse(Anak entity.Anak) AnakResponse {
	return AnakResponse{
		Id:            Anak.Id,
		Nama:          Anak.Nama,
		Tinggi:        Anak.Tinggi,
		BeratBadan:    Anak.BeratBadan,
		Gender:        Anak.Gender,
		LingkarKepala: Anak.LingkarKepala,
		LingkarLengan: Anak.LingkarLengan,
		BMI:           Anak.BMI,
		StatusGizi:    Anak.StatusGizi,
	}
}

func ToAnakResponses(Anak []entity.Anak) []AnakResponse {
	var responses []AnakResponse
	for _, Anak := range Anak {
		responses = append(responses, ToAnakResponse(Anak))
	}

	return responses
}
