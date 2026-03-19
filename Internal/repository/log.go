package repository

import (
	"bcc-geazy/internal/entity"
	"bcc-geazy/internal/model"
	"context"

	"gorm.io/gorm"
)

type ILogRepository interface {
	GetLog(ctx context.Context, pagination model.Pagination) ([]entity.Log, error)
}

type LogRepository struct {
	db *gorm.DB
}

func NewLogRepository(db *gorm.DB) *LogRepository {
	return &LogRepository{db}
}

func (p *LogRepository) GetLog(ctx context.Context, pagination model.Pagination) ([]entity.Log, error) {
	log, err := gorm.G[entity.Log](p.db).
		Limit(pagination.Limit).
		Offset(pagination.Offset()).
		Order("Dibuat pada ").
		Find(ctx)
	if err != nil {
		return nil, err
	}

	return log, nil
}
