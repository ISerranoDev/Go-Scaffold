package usecase

import (
	"errors"
	"go-react-backend.iserranodev.net/internal/domain"
	"go-react-backend.iserranodev.net/internal/helper"
)

type AuthUseCase struct {
	userRepo domain.UserRepository
	roleRepo domain.RoleRepository
}

func NewAuthUseCase(u domain.UserRepository, r domain.RoleRepository) *AuthUseCase {
	return &AuthUseCase{u, r}
}

func (uc *AuthUseCase) Authenticate(email, password string) (*domain.User, error) {
	user, err := uc.userRepo.GetByEmail(email)

	if err != nil || !helper.CheckPassword(user.Password, password) {
		return nil, errors.New("credenciales no válidas")
	}
	user.Roles, _ = uc.roleRepo.GetByUserID(user.ID)
	return user, nil
}
