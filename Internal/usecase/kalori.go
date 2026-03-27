package usecase

import (
	"bcc-geazy/internal/entity"
)

func hitungKalori(protein, lemak, karbo float64) float64 {
	return (protein * 4) + (karbo * 4) + (lemak * 9)
}

func HitungTotalNutrisi(makanan []entity.Log) (protein, lemak, karbo float64) {
	for _, a := range makanan {
		for _, m := range a.LogMakanan {
			protein += m.Makanan.Protein * (m.Gram / 100)
			lemak += m.Makanan.Lemak * (m.Gram / 100)
			karbo += m.Makanan.Karbo * (m.Gram / 100)
		}
	}
	return
}

func HitungNutrisiPerLog(log entity.Log) (protein, lemak, karbo float64) {
	for _, m := range log.LogMakanan {
		protein += m.Makanan.Protein * (m.Gram / 100)
		lemak += m.Makanan.Lemak * (m.Gram / 100)
		karbo += m.Makanan.Karbo * (m.Gram / 100)
	}
	return
}

func HitungPersen(total, target float64) float64 {
	if target <= 0 {
		return 0.0
	}
	return (total / target) * 100
}
