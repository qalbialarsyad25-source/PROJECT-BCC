package utils

import "strings"

const (
	GolonganO  = "O"
	GolonganA  = "A"
	GolonganB  = "B"
	GolonganAB = "AB"
)

func ValidGolonganDarah(a string) bool {
	a = strings.ToUpper(a)
	return a == GolonganA || a == GolonganAB || a == GolonganB || a == GolonganO
}
