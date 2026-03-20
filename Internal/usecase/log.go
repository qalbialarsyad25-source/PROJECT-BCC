package usecase

import (
	"bcc-geazy/internal/model"
	"bcc-geazy/internal/repository"
	"context"
)

type ILogUsecase interface {
	GetLog(ctx context.Context, pagination model.Pagination) ([]model.LogResponse, error)
}

type LogUsecase struct {
	LogRepository repository.ILogRepository
}

func NewLogUsecase(logRepository repository.ILogRepository) *LogUsecase {
	return &LogUsecase{logRepository}
}

func (p *LogUsecase) GetLog(ctx context.Context, pagination model.Pagination) ([]model.LogResponse, error) {
	log, err := p.LogRepository.GetLog(ctx, pagination)
	if err != nil {
		return nil, err
	}

	responses := model.ToLogResponses(log)
	return responses, nil
}
