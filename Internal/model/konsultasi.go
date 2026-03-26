package model

import (
	"bcc-geazy/internal/entity"
	"time"

	"github.com/google/uuid"
)

type BuatPesan struct {
	UserID   uuid.UUID `json:"user_id"`
	DokterID uuid.UUID `json:"dokter_id"`
	SenderID uuid.UUID `json:"sender_id"`
	Pesan    string    `json:"pesan"`
}

type KonsultasiResponse struct {
	Id         uuid.UUID `json:"id"`
	Pesan      string    `json:"pesan"`
	WaktuPesan time.Time `json:"waktu_pesan"`
}

type EditPesan struct {
	Pesan string `json:"pesan"`
}

func ToKonsultasiResponse(Konsultasi entity.Konsultasi) KonsultasiResponse {
	return KonsultasiResponse{
		Id:         Konsultasi.Id,
		Pesan:      Konsultasi.Pesan,
		WaktuPesan: Konsultasi.WaktuPesan,
	}
}

func ToKonsultasiResponses(Konsultasi []entity.Konsultasi) []KonsultasiResponse {
	var responses []KonsultasiResponse
	for _, Konsultasi := range Konsultasi {
		responses = append(responses, ToKonsultasiResponse(Konsultasi))
	}

	return responses
}

func (e *EditPesan) ToMap() map[string]any {
	updates := map[string]any{}

	if e.Pesan != "" {
		updates["pesan"] = e.Pesan
	}

	return updates
}
