package jadwal

import (
	"bcc-geazy/internal/model"
	"bcc-geazy/internal/repository"
	"bcc-geazy/internal/usecase"
	"bcc-geazy/pkg/email"
	"context"

	"github.com/robfig/cron/v3"
)

func StartCron(notifUsecase usecase.INotifikasiUsecase, userRepo repository.IUserRepository, anakRepo repository.IAnakRepository) {
	c := cron.New()

	c.AddFunc("@every 1m", func() {
		SimpleCheck(notifUsecase, userRepo, anakRepo)
	})

	c.Start()
}

func SimpleCheck(notifUsecase usecase.INotifikasiUsecase, userRepo repository.IUserRepository, anakRepo repository.IAnakRepository) {
	ctx := context.Background()

	users, err := userRepo.GetAll(ctx)
	if err != nil {
		return
	}

	for _, user := range users {

		anakList, _ := anakRepo.FindUserById(ctx, user.Id)

		if len(anakList) == 0 {

			pesan := "Halo Ayah & Bunda!Sedikit lagi nih! 👋 Lengkapi data si Kecil sekarang agar Geazy bisa bantu pantau tumbuh kembangnya secara akurat. Cuma butuh 1 menit kok!"

			_, err := notifUsecase.CreateNotifikasi(
				ctx,
				user.Id,
				model.BuatNotifikasi{
					Judul: "Lengkapi Data Anak",
					Pesan: pesan,
				},
			)

			if err == nil {
				_ = email.SendEmail(user.Email, "Lengkapi Data Anak", pesan)
			}

			continue
		}

		for _, anak := range anakList {

			if anak.ProteinHarian == 0 {

				pesan := "Nutrisi Harian Si Kecil sudah makan apa hari ini? 🍲 Nutrisi harian si Kecil masih kosong nih. Yuk, catat menu makannya sekarang agar Ayah & Bunda tahu apakah kebutuhan proteinnya sudah terpenuhi!"

				_, err := notifUsecase.CreateNotifikasi(
					ctx,
					user.Id,
					model.BuatNotifikasi{
						Judul: "Peringatan Nutrisi",
						Pesan: pesan,
					},
				)

				if err == nil {
					_ = email.SendEmail(user.Email, "Peringatan Nutrisi", pesan)
				}
			}
		}
	}
}
