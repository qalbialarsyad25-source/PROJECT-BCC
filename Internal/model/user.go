package model

type BuatUser struct {
	Nama string `json:"nama"`
}

type EditUser struct {
	Nama string `json:"nama"`
}

type UserRegister struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=5"`
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
