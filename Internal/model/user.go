package model

import (
	"bcc-geazy/internal/entity"
)

type BuatUser struct {
	Nama string `json:"nama"`
}

type EditUser struct {
	Nama string `json:"nama"`
}

func (p *EditUser) ToMap() map[string]any {
	perbaruan := map[string]any{}

	if p.Nama != "" {
		perbaruan["nama"] = p.Nama
	}

	return perbaruan
}

func toUserResponse(User entity.User) UserResponse {
	return UserResponse{
		Id:   User.Id,
		Nama: User.Nama,
	}
}

func toUserResponses(User []entity.User) []UserResponse {
	var responses []UserResponse
	for _, User := range User {
		responses = append(responses, toUserResponse(User))
	}

	return responses
}
