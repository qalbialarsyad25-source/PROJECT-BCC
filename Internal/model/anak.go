package model

import (
	"bcc-geazy/internal/entity"
	"time"

	"github.com/google/uuid"
)

type TambahDataAnak struct {
	Nama            string  `json:"nama"`
	TanggalLahir    string  `json:"tanggal_lahir"`
	Tinggi          float64 `json:"tinggi"`
	BeratBadan      float64 `json:"berat_badan"`
	Gender          string  `json:"gender"`
	AnakKe          int     `json:"anak_ke"`
	LingkarKepala   float64 `json:"lingkar_kepala"`
	LingkarLengan   float64 `json:"lingkar_lengan"`
	GolonganDarah   string  `json:"golongan_darah"`
	Alergi          string  `json:"alergi"`
	RiwayatPenyakit string  `json:"riwayat_penyakit"`
}

type EditDataAnak struct {
	Nama            string   `json:"nama"`
	TanggalLahir    string   `json:"tanggal_lahir"`
	Tinggi          *float64 `json:"tinggi"`
	BeratBadan      *float64 `json:"berat_badan"`
	Gender          string   `json:"gender"`
	AnakKe          int      `json:"anak_ke"`
	LingkarKepala   *float64 `json:"lingkar_kepala"`
	LingkarLengan   *float64 `json:"lingkar_lengan"`
	GolonganDarah   string   `json:"golongan_darah"`
	Alergi          string   `json:"alergi"`
	RiwayatPenyakit string   `json:"riwayat_penyakit"`
}

type AnakResponse struct {
	Id              uuid.UUID `json:"id"`
	Nama            string    `json:"nama"`
	TanggalLahir    time.Time `json:"tanggal_lahir"`
	Umur            int       `json:"umur"`
	Tinggi          float64   `json:"tinggi"`
	BeratBadan      float64   `json:"berat_badan"`
	Gender          string    `json:"gender"`
	AnakKe          int       `json:"anak_ke"`
	AnakKeLabel     string    `json:"anak_ke_label"`
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

	if p.TanggalLahir != "" {
		perbaruan["tanggal_lahir"] = p.TanggalLahir
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

	if p.AnakKe != 0 {
		perbaruan["anak_ke"] = p.AnakKe
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
		Umur:            Anak.Umur,
		Tinggi:          Anak.Tinggi,
		BeratBadan:      Anak.BeratBadan,
		Gender:          Anak.Gender,
		AnakKe:          Anak.AnakKe,
		AnakKeLabel:     AnakKeToLabel(int(Anak.AnakKe)),
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

func AnakKeToLabel(n int) string {
	labels := []string{
		"Pertama", "Kedua", "Ketiga",
	}

	if n >= 1 && n <= len(labels) {
		return "Anak " + labels[n-1]
	}
	return ""
}
