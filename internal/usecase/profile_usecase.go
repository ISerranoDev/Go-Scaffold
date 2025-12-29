package usecase

import (
	"errors"
	"go-scaffold.iserranodev.net/internal/domain"
	"go-scaffold.iserranodev.net/internal/dto"
	"go-scaffold.iserranodev.net/internal/helper"
)

type ProfileUseCase struct {
	userRepo domain.UserRepository
}

func NewProfileUseCase(u domain.UserRepository) *ProfileUseCase {
	return &ProfileUseCase{u}
}

func (uc *ProfileUseCase) GetByID(id string) (*domain.User, error) {
	return uc.userRepo.GetByID(id)
}

func (uc *ProfileUseCase) Update(input *dto.UpdateProfileRequest) (*domain.User, error) {

	user, err := uc.userRepo.GetByID(input.ID)
	if err != nil {
		return nil, errors.New("usuario no encontrado")
	}

	user.Email = input.Email

	if input.Password != "" || input.RePassword != "" {
		if input.Password != input.RePassword {
			return nil, errors.New("las contraseñas no coinciden")
		}
		hashedPassword, err := helper.HashPassword(input.Password)
		if err != nil {
			return nil, err
		}
		user.Password = hashedPassword
	}

	err = uc.userRepo.UpdateProfile(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
