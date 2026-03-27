package utils

import (
	"time"
)

func HitungUmur(tanggalLahir time.Time) int {
	now := time.Now()
	umur := now.Year() - tanggalLahir.Year()

	if now.YearDay() < tanggalLahir.YearDay() {
		umur--
	}

	return umur
}
