package repository

import (
	"bcc-geazy/internal/entity"
	"bcc-geazy/internal/model"
	"context"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

type INotifikasiRepository interface {
	GetNotifikasi(ctx context.Context, pagination model.Pagination) ([]entity.Notifikasi, error)
	CreateNotifikasi(ctx context.Context, notifikasi entity.Notifikasi) error
	DeleteNotifikasi(ctx context.Context, id uuid.UUID) error
	EditNotifikasi(ctx context.Context, id uuid.UUID, edit model.EditNotifikasi) error
}

type NotifikasiRepository struct {
	db *gorm.DB
}

func NewNotifikasiRepository(db *gorm.DB) *NotifikasiRepository {
	return &NotifikasiRepository{db}
}

func (p *NotifikasiRepository) CreateNotifikasi(ctx context.Context, notip entity.Notifikasi) error {
	err := gorm.G[entity.Notifikasi](p.db).Create(ctx, &notip)
	if err != nil {
		return err
	}

	return nil
}

func (p *NotifikasiRepository) GetNotifikasi(ctx context.Context, pagination model.Pagination) ([]entity.Notifikasi, error) {
	notip, err := gorm.G[entity.Notifikasi](p.db).
		Limit(pagination.Limit).
		Offset(pagination.Offset()).
		Order("dibuat_pada DESC").
		Find(ctx)
	if err != nil {
		return nil, err
	}

	return notip, nil
}

func (p *NotifikasiRepository) DeleteNotifikasi(ctx context.Context, id uuid.UUID) error {
	rows, err := gorm.G[entity.Notifikasi](p.db).Where("id = ?", id).Delete(ctx)
	if err != nil {
		return err
	}

	if rows == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (p *NotifikasiRepository) EditNotifikasi(ctx context.Context, id uuid.UUID, edit model.EditNotifikasi) error {
	result := p.db.WithContext(ctx).Model(&entity.Notifikasi{}).
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
