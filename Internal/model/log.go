package model

import (
	"bcc-geazy/internal/entity"
	"time"
)

type LogResponse struct {
	WaktuMakan time.Time `json:"waktu_makan"`
}

func ToLogResponse(log entity.Log) LogResponse {
	return LogResponse{
		WaktuMakan: log.WaktuMakan,
	}
}

func ToLogResponses(log []entity.Log) []LogResponse {
	var responses []LogResponse
	for _, Log := range log {
		responses = append(responses, ToLogResponse(Log))
	}

	return responses
}
