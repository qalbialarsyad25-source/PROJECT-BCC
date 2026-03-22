package repository

import (
	"bcc-geazy/internal/entity"
	"bcc-geazy/internal/model"
	"context"

	"gorm.io/gorm"
)

type IMakananRepository interface {
	GetMakanan(ctx context.Context, pagination model.Pagination) ([]entity.Makanan, error)
}

type MakananRepository struct {
	db *gorm.DB
}

func NewMakananRepository(db *gorm.DB) *MakananRepository {
	return &MakananRepository{db}
}

func (p *MakananRepository) GetMakanan(ctx context.Context, pagination model.Pagination) ([]entity.Makanan, error) {
	makanan, err := gorm.G[entity.Makanan](p.db).
		Limit(pagination.Limit).
		Offset(pagination.Offset()).
		Find(ctx)
	if err != nil {
		return nil, err
	}

	return makanan, nil
}
