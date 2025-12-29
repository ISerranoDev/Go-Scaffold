package dto

import "go-scaffold.iserranodev.net/internal/domain"

type CreateUserRequest struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	RePassword string `json:"re_password"`
	Roles      []int  `json:"roles"`
}

type UpdateUserRequest struct {
	ID         string `json:"id"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	RePassword string `json:"re_password"`
	Roles      []int  `json:"roles"`
}

type UserFilter struct {
	Filters  FilterFields      `json:"filters"`
	Page     int               `json:"page"`
	PageSize int               `json:"pageSize"`
	Sort     map[string]string `json:"sort"`
}

type FilterFields struct {
	Email string `json:"email"`
	Roles []int  `json:"roles"`
}

type UserResponse struct {
	ID    string         `json:"id"`
	Email string         `json:"email"`
	Roles []RoleResponse `json:"roles"`
}

func FromDomainUser(u *domain.User) UserResponse {
	roles := make([]RoleResponse, len(u.Roles))
	for i, r := range u.Roles {
		roles[i] = FromDomainRole(r)
	}
	return UserResponse{
		ID:    u.ID,
		Email: u.Email,
		Roles: roles,
	}
}
