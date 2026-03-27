package usecase

import (
	"bcc-geazy/internal/model"
	"bcc-geazy/internal/repository"
	"bcc-geazy/internal/utils"
	"context"

	"github.com/google/uuid"
)

type INutrisiUsecase interface {
	GetNutrisiHarian(ctx context.Context, anakID uuid.UUID) (model.NutrisiHarianResponse, error)
}

type NutrisiUsecase struct {
	LogRepository  repository.ILogRepository
	AnakRepository repository.IAnakRepository
}

func NewNutrisiUsecase(logRepo repository.ILogRepository, anakRepo repository.IAnakRepository) *NutrisiUsecase {
	return &NutrisiUsecase{
		LogRepository:  logRepo,
		AnakRepository: anakRepo,
	}
}

func (u *NutrisiUsecase) GetNutrisiHarian(ctx context.Context, anakID uuid.UUID) (model.NutrisiHarianResponse, error) {

	logs, err := u.LogRepository.GetLogHariIni(ctx, anakID)
	if err != nil {
		return model.NutrisiHarianResponse{}, err
	}

	anak, err := u.AnakRepository.GetAnakByID(ctx, anakID)
	if err != nil {
		return model.NutrisiHarianResponse{}, err
	}

	protein, lemak, karbo := HitungTotalNutrisi(logs)
	totalKalori := hitungKalori(protein, lemak, karbo)

	umur := utils.HitungUmur(anak.TanggalLahir)

	targetKalori := utils.HitungKebutuhanKalori(umur, anak.BeratBadan, anak.Gender)
	targetProtein := utils.HitungKebutuhanProtein(anak.BeratBadan)
	targetLemak := utils.HitungKebutuhanLemak(targetKalori)

	persenKalori := HitungPersen(totalKalori, targetKalori)
	persenProtein := HitungPersen(protein, targetProtein)
	persenLemak := HitungPersen(lemak, targetLemak)

	return model.NutrisiHarianResponse{
		TotalKalori:   utils.Pembulatan(totalKalori),
		TargetKalori:  utils.Pembulatan(targetKalori),
		TargetProtein: utils.Pembulatan(targetProtein),
		TargetLemak:   utils.Pembulatan(targetLemak),
		PersenProtein: utils.Pembulatan(persenProtein),
		PersenLemak:   utils.Pembulatan(persenLemak),
		PersenKalori:  utils.Pembulatan(persenKalori),
	}, nil
}
