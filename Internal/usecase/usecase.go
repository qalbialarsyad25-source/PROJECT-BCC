package usecase

import (
	"bcc-geazy/internal/repository"
	"bcc-geazy/pkg/bcrypt"
	"bcc-geazy/pkg/jwt"

	"golang.org/x/oauth2"
)

type Usecase struct {
	AnakUsecase       IAnakUsecase
	MakananUsecase    IMakananUsecase
	LogUsecase        ILogUsecase
	LogMakananUsecase ILogMakananUsecase
	DokterUsecase     IDokterUsecase
	KonsultasiUsecase IKonsultasiUsecase
	InformasiUsecase  IInformasiUsecase
	AuthUsecase       IAuthUsecase
}

func NewUsecase(jwt *jwt.JWT, bcrypt bcrypt.IBcrypt, oauth *oauth2.Config, repository *repository.Repository) *Usecase {
	return &Usecase{
		AnakUsecase:       NewAnakUsecase(repository.AnakRepository),
		MakananUsecase:    NewMakananUsecase(repository.MakananRepository),
		LogUsecase:        NewLogUsecase(repository.LogRepository),
		LogMakananUsecase: NewLogMakananUsecase(repository.LogMakananRepository),
		DokterUsecase:     NewDokterUsecase(repository.DokterRepository),
		KonsultasiUsecase: NewKonsultasiUsecase(repository.KonsultasiRepository),
		InformasiUsecase:  NewInformasiUsecase(repository.InformasiRepository),
		AuthUsecase:       NewAuthUsecase(jwt, bcrypt, oauth, repository.UserRepository),
	}
}
