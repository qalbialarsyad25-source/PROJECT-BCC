package usecase

import (
	"bcc-geazy/internal/entity"
	"bcc-geazy/internal/model"
	"bcc-geazy/internal/repository"
	"context"

	"github.com/google/uuid"
)

type IDokterUsecase interface {
	CreateDataDokter(ctx context.Context, buatDokter model.BuatUserDokter) (*model.DokterResponse, error)
	GetDokter(ctx context.Context, pagination model.Pagination) ([]model.DokterResponse, error)
	DeleteDokter(ctx context.Context, id uuid.UUID) error
	EditDokter(ctx context.Context, id uuid.UUID, edit model.EditDokter) error
}

type DokterUsecase struct {
	DokterRepository repository.IDokterRepository
}

func NewDokterUsecase(dokterRepository repository.IDokterRepository) *DokterUsecase {
	return &DokterUsecase{dokterRepository}
}

func (p *DokterUsecase) CreateDataDokter(ctx context.Context, buatDokter model.BuatUserDokter) (*model.DokterResponse, error) {
	dokter := entity.Dokter{
		Id:   uuid.New(),
		Nama: buatDokter.Nama,
	}

	err := p.DokterRepository.CreateDataDokter(ctx, dokter)
	if err != nil {
		return nil, err
	}

	response := model.ToDokterResponse(dokter)
	return &response, nil
}

func (p *DokterUsecase) GetDokter(ctx context.Context, pagination model.Pagination) ([]model.DokterResponse, error) {
	dokter, err := p.DokterRepository.GetUserDokter(ctx, pagination)
	if err != nil {
		return nil, err
	}

	responses := model.ToDokterResponses(dokter)
	return responses, nil
}

func (p *DokterUsecase) DeleteDokter(ctx context.Context, id uuid.UUID) error {
	return p.DokterRepository.DeleteDataDokter(ctx, id)
}

func (p *DokterUsecase) EditDokter(ctx context.Context, id uuid.UUID, edit model.EditDokter) error {
	return p.DokterRepository.EditDataDokter(ctx, id, edit)
}
