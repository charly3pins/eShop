package application

import (
	"errors"

	"github.com/charly3pins/eShop/domain"
	"github.com/charly3pins/eShop/domain/base"
	"github.com/gofrs/uuid"
)

var (
	ErrInvalidUserID          = errors.New("invalid user id")
	ErrInvalidOrderID         = errors.New("invalid order id")
	ErrOrderWithEmptyProducts = errors.New("order cannot have empty products")
	ErrOrderAlreadyProcessed  = errors.New("order already processed")
)

type OrderService struct {
	OrderRepository domain.OrderRepository
}

type ProductRequest struct {
	ID        string
	Name      string
	UnitPrice int
	Currency  string
	Quantity  int
}

type AddProductsToOrderRequest struct {
	UserID   string
	OrderID  string
	Products []ProductRequest
}

func (os OrderService) AddProductsToOrder(req AddProductsToOrderRequest) (domain.Order, error) {
	var o domain.Order

	reqUserID, err := uuid.FromString(req.UserID)
	if err != nil {
		return o, ErrInvalidUserID
	}
	if len(req.Products) == 0 {
		return o, ErrOrderWithEmptyProducts
	}

	if req.OrderID == "" {
		o, err = createOrder(os, req, reqUserID)
	} else {
		o, err = updateOrder(os, req)
	}

	return o, nil
}

func createOrder(os OrderService, req AddProductsToOrderRequest, reqUserID uuid.UUID) (domain.Order, error) {
	o := domain.NewOrder(
		os.OrderRepository.NextIdentity(),
		reqUserID,
	)

	addProducts(&o, req.Products, os)

	if err := o.Validate(); err != nil {
		return o, err
	}
	oDB, err := os.OrderRepository.Create(o)
	if err != nil {
		return o, err
	}
	return oDB, nil
}

func updateOrder(os OrderService, req AddProductsToOrderRequest) (domain.Order, error) {
	var o domain.Order

	reqOrderID, err := uuid.FromString(req.OrderID)
	if err == nil {
		id := base.Identity{ID: uuid.NullUUID{UUID: reqOrderID, Valid: true}}
		oDB, err := os.OrderRepository.GetOrderByID(id)
		if err != nil {
			return o, err
		}
		if oDB == nil {
			return o, ErrInvalidOrderID
		}
		if oDB.Processed {
			return o, ErrOrderAlreadyProcessed
		}
		o = *oDB
	}

	addProducts(&o, req.Products, os)

	if err := o.Validate(); err != nil {
		return o, err
	}
	o, err = os.OrderRepository.Update(o)
	if err != nil {
		return o, err
	}
	return o, nil
}

func addProducts(o *domain.Order, products []ProductRequest, os OrderService) error {
	for _, p := range products {
		var id base.Identity
		if p.ID == "" {
			id = os.OrderRepository.NextIdentity()
		} else {
			pUUID, err := uuid.FromString(p.ID)
			if err != nil {
				return err
			}
			id.ID = uuid.NullUUID{UUID: pUUID, Valid: true}
		}
		err := o.AddProduct(id, p.Name, p.Currency, p.UnitPrice, p.Quantity)
		if err != nil {
			return err
		}
	}
	return nil
}
