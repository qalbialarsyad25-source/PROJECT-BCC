package usecase

import (
	"bcc-geazy/internal/entity"
	"bcc-geazy/internal/model"
	"bcc-geazy/internal/repository"
	"context"

	"github.com/google/uuid"
)

type IMakananUsecase interface {
	GetMakanan(ctx context.Context, pagination model.Pagination) ([]model.MakananResponse, error)
	CreateMakanan(ctx context.Context, buatmakanan model.TambahMakanan) (*model.MakananResponse, error)
	DeleteMakanan(ctx context.Context, id uuid.UUID) error
	EditMakanan(ctx context.Context, id uuid.UUID, edit model.EditMakanan) error
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

func (p *MakananUsecase) CreateMakanan(ctx context.Context, buatmakanan model.TambahMakanan) (*model.MakananResponse, error) {
	makanan := entity.Makanan{
		Id:      uuid.New(),
		Nama:    buatmakanan.Nama,
		Energi:  buatmakanan.Energi,
		Protein: buatmakanan.Protein,
		Lemak:   buatmakanan.Lemak,
		Karbo:   buatmakanan.Karbo,
	}

	err := p.MakananRepository.CreateMakanan(ctx, makanan)
	if err != nil {
		return nil, err
	}

	response := model.ToMakananResponse(makanan)
	return &response, nil
}

func (p *MakananUsecase) DeleteMakanan(ctx context.Context, id uuid.UUID) error {
	return p.MakananRepository.DeleteMakanan(ctx, id)
}

func (p *MakananUsecase) EditMakanan(ctx context.Context, id uuid.UUID, edit model.EditMakanan) error {
	return p.MakananRepository.EditMakanan(ctx, id, edit)
}
