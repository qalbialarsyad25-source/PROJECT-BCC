package model

type UserRegister struct {
	Nama            string `json:"nama" validate:"required"`
	UserName        string `json:"username" validate:"required,min=3"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=8"`
	ConfirmPassword string `json:"confirm_password" validate:"required"`
}

type UserLogin struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

var (
	UserRoleUser   = "user"
	UserRoleAdmin  = "admin"
	UserRoleDokter = "dokter"
)
