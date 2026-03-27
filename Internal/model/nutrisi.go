package model

type NutrisiHarianResponse struct {
	TotalKalori   float64 `json:"total_kalori"`
	TargetProtein float64 `json:"target_protein"`
	TargetLemak   float64 `json:"target_lemak"`
	TargetKalori  float64 `json:"target_kalori"`
	PersenKalori  float64 `json:"persen_kalori"`
	PersenLemak   float64 `json:"persen_lemak"`
	PersenProtein float64 `json:"persen_protein"`
}
