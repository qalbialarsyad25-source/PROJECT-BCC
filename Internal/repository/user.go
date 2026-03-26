package repository

import (
	"bcc-geazy/internal/entity"
	"bcc-geazy/internal/model"
	"context"
	"errors"

	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(ctx context.Context, user entity.User) error
	GetUser(ctx context.Context, pagination model.Pagination) ([]entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
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
		Find(ctx)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User
	err := u.db.WithContext(ctx).
		Where("email = ?", email).
		First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
