package repository

import (
	"bcc-geazy/internal/entity"
	"bcc-geazy/internal/model"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(ctx context.Context, user entity.User) error
	GetUser(ctx context.Context, pagination model.Pagination) ([]entity.User, error)
	DeleteUser(ctx context.Context, id uuid.UUID) error
	EditUser(ctx context.Context, id uuid.UUID, edit model.EditUser) error
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (p *UserRepository) CreateUser(ctx context.Context, user entity.User) error {
	err := gorm.G[entity.User](p.db).Create(ctx, &user)
	if err != nil {
		return err
	}

	return nil
}

func (p *UserRepository) GetUser(ctx context.Context, pagination model.Pagination) ([]entity.User, error) {
	user, err := gorm.G[entity.User](p.db).
		Limit(pagination.Limit).
		Offset(pagination.Offset()).
		Order("Dibuat_pada ").
		Find(ctx)
	if err != nil {
		return nil, err
	}

	return user, nil
}
