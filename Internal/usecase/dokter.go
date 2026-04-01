package usecase

import (
	"bcc-geazy/internal/entity"
	"bcc-geazy/internal/model"
	"bcc-geazy/internal/repository"
	"bcc-geazy/pkg/bcrypt"
	"context"

	"github.com/google/uuid"
)

type IDokterUsecase interface {
	CreateDataDokter(ctx context.Context, buatDokter model.BuatUserDokter) (*model.DokterResponse, error)
	GetDokter(ctx context.Context, pagination model.Pagination) ([]model.DokterResponse, error)
	DeleteDokter(ctx context.Context, id uuid.UUID) error
	EditDokter(ctx context.Context, id uuid.UUID, edit model.EditDokter) error
	GetDokterByID(ctx context.Context, id uuid.UUID) (entity.Dokter, error)
}

type DokterUsecase struct {
	DokterRepository repository.IDokterRepository
	UserRepository   repository.IUserRepository
	Bcrypt           bcrypt.IBcrypt
}

func NewDokterUsecase(dokterRepository repository.IDokterRepository, userRepository repository.IUserRepository, bcrypt bcrypt.IBcrypt) *DokterUsecase {
	return &DokterUsecase{
		DokterRepository: dokterRepository,
		UserRepository:   userRepository,
		Bcrypt:           bcrypt,
	}
}

func (p *DokterUsecase) CreateDataDokter(ctx context.Context, buatDokter model.BuatUserDokter) (*model.DokterResponse, error) {
	hashed, err := p.Bcrypt.GenerateHash(buatDokter.Password)
	if err != nil {
		return nil, err
	}

	user := entity.User{
		Id:       uuid.New(),
		Email:    buatDokter.Email,
		Password: hashed,
		Role:     "dokter",
	}

	err = p.UserRepository.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	dokter := entity.Dokter{
		Id:        uuid.New(),
		UserId:    user.Id,
		Nama:      buatDokter.Nama,
		Spesialis: buatDokter.Spesialis,
	}

	err = p.DokterRepository.CreateDataDokter(ctx, dokter)
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

func (p *DokterUsecase) GetDokterByID(ctx context.Context, id uuid.UUID) (entity.Dokter, error) {
	return p.DokterRepository.GetDokterByID(ctx, id)
}
