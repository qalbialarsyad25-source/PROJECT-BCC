package seeder

import (
	"bcc-geazy/internal/entity"
	"bcc-geazy/internal/repository"
	"context"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func SeedAdmin(repo repository.IUserRepository) {
	adminEmail := "admin28@gmail.com"

	existing, _ := repo.GetUserByEmail(context.Background(), adminEmail)
	if existing != nil {
		return
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte("admin282828"), bcrypt.DefaultCost)

	admin := entity.User{
		Id:       uuid.New(),
		Nama:     "Admin",
		Email:    adminEmail,
		Password: string(hash),
		Role:     "admin",
	}

	repo.CreateUser(context.Background(), admin)
}
