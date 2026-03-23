package usecase

import (
	"bcc-geazy/internal/entity"
	"bcc-geazy/internal/model"
	"bcc-geazy/internal/repository"
	"context"

	"github.com/google/uuid"
)

type ILogUsecase interface {
	GetLog(ctx context.Context, AnakID uuid.UUID, pagination model.Pagination) ([]model.LogResponse, error)
	CreateLog(ctx context.Context, AnakID uuid.UUID, log model.BuatLog) error
	DeleteLog(ctx context.Context, id uuid.UUID) error
}

type LogUsecase struct {
	LogRepository repository.ILogRepository
}

func NewLogUsecase(logRepository repository.ILogRepository) *LogUsecase {
	return &LogUsecase{logRepository}
}

func (p *LogUsecase) GetLog(ctx context.Context, AnakID uuid.UUID, pagination model.Pagination) ([]model.LogResponse, error) {
	log, err := p.LogRepository.GetLog(ctx, AnakID, pagination)
	if err != nil {
		return nil, err
	}

	var responses []model.LogResponse
	for _, log := range log {
		makanan := model.ToLogMakananResponses(log.LogMakanan)

		var totalProtein, totalLemak, totalKarbo float64
		for _, m := range makanan {
			totalProtein += m.Protein
			totalLemak += m.Lemak
			totalKarbo += m.Karbo
		}
		TotalKalori := (totalProtein * 4) + (totalKarbo * 4) + (totalLemak * 9)

		responses = append(responses, model.LogResponse{
			WaktuMakan:  log.WaktuMakan,
			Makanan:     makanan,
			TotalKalori: TotalKalori,
		})
	}

	return responses, nil
}

func (p *LogUsecase) CreateLog(ctx context.Context, AnakID uuid.UUID, req model.BuatLog) error {
	log := entity.Log{
		Id:     uuid.New(),
		AnakID: AnakID,
	}

	var logmakanan []entity.LogMakanan

	for _, m := range req.Makanan {
		logmakanan = append(logmakanan, entity.LogMakanan{
			Id:        uuid.New(),
			MakananId: m.MakananId,
			Gram:      m.Gram,
		})
	}

	return p.LogRepository.CreateLog(ctx, log, logmakanan)

}

func (p *LogUsecase) DeleteLog(ctx context.Context, id uuid.UUID) error {
	return p.LogRepository.DeleteLog(ctx, id)
}
