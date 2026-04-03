package seeder

import (
	"bcc-geazy/internal/entity"
	"bcc-geazy/pkg/bcrypt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeedDokter(db *gorm.DB, bc bcrypt.IBcrypt) {
	var count int64
	db.Model(&entity.Dokter{}).Count(&count)

	if count > 0 {
		return
	}
	hashed, err := bc.GenerateHash("123456")

	user := entity.User{
		Id:       uuid.New(),
		Email:    "dokter1@gmail.com",
		Password: hashed,
		Role:     "dokter",
	}

	err = db.Create(&user).Error
	if err != nil {
		panic(err)
	}

	dokter := entity.Dokter{
		Id:        uuid.New(),
		UserId:    user.Id,
		Nama:      "Dr. Andi",
		Spesialis: "Anak",
	}

	err = db.Create(&dokter).Error
	if err != nil {
		panic(err)
	}
}
