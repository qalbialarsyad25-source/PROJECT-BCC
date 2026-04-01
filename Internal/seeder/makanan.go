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
			Nama:    "Ayam Dada",
			Energi:  165,
			Protein: 31,
			Lemak:   3.6,
			Karbo:   0,
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
			Nama:    "Tempe",
			Energi:  192,
			Protein: 19,
			Lemak:   11,
			Karbo:   9.4,
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
			Nama:    "Telur",
			Energi:  155,
			Protein: 13,
			Lemak:   11,
			Karbo:   1.1,
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
