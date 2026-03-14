package model

type GoogleUser struct {
	Id    string `json:"id" binding:"required"`
	Nama  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}
