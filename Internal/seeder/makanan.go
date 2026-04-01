package seeder

import (
	"bcc-geazy/internal/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func InfoMakanan(db *gorm.DB) {
	makanan := []entity.Makanan{
		{
			Id:      uuid.New(),
			Nama:    "Dada Ayam",
			Energi:  165,
			Protein: 31,
			Lemak:   3.6,
			Karbo:   0,
		},
		{
			Id:      uuid.New(),
			Nama:    "Paha Atas Ayam",
			Energi:  209,
			Protein: 26,
			Lemak:   11,
			Karbo:   0,
		},
		{
			Id:      uuid.New(),
			Nama:    "Hati Ayam",
			Energi:  167,
			Protein: 24,
			Lemak:   4.8,
			Karbo:   0.9,
		},
		{
			Id:      uuid.New(),
			Nama:    "Ceker Ayam",
			Energi:  215,
			Protein: 19,
			Lemak:   15,
			Karbo:   0.2,
		},
		{
			Id:      uuid.New(),
			Nama:    "Daging Sapi Giling",
			Energi:  217,
			Protein: 26,
			Lemak:   12,
			Karbo:   0,
		},
		{
			Id:      uuid.New(),
			Nama:    "Hati Sapi",
			Energi:  135,
			Protein: 20,
			Lemak:   3.6,
			Karbo:   3.9,
		},
		{
			Id:      uuid.New(),
			Nama:    "Ikan Lele",
			Energi:  105,
			Protein: 18,
			Lemak:   2.9,
			Karbo:   0,
		},
		{
			Id:      uuid.New(),
			Nama:    "Ikan Kembung",
			Energi:  167,
			Protein: 19,
			Lemak:   9.4,
			Karbo:   0,
		},
		{
			Id:      uuid.New(),
			Nama:    "Ikan Tuna",
			Energi:  132,
			Protein: 28,
			Lemak:   1.3,
			Karbo:   0,
		},
		{
			Id:      uuid.New(),
			Nama:    "Ikan Salmon",
			Energi:  208,
			Protein: 20,
			Lemak:   13,
			Karbo:   0,
		},
		{
			Id:      uuid.New(),
			Nama:    "Udang",
			Energi:  99,
			Protein: 24,
			Lemak:   0.3,
			Karbo:   0.2,
		},
		{
			Id:      uuid.New(),
			Nama:    "Telur Ayam Negeri",
			Energi:  155,
			Protein: 13,
			Lemak:   11,
			Karbo:   1.1,
		},
		{
			Id:      uuid.New(),
			Nama:    "Telur Ayam Kampung",
			Energi:  150,
			Protein: 13,
			Lemak:   10,
			Karbo:   1.2,
		},
		{
			Id:      uuid.New(),
			Nama:    "Telur Puyuh",
			Energi:  158,
			Protein: 13,
			Lemak:   11,
			Karbo:   0.4,
		},
		{
			Id:      uuid.New(),
			Nama:    "Tempe",
			Energi:  192,
			Protein: 20,
			Lemak:   11,
			Karbo:   9,
		},
		{
			Id:      uuid.New(),
			Nama:    "Tahu",
			Energi:  76,
			Protein: 8,
			Lemak:   4.8,
			Karbo:   1.9,
		},
		{
			Id:      uuid.New(),
			Nama:    "Kacang Hijau",
			Energi:  347,
			Protein: 24,
			Lemak:   1.2,
			Karbo:   63,
		},
		{
			Id:      uuid.New(),
			Nama:    "Kacang Merah",
			Energi:  333,
			Protein: 24,
			Lemak:   0.8,
			Karbo:   60,
		},
		{
			Id:      uuid.New(),
			Nama:    "Nasi Putih",
			Energi:  130,
			Protein: 2.7,
			Lemak:   0.3,
			Karbo:   28,
		},
		{
			Id:      uuid.New(),
			Nama:    "Nasi Merah",
			Energi:  111,
			Protein: 2.6,
			Lemak:   0.9,
			Karbo:   23,
		},
		{
			Id:      uuid.New(),
			Nama:    "Kentang",
			Energi:  87,
			Protein: 1.9,
			Lemak:   0.1,
			Karbo:   20,
		},
		{
			Id:      uuid.New(),
			Nama:    "Ubi Jalar",
			Energi:  86,
			Protein: 1.6,
			Lemak:   0.1,
			Karbo:   20,
		},
		{
			Id:      uuid.New(),
			Nama:    "Singkong",
			Energi:  160,
			Protein: 1.4,
			Lemak:   0.3,
			Karbo:   38,
		},
		{
			Id:      uuid.New(),
			Nama:    "Wortel",
			Energi:  41,
			Protein: 0.9,
			Lemak:   0.2,
			Karbo:   9.6,
		},
		{
			Id:      uuid.New(),
			Nama:    "Bayam",
			Energi:  23,
			Protein: 2.9,
			Lemak:   0.4,
			Karbo:   3.6,
		},
		{
			Id:      uuid.New(),
			Nama:    "Brokoli",
			Energi:  34,
			Protein: 2.8,
			Lemak:   0.4,
			Karbo:   6.6,
		},
	}

	for _, infomakanan := range makanan {
		var info entity.Makanan

		err := db.Where("nama = ?", infomakanan.Nama).First(&info).Error

		if err == gorm.ErrRecordNotFound {
			db.Create(&infomakanan)
		}
	}
}
