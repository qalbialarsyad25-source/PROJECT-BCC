package usecase

import (
	"bcc-geazy/internal/repository"
	"bcc-geazy/pkg/bcrypt"
	"bcc-geazy/pkg/jwt"

	"golang.org/x/oauth2"
)

type Usecase struct {
	UserUsecase       IUserUsecase
	AnakUsecase       IAnakUsecase
	MakananUsecase    IMakananUsecase
	LogUsecase        ILogUsecase
	DokterUsecase     IDokterUsecase
	KonsultasiUsecase IKonsulUsecase
	InformasiUsecase  IInformasiUsecase
	AuthUsecase       IAuthUsecase
}

func NewUsecase(jwt jwt.JWT, bcrypt bcrypt.IBcrypt, oauth *oauth2.Config, repository *repository.Repository) *Usecase {
	return &Usecase{
		AnakUsecase:       NewAnakUsecase(repository.AnakRepository),
		MakananUsecase:    NewMakananUsecase(repository.MakananRepository),
		LogUsecase:        NewLogUsecase(repository.LogRepository),
		DokterUsecase:     NewDokterUsecase(repository.DokterRepository),
		KonsultasiUsecase: NewKonsulUsecase(repository.KonsultasiRepository),
		InformasiUsecase:  NewInformasiUsecase(repository.InformasiRepository),
		AuthUsecase:       NewAuthUsecase(jwt, bcrypt, oauth, repository.UserRepository),
	}
}
