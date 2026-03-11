package repository

import "gorm.io/gorm"

type Repository struct {
	UserRepository       IUserRepository
	AnakRepository       IAnakRepository
	MakananRepository    IMakananRepository
	LogRepository        ILogRepository
	DokterRepository     IDokterRepository
	KonsultasiRepository IKonsulRepository
	InformasiRepository  IInformasiRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		UserRepository:       NewUserRepository(db),
		AnakRepository:       NewAnakRepository(db),
		MakananRepository:    NewMakananRepository(db),
		LogRepository:        NewLogRepository(db),
		DokterRepository:     NewDokterRepository(db),
		KonsultasiRepository: NewKonsulRepository(db),
		InformasiRepository:  NewInformasiRepository(db),
	}
}
