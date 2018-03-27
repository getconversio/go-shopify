package goshopify

import (
	"fmt"
)

const checkoutsBasePath = "admin/checkouts"

// CheckoutService is an interface for interfacing with the checkout endpoints
// of the Shopify API.
// See: https://help.shopify.com/api/reference/checkout
// See also: https://help.shopify.com/api/tutorials/sell-through-the-checkout-api#creating-a-checkout
type CheckoutService interface {
	Get(string, interface{}) (*Checkout, error)
	Create(Checkout) (*Checkout, error)
	// Update(Checkout) (*Checkout, error)
}

// CheckoutServiceOp handles communication with the checkout related methods of
// the Shopify API.
type CheckoutServiceOp struct {
	client *Client
}

// Checkout represents a Shopify checkout
type Checkout struct {
	CreatedAt           string        `json:"created_at"`
	Currency            string        `json:"currency"`
	CustomerID          int           `json:"customer_id"`
	Email               string        `json:"email"`
	LocationID          int           `json:"location_id"`
	OrderID             int           `json:"order_id"`
	RequiresShipping    bool          `json:"requires_shipping"`
	ReservationTime     int           `json:"reservation_time"`
	SourceName          string        `json:"source_name"`
	SourceIdentifier    interface{}   `json:"source_identifier"`
	SourceURL           string        `json:"source_url"`
	TaxesIncluded       bool          `json:"taxes_included"`
	Token               string        `json:"token"`
	UpdatedAt           string        `json:"updated_at"`
	PaymentDue          string        `json:"payment_due"`
	PaymentURL          string        `json:"payment_url"`
	ReservationTimeLeft int           `json:"reservation_time_left"`
	SubtotalPrice       string        `json:"subtotal_price"`
	TotalPrice          string        `json:"total_price"`
	TotalTax            string        `json:"total_tax"`
	Attributes          []interface{} `json:"attributes"`
	Note                string        `json:"note"`
	Order               interface{}   `json:"order"`
	PrivacyPolicyURL    string        `json:"privacy_policy_url"`
	RefundPolicyURL     string        `json:"refund_policy_url"`
	TermsOfServiceURL   string        `json:"terms_of_service_url"`
	UserID              int           `json:"user_id"`
	WebURL              string        `json:"web_url"`
	TaxLines            []struct {
		Price string  `json:"price"`
		Rate  float64 `json:"rate"`
		Title string  `json:"title"`
	} `json:"tax_lines"`
	LineItems       []CheckoutLineItem `json:"line_items"`
	GiftCards       []interface{}      `json:"gift_cards"`
	ShippingRate    interface{}        `json:"shipping_rate"`
	ShippingAddress *Address           `json:"shipping_address"`
	CreditCard      interface{}        `json:"credit_card"`
	BillingAddress  *Address           `json:"billing_address"`
	Discount        interface{}        `json:"discount"`
}

// Represents the result from the checkouts/X.json endpoint
type CheckoutResource struct {
	Checkout *Checkout `json:"checkout"`
}

// Represents the result from the checkouts.json endpoint
type CheckoutsResource struct {
	Checkouts []Checkout `json:"checkouts"`
}

// Get individual checkout
func (s *CheckoutServiceOp) Get(checkoutToken string, options interface{}) (*Checkout, error) {
	path := fmt.Sprintf("%s/%s.json", checkoutsBasePath, checkoutToken)
	resource := new(CheckoutResource)
	err := s.client.Get(path, resource, options)
	return resource.Checkout, err
}

// Create a new checkout
func (s *CheckoutServiceOp) Create(checkout Checkout) (*Checkout, error) {
	path := fmt.Sprintf("%s.json", checkoutsBasePath)
	wrappedData := CheckoutResource{Checkout: &checkout}
	resource := new(CheckoutResource)
	err := s.client.Post(path, wrappedData, resource)
	return resource.Checkout, err
}

// // Update an existing checkout
// func (s *CheckoutServiceOp) Update(checkout Checkout) (*Checkout, error) {
// 	path := fmt.Sprintf("%s/%d.json", checkoutsBasePath, checkout.ID)
// 	wrappedData := CheckoutResource{Checkout: &checkout}
// 	resource := new(CheckoutResource)
// 	err := s.client.Put(path, wrappedData, resource)
// 	return resource.Checkout, err
// }
