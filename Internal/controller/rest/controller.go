package rest

import "bcc-geazy/internal/usecase"

type V1 struct {
	usecase *usecase.Usecase
}

func NewV1(usecase *usecase.Usecase) *V1 {
	return &V1{usecase}
}
