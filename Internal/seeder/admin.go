package seeder

import (
	"bcc-geazy/internal/entity"
	"bcc-geazy/internal/repository"
	"context"
	"os"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func SeedAdmin(repo repository.IUserRepository) {
	adminEmail := os.Getenv("EMAIL_ADMIN")
	adminSandi := os.Getenv("SANDI_ADMIN")

	existing, _ := repo.GetUserByEmail(context.Background(), adminEmail)
	if existing != nil {
		return
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(adminSandi), bcrypt.DefaultCost)

	admin := entity.User{
		Id:       uuid.New(),
		Nama:     "Admin",
		Email:    adminEmail,
		Password: string(hash),
		Role:     "admin",
	}

	repo.CreateUser(context.Background(), admin)
}
