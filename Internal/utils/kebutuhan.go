package utils

import "strings"

func HitungKebutuhanKalori(umur int, berat float64, g string) float64 {
	g = strings.ToLower(g)

	if g == "laki laki" || g == "laki-laki" {
		g = "l"
	} else if g == "perempuan" {
		g = "p"
	}

	switch {
	case umur <= 3:
		if g == "l" {
			return (59.5 * berat) - 30
		}
		return (58.3 * berat) - 31

	case umur <= 10:
		if g == "l" {
			return (22.7 * berat) + 495
		}
		return (20.3 * berat) + 485

	case umur <= 18:
		if g == "l" {
			return (17.5 * berat) + 651
		}
		return (12.2 * berat) + 764
	}

	return 1500
}

func HitungKebutuhanProtein(berat float64) float64 {
	return 1.2 * berat
}

func HitungKebutuhanLemak(kalori float64) float64 {
	return (0.3 * kalori) / 9
}
