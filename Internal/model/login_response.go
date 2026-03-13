package model

type LoginResponse struct {
	Token string       `json:"token" binding:"required"`
	User  UserResponse `json:"user" binding:"required"`
}
