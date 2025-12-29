package usecase

import (
	"go-scaffold.iserranodev.net/internal/domain"
)

type RoleUseCase struct {
	roleRepo domain.RoleRepository
}

func NewRoleUseCase(r domain.RoleRepository) *RoleUseCase {
	return &RoleUseCase{r}
}

func (uc *RoleUseCase) GetAll() ([]*domain.Role, error) {
	return uc.roleRepo.GetAll()
}
