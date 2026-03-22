package usecase

import "math"

func HitungBMI(berat, tinggi float64) (float64, string) {
	tinggiMeter := tinggi / 100
	bmi := berat / (tinggiMeter * tinggiMeter)

	bmiformat := math.Round(bmi*100) / 100

	var status string
	if bmi < 18.5 {
		status = "Kurus"
	} else if bmi < 25 {
		status = "Normal"
	} else if bmi < 30 {
		status = "Gemuk"
	} else if bmi >= 30 {
		status = "Obesitas"
	}

	return bmiformat, status
}
