package rest

import (
	"bcc-geazy/internal/usecase"
	"bcc-geazy/pkg/middleware"

	"github.com/go-playground/validator/v10"
)

type V1 struct {
	middleware.IMiddleware
	validator *validator.Validate
	usecase   *usecase.Usecase
}

func NewV1(middleware middleware.IMiddleware, validator *validator.Validate, usecase *usecase.Usecase) *V1 {
	return &V1{middleware, validator, usecase}
}
