package db

import (
	"github.com/oklog/ulid/v2"
	"go-scaffold.iserranodev.net/internal/domain"
	"gorm.io/gorm"
	"math/rand"
	"strings"
	"time"
)

type UserModel struct {
	ID       string `gorm:"primaryKey;size:26"`
	Email    string `gorm:"unique"`
	Password string
	Roles    []*RoleModel `gorm:"many2many:user_roles;joinForeignKey:ID;joinReferences:ID;joinForeignKey:user_id;joinReferences:role_id"`
}

func (u *UserModel) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == "" {
		t := time.Now()
		entropy := rand.New(rand.NewSource(t.UnixNano()))
		u.ID = ulid.MustNew(ulid.Timestamp(t), entropy).String()
	}
	return
}

func (UserModel) TableName() string {
	return "users"
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) GetByEmail(email string) (*domain.User, error) {
	var m UserModel
	err := r.db.Preload("Roles").Where("email = ?", email).First(&m).Error
	return toUserDomain(m), err
}

func toUserDomain(m UserModel) *domain.User {
	var roles []*domain.Role
	for _, rm := range m.Roles {
		roles = append(roles, &domain.Role{ID: rm.ID, Name: rm.Name, Label: rm.Label})
	}
	return &domain.User{
		ID:       m.ID,
		Email:    m.Email,
		Password: m.Password,
		Roles:    roles,
	}
}

func (r *UserRepository) GetByID(id string) (*domain.User, error) {
	var m UserModel
	err := r.db.Preload("Roles").Where("id = ?", id).First(&m).Error
	return toUserDomain(m), err
}

func (r *UserRepository) List(filters map[string]interface{}, sort map[string]string, page int, pageSize int) ([]*domain.User, *int64, error) {
	var userModels []UserModel

	query := r.db.Model(&UserModel{}).Preload("Roles")

	for field, value := range filters {
		switch field {
		case "roles":
			roles, ok := value.([]int)
			if ok && len(roles) > 0 {
				// JOIN con tabla intermedia user_roles (ajusta el nombre según tu modelo)
				query = query.Joins("JOIN user_roles ur ON ur.user_id = users.id").
					Where("ur.role_id IN ?", roles).
					Group("users.id") // importante para no duplicar resultados
			}
		case "email", "id":
			v, ok := value.(string)
			if ok && v != "" {
				query = query.Where(field+" ILIKE ?", "%"+v+"%")
			}
		}
	}

	// Contar total de resultados filtrados (sin paginación)
	var total int64
	err := query.Count(&total).Error
	if err != nil {
		return nil, nil, err
	}

	// Ordenamiento
	for field, order := range sort {
		order = strings.ToUpper(order)
		if order != "ASC" && order != "DESC" {
			order = "ASC"
		}
		query = query.Order(field + " " + order)
	}

	// Paginación
	if page > 0 && pageSize > 0 {
		offset := (page - 1) * pageSize
		query = query.Offset(offset).Limit(pageSize)
	}

	err = query.Find(&userModels).Error
	if err != nil {
		return nil, nil, err
	}

	var users []*domain.User
	for _, m := range userModels {
		users = append(users, toUserDomain(m))
	}

	return users, &total, nil
}

func (r *UserRepository) Create(u *domain.User) error {
	m := UserModel{
		Email:    u.Email,
		Password: u.Password,
	}

	// Si tiene roles, buscamos los modelos correspondientes
	if len(u.Roles) > 0 {
		var roleModels []*RoleModel
		for _, role := range u.Roles {
			roleModels = append(roleModels, &RoleModel{
				ID:    role.ID,
				Name:  role.Name,
				Label: role.Label,
			})
		}
		m.Roles = roleModels
	}

	err := r.db.Create(&m).Error
	if err != nil {
		return err
	}

	// Actualiza el ID generado en el domain.User recibido
	u.ID = m.ID

	return nil
}

func (r *UserRepository) Update(user *domain.User) error {
	var m UserModel
	err := r.db.Preload("Roles").First(&m, "id = ?", user.ID).Error
	if err != nil {
		return err
	}

	m.Email = user.Email
	m.Password = user.Password

	// Actualizar roles: buscamos los modelos de roles según domain.User.Roles
	var roleModels []*RoleModel
	for _, role := range user.Roles {
		roleModels = append(roleModels, &RoleModel{
			ID:    role.ID,
			Name:  role.Name,
			Label: role.Label,
		})
	}

	// Actualiza la asociación many2many
	err = r.db.Model(&m).Association("Roles").Replace(roleModels)
	if err != nil {
		return err
	}

	// Guarda los cambios del usuario
	return r.db.Save(&m).Error
}

func (r *UserRepository) UpdateProfile(user *domain.User) error {
	var m UserModel
	err := r.db.Preload("Roles").First(&m, "id = ?", user.ID).Error
	if err != nil {
		return err
	}

	m.Email = user.Email
	m.Password = user.Password

	// Guarda los cambios del usuario
	return r.db.Save(&m).Error
}

func (r *UserRepository) Delete(id string) error {
	var m UserModel
	err := r.db.Preload("Roles").First(&m, "id = ?", id).Error
	if err != nil {
		return err
	}

	// Limpia las relaciones many2many antes de borrar
	err = r.db.Model(&m).Association("Roles").Clear()
	if err != nil {
		return err
	}

	return r.db.Delete(&m).Error
}
