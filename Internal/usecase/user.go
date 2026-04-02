package usecase

import (
	"bcc-geazy/internal/model"
	"bcc-geazy/internal/repository"
	"bcc-geazy/pkg/storage"
	"context"
	"errors"
	"mime/multipart"

	"github.com/google/uuid"
)

type IUserUsecase interface {
	GetProfile(ctx context.Context, id uuid.UUID) (*model.UserResponse, error)
	UploadFotoUser(ctx context.Context, userID uuid.UUID, file multipart.File, filename string) (string, error)
}

type UserUsecase struct {
	UserRepository repository.IUserRepository
}

func NewUserUsecase(userRepository repository.IUserRepository) *UserUsecase {
	return &UserUsecase{userRepository}
}

func (p *UserUsecase) GetProfile(ctx context.Context, id uuid.UUID) (*model.UserResponse, error) {
	user, err := p.UserRepository.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user tidak ditemukan")
	}

	return &model.UserResponse{
		Id:       user.Id,
		Nama:     user.Nama,
		Email:    user.Email,
		UserName: user.UserName,
		Profil:   user.Profil,
	}, nil
}

func (p *UserUsecase) UploadFotoUser(ctx context.Context, userID uuid.UUID, file multipart.File, filename string) (string, error) {

	user, err := p.UserRepository.GetUserById(ctx, userID)
	if err != nil {
		return "", err
	}

	url, err := storage.UploadFile(file, filename)
	if err != nil {
		return "", err
	}

	user.Profil = url

	err = p.UserRepository.UpdateUser(ctx, user)
	if err != nil {
		return "", err
	}

	return url, nil
}
