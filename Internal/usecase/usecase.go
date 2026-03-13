package usecase

import (
	"bcc-geazy/internal/repository"
)

type Usecase struct {
	UserUsecase       IUserUsecase
	AnakUsecase       IAnakUsecase
	MakananUsecase    IMakananUsecase
	LogUsecase        ILogUsecase
	DokterUsecase     IDokterUsecase
	KonsultasiUsecase IKonsulUsecase
	InformasiUsecase  IInformasiUsecase
}

func NewUsecase(repository *repository.Repository) *Usecase {
	return &Usecase{
		UserUsecase:       NewUserUsecase(repository.UserRepository),
		AnakUsecase:       NewAnakUsecase(repository.AnakRepository),
		MakananUsecase:    NewMakananUsecase(repository.MakananRepository),
		LogUsecase:        NewLogUsecase(repository.LogRepository),
		DokterUsecase:     NewDokterUsecase(repository.DokterRepository),
		KonsultasiUsecase: NewKonsulUsecase(repository.KonsultasiRepository),
		InformasiUsecase:  NewInformasiUsecase(repository.InformasiRepository),
	}
}
