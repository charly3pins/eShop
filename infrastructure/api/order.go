package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/charly3pins/eShop/application"
)

type OrderHandler struct {
	OrderService application.OrderService
}

type ProductBody struct {
	ProductID string `json:"product_id,omitempty"`
	Name      string `json:"name"`
	UnitPrice int    `json:"unit_price"`
	Currency  string `json:"currency"`
	Quantity  int    `json:"quantity"`
}

type AddProductsToOrderBody struct {
	UserID   string        `json:"user_id"`
	OrderID  string        `json:"order_id,omitempty"`
	Products []ProductBody `json:"products"`
}

func (oh OrderHandler) AddProductsToOrder(w http.ResponseWriter, r *http.Request) {
	body := &AddProductsToOrderBody{}
	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	products := make([]application.ProductRequest, len(body.Products))
	for i, p := range body.Products {
		product := application.ProductRequest{
			ID:        p.ProductID,
			Name:      p.Name,
			UnitPrice: p.UnitPrice,
			Currency:  p.Currency,
			Quantity:  p.Quantity,
		}
		products[i] = product
	}
	req := application.AddProductsToOrderRequest{
		UserID:   body.UserID,
		OrderID:  body.OrderID,
		Products: products,
	}
	if _, err := oh.OrderService.AddProductsToOrder(req); err != nil {
		log.Printf("error adding products to an order: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
