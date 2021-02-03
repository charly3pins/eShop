package postgres

import (
	"github.com/charly3pins/eShop/domain"
	"github.com/charly3pins/eShop/domain/base"
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
)

type OrderRepository struct {
	Db *gorm.DB
}

func (or OrderRepository) Create(o domain.Order) (domain.Order, error) {
	if err := or.Db.Create(&o).Error; err != nil {
		return o, err
	}
	return o, nil
}

func (or OrderRepository) Update(o domain.Order) (domain.Order, error) {
	if err := or.Db.Save(&o).Error; err != nil {
		return o, err
	}
	return o, nil
}

func (or OrderRepository) GetOrderByID(id base.Identity) (*domain.Order, error) {
	var o domain.Order
	if err := or.Db.First(&o, "id = ?", id.ID).Error; err != nil {
		return nil, err
	}
	return &o, nil
}

func (or OrderRepository) NextIdentity() base.Identity {
	return base.Identity{
		ID: uuid.NullUUID{
			UUID:  uuid.Must(uuid.NewV4()),
			Valid: true,
		},
	}
}
