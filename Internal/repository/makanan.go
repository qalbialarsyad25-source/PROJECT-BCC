package repository

import (
	"bcc-geazy/internal/entity"
	"bcc-geazy/internal/model"
	"context"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

type IMakananRepository interface {
	GetMakanan(ctx context.Context, pagination model.Pagination) ([]entity.Makanan, error)
	CreateMakanan(ctx context.Context, makanan entity.Makanan) error
	EditMakanan(ctx context.Context, id uuid.UUID, edit model.EditMakanan) error
	DeleteMakanan(ctx context.Context, id uuid.UUID) error
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

func (p *MakananRepository) CreateMakanan(ctx context.Context, makanan entity.Makanan) error {
	err := p.db.WithContext(ctx).Create(&makanan).Error
	if err != nil {
		return err
	}

	return nil
}

func (p *MakananRepository) EditMakanan(ctx context.Context, id uuid.UUID, edit model.EditMakanan) error {
	result := p.db.WithContext(ctx).Model(&entity.Makanan{}).
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

func (p *MakananRepository) DeleteMakanan(ctx context.Context, id uuid.UUID) error {
	rows, err := gorm.G[entity.Makanan](p.db).Where("id = ?", id).Delete(ctx)
	if err != nil {
		return err
	}

	if rows == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
