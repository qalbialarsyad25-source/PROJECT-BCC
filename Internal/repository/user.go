package repository

import (
	"bcc-geazy/internal/entity"
	"bcc-geazy/internal/model"
	"context"
	"time"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(ctx context.Context, user entity.User) error
	GetUser(ctx context.Context, pagination model.Pagination) ([]entity.User, error)
	GetUserById(ctx context.Context, id uuid.UUID) (*entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	SaveResetToken(ctx context.Context, userID uuid.UUID, token string, expired time.Time) error
	GetUserByResetToken(ctx context.Context, token string) (*entity.User, error)
	UpdatePassword(ctx context.Context, userID uuid.UUID, password string) error
	ClearResetToken(ctx context.Context, userID uuid.UUID) error
	UpdateUser(ctx context.Context, user *entity.User) error
	GetAll(ctx context.Context) ([]entity.User, error)
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

func (p *UserRepository) GetUserByEmail(ctx context.Context, userEmail string) (*entity.User, error) {
	var user entity.User

	err := p.db.WithContext(ctx).
		Where("email = ?", userEmail).
		First(&user).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (p *UserRepository) SaveResetToken(ctx context.Context, userID uuid.UUID, token string, expired time.Time) error {
	return p.db.WithContext(ctx).
		Model(&entity.User{}).
		Where("id = ?", userID).
		Updates(map[string]interface{}{
			"reset_token":            token,
			"reset_token_expired_at": expired,
		}).Error
}

func (p *UserRepository) GetUserByResetToken(ctx context.Context, token string) (*entity.User, error) {
	var user entity.User

	err := p.db.WithContext(ctx).
		Where("reset_token = ? AND reset_token_expired_at > ?", token, time.Now()).
		First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (p *UserRepository) UpdatePassword(ctx context.Context, userID uuid.UUID, password string) error {
	return p.db.WithContext(ctx).
		Model(&entity.User{}).
		Where("id = ?", userID).
		Update("password", password).Error
}

func (p *UserRepository) ClearResetToken(ctx context.Context, userID uuid.UUID) error {
	return p.db.WithContext(ctx).
		Model(&entity.User{}).
		Where("id = ?", userID).
		Updates(map[string]interface{}{
			"reset_token":            "",
			"reset_token_expired_at": nil,
		}).Error
}

func (p *UserRepository) GetUserById(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	var user entity.User

	err := p.db.WithContext(ctx).
		Where("id = ?", id).
		First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (p *UserRepository) UpdateUser(ctx context.Context, user *entity.User) error {
	return p.db.WithContext(ctx).Save(&user).Error
}

func (p *UserRepository) GetAll(ctx context.Context) ([]entity.User, error) {
	var users []entity.User

	err := p.db.WithContext(ctx).Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}
