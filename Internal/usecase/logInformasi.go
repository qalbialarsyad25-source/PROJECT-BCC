package usecase

import (
	"bcc-geazy/internal/entity"
	"bcc-geazy/internal/model"
	"bcc-geazy/internal/repository"
	"context"

	"github.com/google/uuid"
)

type ILogInformasiUsecase interface {
	GetLogInformasi(ctx context.Context, userID uuid.UUID, pagination model.Pagination) ([]model.LogInformasiResponse, error)
	DeleteLogInformasi(ctx context.Context, id uuid.UUID) error
	CreateLogInformasi(ctx context.Context, userID uuid.UUID, loginfo model.BuatLogInformasi) (*model.LogInformasiResponse, error)
}

type LogInformasiUsecase struct {
	LogInformasiRepository repository.ILogInformasiRepository
}

func NewLogInformasiUsecase(LogInformasiRepo repository.ILogInformasiRepository) *LogInformasiUsecase {
	return &LogInformasiUsecase{LogInformasiRepo}
}

func (p *LogInformasiUsecase) GetLogInformasi(ctx context.Context, userID uuid.UUID, pagination model.Pagination) ([]model.LogInformasiResponse, error) {
	loginformasi, err := p.LogInformasiRepository.GetLogInformasi(ctx, userID, pagination)
	if err != nil {
		return nil, err
	}

	responses := model.ToLogInformasiResponses(loginformasi)
	return responses, nil
}

func (p *LogInformasiUsecase) DeleteLogInformasi(ctx context.Context, id uuid.UUID) error {
	return p.LogInformasiRepository.DeleteLogInformasi(ctx, id)
}

func (p *LogInformasiUsecase) CreateLogInformasi(ctx context.Context, userID uuid.UUID, loginfo model.BuatLogInformasi) (*model.LogInformasiResponse, error) {
	info := entity.LogInformasi{
		Id:          uuid.New(),
		UserID:      userID,
		InformasiId: loginfo.InformasiID,
	}

	err := p.LogInformasiRepository.CreateLogInformasi(ctx, info)
	if err != nil {
		return nil, err
	}

	simpan, err := p.LogInformasiRepository.GetById(ctx, info.Id)
	if err != nil {
		return nil, err
	}

	response := model.ToLogInformasiResponse(*simpan)
	return &response, nil
}
