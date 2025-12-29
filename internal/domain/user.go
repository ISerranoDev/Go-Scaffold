package domain

type User struct {
	ID       string
	Email    string
	Password string
	Roles    []*Role
}

type UserRepository interface {
	GetByEmail(email string) (*User, error)
	GetByID(id string) (*User, error)
	List(filters map[string]interface{}, sort map[string]string, page, pageSize int) ([]*User, *int64, error)
	Create(*User) error
	Update(*User) error
	UpdateProfile(*User) error
	Delete(id string) error
}

type UserService interface {
	Create(user *User, roleIDs []int, rePassword string) (*User, error)
	GetByID(id int) (*User, error)
	Update(user *User, roleIDs []int, rePassword string) (*User, error)
	Delete(id int) error
	List(filters map[string]interface{}, sort map[string]string, page, pageSize int) ([]*User, error)
}
