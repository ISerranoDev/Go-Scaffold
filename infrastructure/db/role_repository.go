package db

import (
	"go-react-backend.iserranodev.net/internal/domain"
	"gorm.io/gorm"
)

type RoleModel struct {
	ID    int    `gorm:"primaryKey"`
	Name  string `gorm:"unique"`
	Label string `gorm:"not null"`
}

func (RoleModel) TableName() string {
	return "roles"
}

type RoleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{db}
}

func (r *RoleRepository) GetByUserID(userID string) ([]*domain.Role, error) {
	var user UserModel
	err := r.db.Preload("Roles").First(&user, "id = ?", userID).Error
	if err != nil {
		return nil, err
	}
	var roles []*domain.Role
	for _, role := range user.Roles {
		roles = append(roles, &domain.Role{ID: role.ID, Name: role.Name, Label: role.Label})
	}
	return roles, nil
}

func (r *RoleRepository) GetAll() ([]*domain.Role, error) {
	var roleModels []RoleModel
	err := r.db.Find(&roleModels).Error
	if err != nil {
		return nil, err
	}

	var roles []*domain.Role
	for _, m := range roleModels {
		roles = append(roles, &domain.Role{ID: m.ID, Name: m.Name, Label: m.Label})
	}

	return roles, nil
}

func (r *RoleRepository) GetByName(name string) (*domain.Role, error) {
	var role RoleModel
	err := r.db.Where("name = ?", name).First(&role).Error
	if err != nil {
		return nil, err
	}
	return &domain.Role{
		ID:    role.ID,
		Name:  role.Name,
		Label: role.Label,
	}, nil
}

func (r *RoleRepository) GetByID(id int) (*domain.Role, error) {
	var role RoleModel
	err := r.db.Where("id = ?", id).First(&role).Error
	if err != nil {
		return nil, err
	}
	return &domain.Role{
		ID:    role.ID,
		Name:  role.Name,
		Label: role.Label,
	}, nil
}
