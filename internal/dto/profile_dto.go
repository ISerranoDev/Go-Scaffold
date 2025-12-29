package dto

import "go-scaffold.iserranodev.net/internal/domain"

type UpdateProfileRequest struct {
	ID         string `json:"id"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	RePassword string `json:"re_password"`
}

type ProfileResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

func FromDomainProfile(u *domain.User) ProfileResponse {

	return ProfileResponse{
		ID:    u.ID,
		Email: u.Email,
	}
}
