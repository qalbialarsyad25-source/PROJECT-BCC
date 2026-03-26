package model

import (
	"bcc-geazy/internal/entity"

	"github.com/google/uuid"
)

type TambahDataAnak struct {
	Nama            string  `json:"nama"`
	TanggalLahir    float64 `json:"tanggal_lahir"`
	BulanLahir      float64 `json:"bulan_lahir"`
	TahunLahir      float64 `json"tahun_lahir"`
	Umur            float64 `json:"umur"`
	Tinggi          float64 `json:"tinggi"`
	BeratBadan      float64 `json:"berat_badan"`
	Gender          string  `json:"gender"`
	LingkarKepala   float64 `json:"lingkar_kepala"`
	LingkarLengan   float64 `json:"lingkar_lengan"`
	GolonganDarah   string  `json:"golongan_darah"`
	Alergi          string  `json:"alergi"`
	RiwayatPenyakit string  `json:"riwayat_penyakit"`
}

type EditDataAnak struct {
	Nama            string   `json:"nama"`
	TanggalLahir    *float64 `json:"tanggal_lahir"`
	BulanLahir      *float64 `json:"bulan_lahir"`
	TahunLahir      *float64 `json"tahun_lahir"`
	Umur            *float64 `json:"umur"`
	Tinggi          *float64 `json:"tinggi"`
	BeratBadan      *float64 `json:"berat_badan"`
	Gender          string   `json:"gender"`
	LingkarKepala   *float64 `json:"lingkar_kepala"`
	LingkarLengan   *float64 `json:"lingkar_lengan"`
	GolonganDarah   string   `json:"golongan_darah"`
	Alergi          string   `json:"alergi"`
	RiwayatPenyakit string   `json:"riwayat_penyakit"`
}

type AnakResponse struct {
	Id              uuid.UUID `json:"id"`
	Nama            string    `json:"nama"`
	TanggalLahir    float64   `json:"tanggal_lahir"`
	BulanLahir      float64   `json:"bulan_lahir"`
	TahunLahir      float64   `json"tahun_lahir"`
	Umur            float64   `json:"umur"`
	Tinggi          float64   `json:"tinggi"`
	BeratBadan      float64   `json:"berat_badan"`
	Gender          string    `json:"gender"`
	LingkarKepala   float64   `json:"lingkar_kepala"`
	LingkarLengan   float64   `json:"lingkar_lengan"`
	GolonganDarah   string    `json:"golongan_darah"`
	Alergi          string    `json:"alergi"`
	RiwayatPenyakit string    `json:"riwayat_penyakit"`
	BMI             float64   `json:"bmi"`
	StatusGizi      string    `json:"status"`
}

func (p *EditDataAnak) ToMap() map[string]any {
	perbaruan := map[string]any{}

	if p.Nama != "" {
		perbaruan["nama"] = p.Nama
	}

	if p.TanggalLahir != nil {
		perbaruan["tanggal_lahir"] = p.TanggalLahir
	}

	if p.BulanLahir != nil {
		perbaruan["bulan_lahir"] = p.BulanLahir
	}

	if p.TahunLahir != nil {
		perbaruan["tahun_lahir"] = p.TahunLahir
	}

	if p.Umur != nil {
		perbaruan["umur"] = p.Umur
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

	if p.GolonganDarah != "" {
		perbaruan["golongan_darah"] = p.GolonganDarah
	}

	if p.Alergi != "" {
		perbaruan["alergi"] = p.Alergi
	}

	if p.RiwayatPenyakit != "" {
		perbaruan["riwayat_penyakit"] = p.RiwayatPenyakit
	}

	return perbaruan
}

func ToAnakResponse(Anak entity.Anak) AnakResponse {
	return AnakResponse{
		Id:              Anak.Id,
		Nama:            Anak.Nama,
		TanggalLahir:    Anak.TanggalLahir,
		BulanLahir:      Anak.BulanLahir,
		TahunLahir:      Anak.TahunLahir,
		Umur:            Anak.Umur,
		Tinggi:          Anak.Tinggi,
		BeratBadan:      Anak.BeratBadan,
		Gender:          Anak.Gender,
		LingkarKepala:   Anak.LingkarKepala,
		LingkarLengan:   Anak.LingkarLengan,
		GolonganDarah:   Anak.GolonganDarah,
		Alergi:          Anak.Alergi,
		RiwayatPenyakit: Anak.RiwayatPenyakit,
		BMI:             Anak.BMI,
		StatusGizi:      Anak.StatusGizi,
	}
}

func ToAnakResponses(Anak []entity.Anak) []AnakResponse {
	var responses []AnakResponse
	for _, Anak := range Anak {
		responses = append(responses, ToAnakResponse(Anak))
	}

	return responses
}
