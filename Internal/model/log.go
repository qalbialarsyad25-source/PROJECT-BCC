package model

import (
	"bcc-geazy/internal/entity"
	"time"
)

type LogReponse struct {
	WaktuMakan time.Time `json:"waktu_makan"`
}

func toLogResponse(log entity.Log) LogReponse {
	return LogReponse{
		WaktuMakan: log.WaktuMakan,
	}
}

func toLogResponses(log []entity.Log) []LogReponse {
	var responses []LogReponse
	for _, Log := range log {
		responses = append(responses, toLogResponse(Log))
	}

	return responses
}
