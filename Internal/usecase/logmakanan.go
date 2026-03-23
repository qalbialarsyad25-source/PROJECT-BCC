package usecase

import (
	"bcc-geazy/internal/model"
	"bcc-geazy/internal/repository"
	"context"
)

type ILogMakananUsecase interface {
	GetLogMakanan(ctx context.Context, pagination model.Pagination) ([]model.LogMakananResponse, error)
}

type LogMakananUsecase struct {
	LogMakananRepository repository.ILogMakananRepository
}

func NewLogMakananUsecase(LogMakananRepository repository.ILogMakananRepository) *LogMakananUsecase {
	return &LogMakananUsecase{LogMakananRepository}
}

func (p *LogMakananUsecase) GetLogMakanan(ctx context.Context, pagination model.Pagination) ([]model.LogMakananResponse, error) {
	logmakanan, err := p.LogMakananRepository.GetLogMakanan(ctx, pagination)
	if err != nil {
		return nil, err
	}

	responses := model.ToLogMakananResponses(logmakanan)
	return responses, nil
}
