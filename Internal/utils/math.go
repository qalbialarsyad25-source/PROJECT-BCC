package utils

import "math"

func Pembulatan(val float64) float64 {
	return math.Round(val*100) / 100
}
