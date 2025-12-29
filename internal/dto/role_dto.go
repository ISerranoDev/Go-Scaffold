package dto

import "go-react-backend.iserranodev.net/internal/domain"

type RoleResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Label string `json:"label"`
}

func FromDomainRole(r *domain.Role) RoleResponse {
	return RoleResponse{
		ID:    r.ID,
		Name:  r.Name,
		Label: r.Label,
	}
}
