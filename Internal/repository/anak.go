package repository

import (
	"bcc-geazy/internal/entity"
	"bcc-geazy/internal/model"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IAnakRepository interface {
	CreateAnak(ctx context.Context, anak entity.Anak) error
	GetUserAnak(ctx context.Context, pagination model.Pagination) ([]entity.Anak, error)
	DeleteAnak(ctx context.Context, id uuid.UUID) error
	EditAnak(ctx context.Context, id uuid.UUID, edit model.EditUser) error
}

type AnakRepository struct {
	db *gorm.DB
}

func NewAnakRepository(db *gorm.DB) *AnakRepository {
	return &AnakRepository{db}
}

func (p *AnakRepository) CreateAnak(ctx context.Context, anak entity.Anak) error {
	err := gorm.G[entity.Anak](p.db).Create(ctx, &anak)
	if err != nil {
		return err
	}

	return nil
}

func (p *AnakRepository) GetUserAnak(ctx context.Context, pagination model.Pagination) ([]entity.Anak, error) {
	anak, err := gorm.G[entity.Anak](p.db).
		Limit(pagination.Limit).
		Offset(pagination.Offset()).
		Order("Dibuat_pada ").
		Find(ctx)
	if err != nil {
		return nil, err
	}

	return anak, nil
}
