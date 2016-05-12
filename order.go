package goshopify

import (
	"fmt"
)

const ordersBasePath = "admin/orders"

// OrderService is an interface for interfacing with the orders endpoints of
// the Shopify API.
// See: https://help.shopify.com/api/reference/order
type OrderService interface {
	Count() (int, error)
}

// OrderServiceOp handles communication with the order related methods of the
// Shopify API.
type OrderServiceOp struct {
	client *Client
}

type orderCountRoot struct {
	Count int `json:"count"`
}

// Count orders
func (s *OrderServiceOp) Count() (int, error) {
	path := fmt.Sprintf("%s/count.json", ordersBasePath)

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return 0, err
	}

	root := new(orderCountRoot)
	err = s.client.Do(req, root)
	if err != nil {
		return 0, err
	}

	return root.Count, err
}
