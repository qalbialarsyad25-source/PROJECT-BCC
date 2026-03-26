package repository

import (
	"bcc-geazy/internal/entity"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IKonsulRepository interface {
	CreateKonsultasi(ctx context.Context, konsul entity.Konsultasi) error
	GetByUserdanDokter(ctx context.Context, userID, dokterID uuid.UUID) ([]entity.Konsultasi, error)
	PesanDibaca(ctx context.Context, userID, dokterID uuid.UUID) error
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

func (p *KonsultasiRepository) GetByUserdanDokter(ctx context.Context, userID, dokterID uuid.UUID) ([]entity.Konsultasi, error) {
	var hasil []entity.Konsultasi

	err := p.db.WithContext(ctx).
		Where("user_id = ? AND dokter_id = ?", userID, dokterID).
		Order("waktu_pesan ASC").
		Find(&hasil).Error

	return hasil, err
}

func (p *KonsultasiRepository) PesanDibaca(ctx context.Context, userID, dokterID uuid.UUID) error {
	return p.db.WithContext(ctx).
		Model(&entity.Konsultasi{}).
		Where("user_id = ? AND dokter_id = ? AND dibaca = false", userID, dokterID).
		Update("dibaca", true).Error
}
