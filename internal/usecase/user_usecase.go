package usecase

import (
	"errors"
	"fmt"
	"go-scaffold.iserranodev.net/internal/domain"
	"go-scaffold.iserranodev.net/internal/dto"
	"go-scaffold.iserranodev.net/internal/helper"
)

type UserUseCase struct {
	userRepo domain.UserRepository
	roleRepo domain.RoleRepository
}

func NewUserUseCase(u domain.UserRepository, r domain.RoleRepository) *UserUseCase {
	return &UserUseCase{u, r}
}

func (uc *UserUseCase) List(filter dto.UserFilter) ([]*domain.User, *int64, error) {
	filters := make(map[string]interface{})
	if filter.Filters.Email != "" {
		filters["email"] = filter.Filters.Email
	}
	if len(filter.Filters.Roles) > 0 {
		filters["roles"] = filter.Filters.Roles
	}

	return uc.userRepo.List(filters, filter.Sort, filter.Page, filter.PageSize)
}

func (uc *UserUseCase) Create(input *dto.CreateUserRequest) (*domain.User, error) {
	if input.Password != input.RePassword {
		return nil, errors.New("las contraseñas no coinciden")
	}

	hashedPassword, err := helper.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	roles := make([]*domain.Role, 0, len(input.Roles))
	for _, id := range input.Roles {
		role, err := uc.roleRepo.GetByID(id)
		if err != nil {
			return nil, errors.New("rol inválido con ID: " + fmt.Sprint(id))
		}
		roles = append(roles, role)
	}

	user := &domain.User{
		Email:    input.Email,
		Password: hashedPassword,
		Roles:    roles,
	}

	err = uc.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *UserUseCase) GetByID(id string) (*domain.User, error) {
	return uc.userRepo.GetByID(id)
}

func (uc *UserUseCase) Update(input *dto.UpdateUserRequest) (*domain.User, error) {
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

	roles := make([]*domain.Role, 0, len(input.Roles))
	for _, id := range input.Roles {
		role, err := uc.roleRepo.GetByID(id)
		if err != nil {
			return nil, errors.New("rol inválido con ID: " + fmt.Sprint(id))
		}
		roles = append(roles, role)
	}
	user.Roles = roles

	err = uc.userRepo.Update(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *UserUseCase) Delete(id string) error {
	return uc.userRepo.Delete(id)
}
