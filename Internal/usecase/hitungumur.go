package usecase

import (
	"time"
)

func HitungUmur(tahun float64) float64 {
	tahunlahir := tahun
	var tahunlahirint int = int(tahunlahir)
	tahunini := time.Now().Year()

	umur := tahunini - tahunlahirint
	var umurfloat float64 = float64(umur)

	return umurfloat
}
