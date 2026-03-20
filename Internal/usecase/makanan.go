package usecase

import (
	"bcc-geazy/internal/model"
	"bcc-geazy/internal/repository"
	"context"
)

type IMakananUsecase interface {
	GetMakanan(ctx context.Context, pagination model.Pagination) ([]model.MakananResponse, error)
}

type MakananUsecase struct {
	MakananRepository repository.IMakananRepository
}

func NewMakananUsecase(makananRepository repository.IMakananRepository) *MakananUsecase {
	return &MakananUsecase{makananRepository}
}

func (p *MakananUsecase) GetMakanan(ctx context.Context, pagination model.Pagination) ([]model.MakananResponse, error) {
	makanan, err := p.MakananRepository.GetMakanan(ctx, pagination)
	if err != nil {
		return nil, err
	}

	responses := model.ToMakananResponses(makanan)
	return responses, nil
}
