package model

import (
	"bcc-geazy/internal/entity"
	"time"

	"github.com/google/uuid"
)

type LogResponse struct {
	Id          uuid.UUID            `json:"id"`
	WaktuMakan  time.Time            `json:"waktu_makan"`
	Makanan     []LogMakananResponse `json:"makanan"`
	TotalKalori float64              `json:"total_kalori"`
}

type DetailMakanan struct {
	MakananId uuid.UUID `json:"makanan_id"`
	Gram      float64   `json:"gram"`
}

type BuatLog struct {
	Makanan []DetailMakanan `json:"makanan"`
}

func ToLogResponse(log entity.Log) LogResponse {
	return LogResponse{
		Id:         log.Id,
		WaktuMakan: log.WaktuMakan,
		Makanan:    ToLogMakananResponses(log.LogMakanan),
	}
}

func ToLogResponses(log []entity.Log) []LogResponse {
	var responses []LogResponse
	for _, Log := range log {
		responses = append(responses, ToLogResponse(Log))
	}

	return responses
}
