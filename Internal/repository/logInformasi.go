package repository

import (
	"bcc-geazy/internal/entity"
	"bcc-geazy/internal/model"
	"context"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

type ILogInformasiRepository interface {
	GetLogInformasi(ctx context.Context, userID uuid.UUID, informasiID uuid.UUID, pagination model.Pagination) ([]entity.LogInformasi, error)
	DeleteLogInformasi(ctx context.Context, id uuid.UUID) error
	CreateLogInformasi(ctx context.Context, loginfo entity.LogInformasi) error
	GetById(ctx context.Context, id uuid.UUID) (*entity.LogInformasi, error)
}

type LogInformasiRepository struct {
	db *gorm.DB
}

func NewLogInformasiRepository(db *gorm.DB) *LogInformasiRepository {
	return &LogInformasiRepository{db}
}

func (p *LogInformasiRepository) GetLogInformasi(ctx context.Context, userID uuid.UUID, informasiID uuid.UUID, pagination model.Pagination) ([]entity.LogInformasi, error) {
	logs, err := gorm.G[entity.LogInformasi](p.db).
		Preload("Informasi", func(db gorm.PreloadBuilder) error {
			return nil
		}).
		Where("user_id = ? AND informasi_id = ?", userID, informasiID).
		Limit(pagination.Limit).
		Offset(pagination.Offset()).
		Find(ctx)

	if err != nil {
		return nil, err
	}

	return logs, nil
}

func (p *LogInformasiRepository) DeleteLogInformasi(ctx context.Context, id uuid.UUID) error {
	hasil, err := gorm.G[entity.LogInformasi](p.db).Where("id = ?", id).Delete(ctx)
	if err != nil {
		return err
	}

	if hasil == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (p *LogInformasiRepository) CreateLogInformasi(ctx context.Context, loginfo entity.LogInformasi) error {
	err := gorm.G[entity.LogInformasi](p.db).Create(ctx, &loginfo)
	if err != nil {
		return err
	}

	return nil
}

func (p *LogInformasiRepository) GetById(ctx context.Context, id uuid.UUID) (*entity.LogInformasi, error) {
	log, err := gorm.G[entity.LogInformasi](p.db).
		Preload("Informasi", func(db gorm.PreloadBuilder) error {
			return nil
		}).
		Where("id = ?", id).
		First(ctx)

	if err != nil {
		return nil, err
	}

	return &log, nil
}
