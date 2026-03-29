package usecase

import (
	"bcc-geazy/internal/entity"
	"bcc-geazy/internal/model"
	"bcc-geazy/internal/repository"
	"context"
	"time"

	"github.com/google/uuid"
)

type INotifikasiUsecase interface {
	CreateNotifikasi(ctx context.Context, userID uuid.UUID, notip model.BuatNotifikasi) (*model.NotifikasiResponse, error)
	GetNotifikasi(ctx context.Context, pagination model.Pagination) ([]model.NotifikasiResponse, error)
	DeleteNotifikasi(ctx context.Context, id uuid.UUID) error
	EditNotifikasi(ctx context.Context, id uuid.UUID, edit model.EditNotifikasi) error
}

type NotifikasiUsecase struct {
	NotifikasiRepository repository.INotifikasiRepository
}

func NewNotifikasiUsecase(notip repository.INotifikasiRepository) *NotifikasiUsecase {
	return &NotifikasiUsecase{notip}
}

func (p *NotifikasiUsecase) CreateNotifikasi(ctx context.Context, userID uuid.UUID, notip model.BuatNotifikasi) (*model.NotifikasiResponse, error) {
	notipikasi := entity.Notifikasi{
		Id:         uuid.New(),
		UserId:     userID,
		Pesan:      notip.Pesan,
		Judul:      notip.Judul,
		Dibaca:     false,
		DibuatPada: time.Now(),
	}

	err := p.NotifikasiRepository.CreateNotifikasi(ctx, notipikasi)
	if err != nil {
		return nil, err
	}

	response := model.ToNotifikasiResponse(notipikasi)
	return &response, nil
}

func (p *NotifikasiUsecase) GetNotifikasi(ctx context.Context, pagination model.Pagination) ([]model.NotifikasiResponse, error) {
	notipikasi, err := p.NotifikasiRepository.GetNotifikasi(ctx, pagination)
	if err != nil {
		return nil, err
	}

	responses := model.ToNotifikasiResponses(notipikasi)
	return responses, nil
}

func (p *NotifikasiUsecase) DeleteNotifikasi(ctx context.Context, id uuid.UUID) error {
	return p.NotifikasiRepository.DeleteNotifikasi(ctx, id)
}

func (p *NotifikasiUsecase) EditNotifikasi(ctx context.Context, id uuid.UUID, edit model.EditNotifikasi) error {
	return p.NotifikasiRepository.EditNotifikasi(ctx, id, edit)
}
