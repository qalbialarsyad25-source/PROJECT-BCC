package repository

import (
	"bcc-geazy/internal/entity"
	"bcc-geazy/internal/model"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IAnakRepository interface {
	CreateDataAnak(ctx context.Context, anak entity.Anak) error
	GetDataAnak(ctx context.Context, pagination model.Pagination) ([]entity.Anak, error)
	DeleteDataAnak(ctx context.Context, id uuid.UUID) error
	EditDataAnak(ctx context.Context, id uuid.UUID, edit model.EditDataAnak) error
	GetAnakByID(ctx context.Context, id uuid.UUID) (*entity.Anak, error)
	UpdateAnak(ctx context.Context, anak *entity.Anak) error
	FindUserById(ctx context.Context, userID uuid.UUID) ([]entity.Anak, error)
}

type AnakRepository struct {
	db *gorm.DB
}

func NewAnakRepository(db *gorm.DB) *AnakRepository {
	return &AnakRepository{db}
}

func (p *AnakRepository) CreateDataAnak(ctx context.Context, anak entity.Anak) error {
	err := gorm.G[entity.Anak](p.db).Create(ctx, &anak)
	if err != nil {
		return err
	}

	return nil
}

func (p *AnakRepository) GetDataAnak(ctx context.Context, pagination model.Pagination) ([]entity.Anak, error) {
	anak, err := gorm.G[entity.Anak](p.db).
		Limit(pagination.Limit).
		Offset(pagination.Offset()).
		Order("dibuat_pada DESC").
		Find(ctx)
	if err != nil {
		return nil, err
	}

	return anak, nil
}

func (p *AnakRepository) DeleteDataAnak(ctx context.Context, id uuid.UUID) error {
	rows, err := gorm.G[entity.Anak](p.db).Where("id = ?", id).Delete(ctx)
	if err != nil {
		return err
	}

	if rows == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (p *AnakRepository) EditDataAnak(ctx context.Context, id uuid.UUID, edit model.EditDataAnak) error {
	result := p.db.WithContext(ctx).Model(&entity.Anak{}).
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

func (p *AnakRepository) GetAnakByID(ctx context.Context, id uuid.UUID) (*entity.Anak, error) {
	var anak entity.Anak

	err := p.db.WithContext(ctx).
		Where("id = ?", id).
		First(&anak).Error

	if err != nil {
		return nil, err
	}

	return &anak, err
}

func (p *AnakRepository) UpdateAnak(ctx context.Context, anak *entity.Anak) error {
	return p.db.WithContext(ctx).
		Model(&entity.Anak{}).
		Where("id = ?", anak.Id).
		Updates(map[string]interface{}{
			"profil": anak.Profil,
		}).Error
}

func (p *AnakRepository) FindUserById(ctx context.Context, userID uuid.UUID) ([]entity.Anak, error) {
	var anak []entity.Anak

	err := p.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Find(&anak).Error

	return anak, err
}
