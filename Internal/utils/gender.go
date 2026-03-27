package utils

import "strings"

const (
	GenderL = "laki-laki"
	GenderP = "perempuan"
)

func GenderValid(a string) bool {
	a = strings.ToLower(a)
	return a == GenderL || a == GenderP
}
