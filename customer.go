package goshopify

import (
	"fmt"
)

const customersBasePath = "admin/customers"

// CustomerService is an interface for interfacing with the customers endpoints
// of the Shopify API.
// See: https://help.shopify.com/api/reference/customer
type CustomerService interface {
	Count() (int, error)
}

// CustomerServiceOp handles communication with the product related methods of
// the Shopify API.
type CustomerServiceOp struct {
	client *Client
}

type customerCountRoot struct {
	Count int `json:"count"`
}

// Count customers
func (s *CustomerServiceOp) Count() (int, error) {
	path := fmt.Sprintf("%s/count.json", customersBasePath)

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return 0, err
	}

	root := new(customerCountRoot)
	err = s.client.Do(req, root)
	if err != nil {
		return 0, err
	}

	return root.Count, err
}
