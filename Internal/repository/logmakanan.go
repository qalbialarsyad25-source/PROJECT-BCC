package repository

import (
	"bcc-geazy/internal/entity"
	"bcc-geazy/internal/model"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ILogMakananRepository interface {
	GetLogMakanan(ctx context.Context, pagination model.Pagination) ([]entity.LogMakanan, error)
	DeleteLogMakanan(ctx context.Context, id uuid.UUID) error
}

type LogMakananRepository struct {
	db *gorm.DB
}

func NewLogMakananRepository(db *gorm.DB) *LogMakananRepository {
	return &LogMakananRepository{db}
}

func (p *LogMakananRepository) GetLogMakanan(ctx context.Context, pagination model.Pagination) ([]entity.LogMakanan, error) {
	logmakanan, err := gorm.G[entity.LogMakanan](p.db).
		Preload("makanan", func(db gorm.PreloadBuilder) error {
			return nil
		}).
		Limit(pagination.Limit).
		Offset(pagination.Offset()).
		Order("Dibuat_pada ").
		Find(ctx)
	if err != nil {
		return nil, err
	}

	return logmakanan, nil
}

func (p *LogMakananRepository) DeleteLogMakanan(ctx context.Context, id uuid.UUID) error {
	rows, err := gorm.G[entity.LogMakanan](p.db).Where("id = ?", id).Delete(ctx)
	if err != nil {
		return err
	}

	if rows == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
