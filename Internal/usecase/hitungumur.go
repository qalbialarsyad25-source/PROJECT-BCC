package usecase

import (
	"fmt"
	"time"
)

func HitungUmur(tahun int) string {
	tahunini := time.Now().Year()

	umur := tahunini - tahun

	return fmt.Sprintf("%d Tahun", umur)
}
