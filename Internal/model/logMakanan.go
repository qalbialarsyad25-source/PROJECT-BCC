package model

import (
	"bcc-geazy/internal/entity"
)

type LogMakananResponse struct {
	Gram float64 `json:"gram"`
}

func toLogMakananResponse(logmakanan entity.LogMakanan) LogMakananResponse {
	return LogMakananResponse{
		Gram: logmakanan.Gram,
	}
}

func toLogMakananResponses(logmakanan []entity.LogMakanan) []LogMakananResponse {
	var responses []LogMakananResponse
	for _, Logmakanan := range logmakanan {
		responses = append(responses, toLogMakananResponse(Logmakanan))
	}

	return responses
}
