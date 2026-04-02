package model

import (
	"bcc-geazy/internal/entity"

	"github.com/google/uuid"
)

type BuatLogInformasi struct {
	InformasiId uuid.UUID `json:"informasi_id"`
}

type LogInformasiResponse struct {
	Id        uuid.UUID         `json:"id"`
	Informasi InformasiResponse `json:"informasi"`
}

func ToLogInformasiResponse(a entity.LogInformasi) LogInformasiResponse {
	return LogInformasiResponse{
		Id:        a.Id,
		Informasi: ToInformasiResponse(a.Informasi),
	}
}

func ToLogInformasiResponses(a []entity.LogInformasi) []LogInformasiResponse {
	var responses []LogInformasiResponse
	for _, LogInformasi := range a {
		responses = append(responses, ToLogInformasiResponse(LogInformasi))
	}
	return responses

}
