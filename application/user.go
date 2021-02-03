package application

import (
	"errors"

	"github.com/charly3pins/eShop/domain"
)

var (
	ErrUserAlreadyExists = errors.New("user already exists")
)

type UserService struct {
	UserRepository domain.UserRepository
}

type SignInUserRequest struct {
	Email string
	Name  string
}

func (us UserService) SignInUser(req SignInUserRequest) (domain.User, error) {
	var u domain.User

	uDB, err := us.UserRepository.GetUserByEmail(req.Email)
	if err != nil {
		return u, err
	}
	if uDB.ID.Valid {
		return u, ErrUserAlreadyExists
	}

	u = domain.NewUser(
		us.UserRepository.NextIdentity(),
		req.Email,
		req.Name,
	)

	if err := u.Validate(); err != nil {
		return u, err
	}

	return us.UserRepository.Save(u)
}
