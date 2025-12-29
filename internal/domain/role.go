package domain

type Role struct {
	ID    int
	Name  string
	Label string
}

type RoleRepository interface {
	GetAll() ([]*Role, error)
	GetByUserID(userID string) ([]*Role, error)
	GetByName(name string) (*Role, error)
	GetByID(id int) (*Role, error)
}
