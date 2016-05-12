package goshopify

import (
	"fmt"

	"github.com/shopspring/decimal"
)

const ordersBasePath = "admin/orders"

// OrderService is an interface for interfacing with the orders endpoints of
// the Shopify API.
// See: https://help.shopify.com/api/reference/order
type OrderService interface {
	List(options interface{}) ([]Order, error)
	Count(options interface{}) (int, error)
}

// OrderServiceOp handles communication with the order related methods of the
// Shopify API.
type OrderServiceOp struct {
	client *Client
}

// Order represents a Shopify order
type Order struct {
	ID    int             `json:"id"`
	Name  string          `json:"name"`
	Total decimal.Decimal `json:"total_price"`
}

// Represents the result from the orders/X.json endpoint
type OrderResource struct {
	Order *Order `json:"order"`
}

// Represents the result from the orders.json endpoint
type OrdersResource struct {
	Orders []Order `json:"orders"`
}

// List orders
func (s *OrderServiceOp) List(options interface{}) ([]Order, error) {
	path := fmt.Sprintf("%s.json", ordersBasePath)
	resource := new(OrdersResource)
	err := s.client.Get(path, resource, options)
	return resource.Orders, err
}

// Count orders
func (s *OrderServiceOp) Count(options interface{}) (int, error) {
	path := fmt.Sprintf("%s/count.json", ordersBasePath)
	return s.client.Count(path, options)
}
