package usecase

import (
	"bcc-geazy/internal/entity"
	"bcc-geazy/internal/model"
	"bcc-geazy/internal/repository"
	"context"
	"time"

	"github.com/google/uuid"
)

type IInformasiUsecase interface {
	CreateInformasi(ctx context.Context, userID uuid.UUID, buatInformasi model.BuatInformasi) (*model.InformasiResponse, error)
	GetInformasi(ctx context.Context, pagination model.Pagination) ([]model.InformasiResponse, error)
	DeleteInformasi(ctx context.Context, id uuid.UUID) error
	EditInformasi(ctx context.Context, id uuid.UUID, edit model.EditInformasi) error
}

type InformasiUsecase struct {
	InformasiRepository repository.IInformasiRepository
}

func NewInformasiUsecase(informasiRepository repository.IInformasiRepository) *InformasiUsecase {
	return &InformasiUsecase{informasiRepository}
}

func (p *InformasiUsecase) CreateInformasi(ctx context.Context, userID uuid.UUID, buatInformasi model.BuatInformasi) (*model.InformasiResponse, error) {
	informasi := entity.Informasi{
		Id:         uuid.New(),
		UserID:     nil,
		Ringkasan:  buatInformasi.Ringkasan,
		Judul:      buatInformasi.Judul,
		DibuatPada: time.Now(),
	}

	err := p.InformasiRepository.CreateInformasi(ctx, informasi)
	if err != nil {
		return nil, err
	}

	response := model.ToInformasiResponse(informasi)
	return &response, nil
}

func (p *InformasiUsecase) GetInformasi(ctx context.Context, pagination model.Pagination) ([]model.InformasiResponse, error) {
	informasi, err := p.InformasiRepository.GetInformasi(ctx, pagination)
	if err != nil {
		return nil, err
	}

	responses := model.ToInformasiResponses(informasi)
	return responses, nil
}

func (p *InformasiUsecase) DeleteInformasi(ctx context.Context, id uuid.UUID) error {
	return p.InformasiRepository.DeleteInformasi(ctx, id)
}

func (p *InformasiUsecase) EditInformasi(ctx context.Context, id uuid.UUID, edit model.EditInformasi) error {
	return p.InformasiRepository.EditInformasi(ctx, id, edit)
}
