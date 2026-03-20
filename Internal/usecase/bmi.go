package usecase

func HitungBMI(berat, tinggi float64) (float64, string) {
	tinggiMeter := tinggi / 100
	bmi := berat / (tinggiMeter * tinggiMeter)

	var status string
	if bmi < 18.5 {
		status = "Kurus"
	} else if bmi < 25 {
		status = "Normal"
	} else if bmi >= 30 {
		status = "Obesitas"
	}

	return bmi, status
}
