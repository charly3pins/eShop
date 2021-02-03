package postgres

import (
	"github.com/charly3pins/eShop/domain"
	"github.com/charly3pins/eShop/domain/base"
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	Db *gorm.DB
}

func (ur UserRepository) Save(u domain.User) (domain.User, error) {
	if err := ur.Db.Create(&u).Error; err != nil {
		return u, err
	}
	return u, nil
}

func (ur UserRepository) GetUserByID(id base.Identity) (*domain.User, error) {
	var u domain.User
	if err := ur.Db.First(&u, "id = ?", id.ID).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (ur UserRepository) GetUserByEmail(email string) (*domain.User, error) {
	var u domain.User
	if err := ur.Db.First(&u, "email = ?", email).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
	}
	return &u, nil
}

func (ur UserRepository) NextIdentity() base.Identity {
	return base.Identity{
		ID: uuid.NullUUID{
			UUID:  uuid.Must(uuid.NewV4()),
			Valid: true,
		},
	}
}
