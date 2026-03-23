package usecase

import (
	"bcc-geazy/internal/model"
)

func hitungKalori(protein, lemak, karbo float64) float64 {
	return (protein * 4) + (karbo * 4) + (lemak * 4)
}

func HitungTotalNutrisi(makanan []model.LogMakananResponse) (protein, lemak, karbo float64) {
	for _, m := range makanan {
		protein += m.Protein
		lemak += m.Lemak
		karbo += m.Karbo
	}
	return
}
