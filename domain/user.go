// domain/user.go
package domain

type DBType string

type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

type UserRepository interface {
	Create(user *User) error
	GetByID(id uint) (*User, error)
	GetByEmail(email string) (*User, error)
	Update(user *User) error
	Delete(id uint) error
}

type UserUsecase interface {
	Register(user *User) error
	Login(email, password string) (string, error)
	GetUser(id uint) (*User, error)
	UpdateUser(user *User) error
	DeleteUser(id uint) error
}
