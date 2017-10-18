package goshopify

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

const customersBasePath = "admin/customers"

// CustomerService is an interface for interfacing with the customers endpoints
// of the Shopify API.
// See: https://help.shopify.com/api/reference/customer
type CustomerService interface {
	List(interface{}) ([]Customer, error)
	Count(interface{}) (int, error)
	Get(int, interface{}) (*Customer, error)
}

// CustomerServiceOp handles communication with the product related methods of
// the Shopify API.
type CustomerServiceOp struct {
	client *Client
}

// Customer represents a Shopify customer
type Customer struct {
	ID                  int              `json:"id"`
	Email               string           `json:"email"`
	FirstName           string           `json:"first_name"`
	LastName            string           `json:"last_name"`
	State               string           `json:"state"`
	Note                string           `json:"note"`
	VerifiedEmail       bool             `json:"verified_email"`
	MultipassIdentifier string           `json:"multipass_identifier"`
	OrdersCount         int              `json:"orders_count"`
	TaxExempt           bool             `json:"tax_exempt"`
	TotalSpent          *decimal.Decimal `json:"total_spent"`
	Phone               string           `json:"phone"`
	Tags                string           `json:"tags"`
	LastOrderId         int              `json:"last_order_id"`
	AcceptsMarketing    bool             `json:"accepts_marketing"`
	CreatedAt           *time.Time       `json:"created_at"`
	UpdatedAt           *time.Time       `json:"updated_at"`
}

// Represents the result from the customers/X.json endpoint
type CustomerResource struct {
	Customer *Customer `json:"customer"`
}

// Represents the result from the customers.json endpoint
type CustomersResource struct {
	Customers []Customer `json:"customers"`
}

// List customers
func (s *CustomerServiceOp) List(options interface{}) ([]Customer, error) {
	path := fmt.Sprintf("%s.json", customersBasePath)
	resource := new(CustomersResource)
	err := s.client.Get(path, resource, options)
	return resource.Customers, err
}

// Count customers
func (s *CustomerServiceOp) Count(options interface{}) (int, error) {
	path := fmt.Sprintf("%s/count.json", customersBasePath)
	return s.client.Count(path, options)
}

// Get customer
func (s *CustomerServiceOp) Get(customerID int, options interface{}) (*Customer, error) {
	path := fmt.Sprintf("%s/%v.json", customersBasePath, customerID)
	resource := new(CustomerResource)
	err := s.client.Get(path, resource, options)
	return resource.Customer, err
}
