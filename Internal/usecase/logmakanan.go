package usecase

import (
	"bcc-geazy/internal/model"
	"bcc-geazy/internal/repository"
	"context"

	"github.com/google/uuid"
)

type ILogMakananUsecase interface {
	GetLogMakanan(ctx context.Context, pagination model.Pagination) ([]model.LogMakananResponse, error)
	DeleteLogMakanan(ctx context.Context, id uuid.UUID) error
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

func (p *LogMakananUsecase) DeleteLogMakanan(ctx context.Context, id uuid.UUID) error {
	return p.LogMakananRepository.DeleteLogMakanan(ctx, id)
}
