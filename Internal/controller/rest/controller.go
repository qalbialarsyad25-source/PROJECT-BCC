package rest

import (
	websocket "bcc-geazy/internal/controller/delivery"
	"bcc-geazy/internal/usecase"
	"bcc-geazy/pkg/middleware"

	"github.com/go-playground/validator/v10"
)

type V1 struct {
	middleware.IMiddleware
	validator *validator.Validate
	usecase   *usecase.Usecase
	wsManager *websocket.WSManager
}

func NewV1(middleware middleware.IMiddleware, validator *validator.Validate, usecase *usecase.Usecase, wsManager *websocket.WSManager) *V1 {
	return &V1{middleware, validator, usecase, wsManager}
}
