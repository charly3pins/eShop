package domain

import (
	"errors"
	"time"

	"github.com/charly3pins/eShop/domain/base"
)

func NewUser(id base.Identity, email, name string) User {
	return User{
		Identity: id,
		Email:    email,
		Name:     name,
	}
}

type User struct {
	base.Identity
	Email     string
	Name      string
	CreatedAt time.Time
}

func (u User) Validate() error {
	if u.Email == "" {
		return errors.New("email cannot be empty")
	}
	if u.Name == "" {
		return errors.New("name cannot be empty")
	}

	return nil
}

type UserRepository interface {
	base.Repository
	Save(u User) (User, error)
	GetUserByID(id base.Identity) (*User, error)
	GetUserByEmail(email string) (*User, error)
}
