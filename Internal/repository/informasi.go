package repository

import (
	"bcc-geazy/internal/entity"
	"bcc-geazy/internal/model"
	"context"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

type IInformasiRepository interface {
	GetInformasi(ctx context.Context, pagination model.Pagination) ([]entity.Informasi, error)
	CreateInformasi(ctx context.Context, informasi entity.Informasi) error
	DeleteInformasi(ctx context.Context, id uuid.UUID) error
	EditInformasi(ctx context.Context, id uuid.UUID, edit model.EditInformasi) error
}

type InformasiRepository struct {
	db *gorm.DB
}

func NewInformasiRepository(db *gorm.DB) *InformasiRepository {
	return &InformasiRepository{db}
}

func (p *InformasiRepository) CreateInformasi(ctx context.Context, informasi entity.Informasi) error {
	err := gorm.G[entity.Informasi](p.db).Create(ctx, &informasi)
	if err != nil {
		return err
	}

	return nil
}

func (p *InformasiRepository) GetInformasi(ctx context.Context, pagination model.Pagination) ([]entity.Informasi, error) {
	informasi, err := gorm.G[entity.Informasi](p.db).
		Limit(pagination.Limit).
		Offset(pagination.Offset()).
		Order("dibuat_pada DESC").
		Find(ctx)
	if err != nil {
		return nil, err
	}

	return informasi, nil
}

func (p *InformasiRepository) DeleteInformasi(ctx context.Context, id uuid.UUID) error {
	rows, err := gorm.G[entity.Informasi](p.db).Where("id = ?", id).Delete(ctx)
	if err != nil {
		return err
	}

	if rows == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (p *InformasiRepository) EditInformasi(ctx context.Context, id uuid.UUID, edit model.EditInformasi) error {
	result := p.db.WithContext(ctx).Model(&entity.Informasi{}).
		Where("id = ?", id).
		Updates(edit.ToMap())

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
