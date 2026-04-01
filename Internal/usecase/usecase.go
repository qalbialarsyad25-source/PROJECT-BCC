package usecase

import (
	websocket "bcc-geazy/internal/controller/delivery"
	"bcc-geazy/internal/repository"
	"bcc-geazy/pkg/bcrypt"
	"bcc-geazy/pkg/jwt"

	"golang.org/x/oauth2"
)

type Usecase struct {
	AnakUsecase         IAnakUsecase
	MakananUsecase      IMakananUsecase
	LogUsecase          ILogUsecase
	LogMakananUsecase   ILogMakananUsecase
	DokterUsecase       IDokterUsecase
	KonsultasiUsecase   IKonsultasiUsecase
	InformasiUsecase    IInformasiUsecase
	AuthUsecase         IAuthUsecase
	NutrisiUsecase      INutrisiUsecase
	LogInformasiUsecase ILogInformasiUsecase
	NotifikasiUsecase   INotifikasiUsecase
	UserUsecase         IUserUsecase
}

func NewUsecase(jwt *jwt.JWT, bcrypt bcrypt.IBcrypt, oauth *oauth2.Config, repository *repository.Repository, ws *websocket.WSManager) *Usecase {
	return &Usecase{
		AnakUsecase:         NewAnakUsecase(repository.AnakRepository),
		MakananUsecase:      NewMakananUsecase(repository.MakananRepository),
		LogUsecase:          NewLogUsecase(repository.LogRepository),
		LogMakananUsecase:   NewLogMakananUsecase(repository.LogMakananRepository),
		DokterUsecase:       NewDokterUsecase(repository.DokterRepository, repository.UserRepository, bcrypt),
		KonsultasiUsecase:   NewKonsultasiUsecase(repository.KonsultasiRepository, ws),
		InformasiUsecase:    NewInformasiUsecase(repository.InformasiRepository),
		AuthUsecase:         NewAuthUsecase(jwt, bcrypt, oauth, repository.UserRepository),
		NutrisiUsecase:      NewNutrisiUsecase(repository.LogRepository, repository.AnakRepository),
		LogInformasiUsecase: NewLogInformasiUsecase(repository.LogInformasiRepository),
		NotifikasiUsecase:   NewNotifikasiUsecase(repository.NotifikasiRepository),
		UserUsecase:         NewUserUsecase(repository.UserRepository),
	}
}
