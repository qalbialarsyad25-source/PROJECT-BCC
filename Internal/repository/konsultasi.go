package repository

import (
	"bcc-geazy/internal/entity"
	"bcc-geazy/internal/model"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IKonsulRepository interface {
	CreateKonsultasi(ctx context.Context, konsul entity.Konsultasi) error
	GetKonsultasi(ctx context.Context, pagination model.Pagination) ([]entity.Konsultasi, error)
	DeleteKonsultasi(ctx context.Context, id uuid.UUID) error
	EditKonsultasi(ctx context.Context, id uuid.UUID, edit model.EditPesan) error
}

type KonsultasiRepository struct {
	db *gorm.DB
}

func NewKonsulRepository(db *gorm.DB) *KonsultasiRepository {
	return &KonsultasiRepository{db}
}

func (p *KonsultasiRepository) CreateKonsultasi(ctx context.Context, konsul entity.Konsultasi) error {
	err := gorm.G[entity.Konsultasi](p.db).Create(ctx, &konsul)
	if err != nil {
		return err
	}

	return nil
}

func (p *KonsultasiRepository) GetKonsultasi(ctx context.Context, pagination model.Pagination) ([]entity.Konsultasi, error) {
	konsul, err := gorm.G[entity.Konsultasi](p.db).
		Limit(pagination.Limit).
		Offset(pagination.Offset()).
		Order("Dibuat_pada ").
		Find(ctx)
	if err != nil {
		return nil, err
	}

	return konsul, nil
}

func (p *KonsultasiRepository) DeleteKonsultasi(ctx context.Context, id uuid.UUID) error {
	rows, err := gorm.G[entity.Konsultasi](p.db).Where("id + ?", id).Delete(ctx)
	if err != nil {
		return err
	}

	if rows == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (p *KonsultasiRepository) EditKonsultasi(ctx context.Context, id uuid.UUID, edit model.EditPesan) error {
	result := p.db.WithContext(ctx).Model(&entity.Konsultasi{}).
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
