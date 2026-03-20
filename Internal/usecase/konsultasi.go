package usecase

import (
	"bcc-geazy/internal/entity"
	"bcc-geazy/internal/model"
	"bcc-geazy/internal/repository"
	"context"

	"github.com/google/uuid"
)

type IKonsultasiUsecase interface {
	CreateKonsultasi(ctx context.Context, buatKonsultasi model.BuatPesan) (*model.KonsultasiResponse, error)
	GetKonsultasi(ctx context.Context, pagination model.Pagination) ([]model.KonsultasiResponse, error)
	DeleteKonsultasi(ctx context.Context, id uuid.UUID) error
	EditKonsultasi(ctx context.Context, id uuid.UUID, edit model.EditPesan) error
}

type KonsultasiUsecase struct {
	KonsultasiRepository repository.IKonsulRepository
}

func NewKonsultasiUsecase(konsultasiRepository repository.IKonsulRepository) *KonsultasiUsecase {
	return &KonsultasiUsecase{konsultasiRepository}
}

func (p *KonsultasiUsecase) CreateKonsultasi(ctx context.Context, buatKonsultasi model.BuatPesan) (*model.KonsultasiResponse, error) {
	konsul := entity.Konsultasi{
		Id:         uuid.New(),
		Pesan:      buatKonsultasi.Pesan,
		WaktuPesan: buatKonsultasi.WaktuPesan,
	}

	err := p.KonsultasiRepository.CreateKonsultasi(ctx, konsul)
	if err != nil {
		return nil, err
	}

	response := model.ToKonsultasiResponse(konsul)
	return &response, nil
}

func (p *KonsultasiUsecase) GetKonsultasi(ctx context.Context, pagination model.Pagination) ([]model.KonsultasiResponse, error) {
	konsul, err := p.KonsultasiRepository.GetKonsultasi(ctx, pagination)
	if err != nil {
		return nil, err
	}

	responses := model.ToKonsultasiResponses(konsul)
	return responses, nil
}

func (p *KonsultasiUsecase) DeleteKonsultasi(ctx context.Context, id uuid.UUID) error {
	return p.KonsultasiRepository.DeleteKonsultasi(ctx, id)
}

func (p *KonsultasiUsecase) EditKonsultasi(ctx context.Context, id uuid.UUID, edit model.EditPesan) error {
	return p.KonsultasiRepository.EditKonsultasi(ctx, id, edit)
}
