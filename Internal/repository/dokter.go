package repository

import (
	"bcc-geazy/internal/entity"
	"bcc-geazy/internal/model"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IDokterRepository interface {
	CreateDataDokter(ctx context.Context, dokter entity.Dokter) error
	GetUserDokter(ctx context.Context, pagination model.Pagination) ([]entity.Dokter, error)
	DeleteDataDokter(ctx context.Context, id uuid.UUID) error
	EditDataDokter(ctx context.Context, id uuid.UUID, edit model.EditDokter) error
	GetDokterByID(ctx context.Context, id uuid.UUID) (entity.Dokter, error)
	UpdateDokter(ctx context.Context, dokter entity.Dokter) error
}

type DokterRepository struct {
	db *gorm.DB
}

func NewDokterRepository(db *gorm.DB) *DokterRepository {
	return &DokterRepository{db}
}

func (p *DokterRepository) CreateDataDokter(ctx context.Context, dokter entity.Dokter) error {
	err := gorm.G[entity.Dokter](p.db).Create(ctx, &dokter)
	if err != nil {
		return err
	}

	return nil
}

func (p *DokterRepository) GetUserDokter(ctx context.Context, pagination model.Pagination) ([]entity.Dokter, error) {
	dokter, err := gorm.G[entity.Dokter](p.db).
		Limit(pagination.Limit).
		Offset(pagination.Offset()).
		Find(ctx)
	if err != nil {
		return nil, err
	}

	return dokter, nil
}

func (p *DokterRepository) DeleteDataDokter(ctx context.Context, id uuid.UUID) error {
	rows, err := gorm.G[entity.Dokter](p.db).Where("id = ?", id).Delete(ctx)
	if err != nil {
		return err
	}

	if rows == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (p *DokterRepository) EditDataDokter(ctx context.Context, id uuid.UUID, edit model.EditDokter) error {
	result := p.db.WithContext(ctx).Model(&entity.Dokter{}).
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

func (p *DokterRepository) GetDokterByID(ctx context.Context, id uuid.UUID) (entity.Dokter, error) {
	var dokter entity.Dokter
	err := p.db.WithContext(ctx).First(&dokter, "id = ?", id).Error
	return dokter, err
}

func (p *DokterRepository) UpdateDokter(ctx context.Context, dokter entity.Dokter) error {
	return p.db.WithContext(ctx).Save(&dokter).Error
}
