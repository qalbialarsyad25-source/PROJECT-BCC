package repository

import (
	"bcc-geazy/internal/entity"
	"bcc-geazy/internal/model"
	"context"
	"time"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

type ILogRepository interface {
	CreateLog(ctx context.Context, log entity.Log, logmakanan []entity.LogMakanan) error
	GetLog(ctx context.Context, AnakID uuid.UUID, pagination model.Pagination) ([]entity.Log, error)
	DeleteLog(ctx context.Context, id uuid.UUID) error
	GetLogHariIni(ctx context.Context, anakID uuid.UUID) ([]entity.Log, error)
}

type LogRepository struct {
	db *gorm.DB
}

func NewLogRepository(db *gorm.DB) *LogRepository {
	return &LogRepository{db}
}

func (p *LogRepository) CreateLog(ctx context.Context, log entity.Log, logmakanan []entity.LogMakanan) error {
	return p.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := gorm.G[entity.Log](tx).Create(ctx, &log); err != nil {
			return err
		}

		for i := range logmakanan {
			logmakanan[i].LogId = log.Id
		}

		for _, lm := range logmakanan {
			if err := gorm.G[entity.LogMakanan](tx).Create(ctx, &lm); err != nil {
				return err
			}
		}

		return nil
	})
}

func (p *LogRepository) GetLog(ctx context.Context, AnakID uuid.UUID, pagination model.Pagination) ([]entity.Log, error) {
	logs, err := gorm.G[entity.Log](p.db).
		Preload("LogMakanan.Makanan", func(db gorm.PreloadBuilder) error {
			return nil
		}).
		Where("anak_id = ?", AnakID).
		Limit(pagination.Limit).
		Offset(pagination.Offset()).
		Order("waktu_makan DESC").
		Find(ctx)
	if err != nil {
		return nil, err
	}

	return logs, nil
}

func (p *LogRepository) DeleteLog(ctx context.Context, id uuid.UUID) error {
	return p.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("log_id = ?", id).Delete(&entity.LogMakanan{}).Error; err != nil {
			return err
		}

		rows, err := gorm.G[entity.Log](tx).Where("id = ?", id).Delete(ctx)
		if err != nil {
			return err
		}

		if rows == 0 {
			return gorm.ErrRecordNotFound
		}

		return nil
	})
}

func (p *LogRepository) GetLogHariIni(ctx context.Context, anakID uuid.UUID) ([]entity.Log, error) {
	var logs []entity.Log

	mulaiHariini := time.Now().Truncate(24 * time.Hour)
	akhirHariini := mulaiHariini.Add(24 * time.Hour)

	err := p.db.WithContext(ctx).
		Preload("LogMakanan.Makanan").
		Where("anak_id = ? AND waktu_makan BETWEEN ? AND ?", anakID, mulaiHariini, akhirHariini).
		Order("waktu_makan DESC").
		Find(&logs).Error

	if err != nil {
		return nil, err
	}

	return logs, nil
}
