package model

type TambahkanDataAnak struct {
	Nama       string  `json:"nama" binding:"required"`
	Tinggi     float64 `json:"tinggi" binding:"required"`
	BeratBadan float64 `json:"berat_badan" binding:"required"`
	Gender     string  `json:"gender" binding:"required"`
}
