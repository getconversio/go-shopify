package goshopify

import (
	"fmt"
	"time"
)

// FulfillmentService is an interface for interfacing with the fulfillment endpoints
// of the Shopify API.
// https://help.shopify.com/api/reference/fulfillment
type FulfillmentService interface {
	List(interface{}) ([]Fulfillment, error)
	Count(interface{}) (int, error)
	Get(int, interface{}) (*Fulfillment, error)
	Create(Fulfillment) (*Fulfillment, error)
	Update(Fulfillment) (*Fulfillment, error)
	Complete(int) (*Fulfillment, error)
	Transition(int) (*Fulfillment, error)
	Cancel(int) (*Fulfillment, error)
}

// FulfillmentsService is an interface for other Shopify resources
// to interface with the fulfillment endpoints of the Shopify API.
// https://help.shopify.com/api/reference/fulfillment
type FulfillmentsService interface {
	ListFulfillments(int, interface{}) ([]Fulfillment, error)
	CountFulfillments(int, interface{}) (int, error)
	GetFulfillment(int, int, interface{}) (*Fulfillment, error)
	CreateFulfillment(int, Fulfillment) (*Fulfillment, error)
	UpdateFulfillment(int, Fulfillment) (*Fulfillment, error)
	CompleteFulfillment(int, int) (*Fulfillment, error)
	TransitionFulfillment(int, int) (*Fulfillment, error)
	CancelFulfillment(int, int) (*Fulfillment, error)
}

// FulfillmentServiceOp handles communication with the fulfillment
// related methods of the Shopify API.
type FulfillmentServiceOp struct {
	client     *Client
	resource   string
	resourceID int
}

// Fulfillment represents a Shopify fulfillment.
type Fulfillment struct {
	ID              int        `json:"id,omitempty"`
	OrderID         int        `json:"order_id,omitempty"`
	LocationID      int        `json:"location_id,omitempty"`
	Status          string     `json:"status,omitempty"`
	CreatedAt       *time.Time `json:"created_at,omitempty"`
	Service         string     `json:"service,omitempty"`
	UpdatedAt       *time.Time `json:"updated_at,omitempty"`
	TrackingCompany string     `json:"tracking_company,omitempty"`
	ShipmentStatus  string     `json:"shipment_status,omitempty"`
	TrackingNumber  string     `json:"tracking_number,omitempty"`
	TrackingNumbers []string   `json:"tracking_numbers,omitempty"`
	TrackingUrl     string     `json:"tracking_url,omitempty"`
	TrackingUrls    []string   `json:"tracking_urls,omitempty"`
	Receipt         Receipt    `json:"receipt,omitempty"`
	LineItems       []LineItem `json:"line_items,omitempty"`
	NotifyCustomer  bool       `json:"notify_customer,omitempty"`
}

// Receipt represents a Shopify receipt.
type Receipt struct {
	TestCase      bool   `json:"testcase,omitempty"`
	Authorization string `json:"authorization,omitempty"`
}

// FulfillmentResource represents the result from the fulfillments/X.json endpoint
type FulfillmentResource struct {
	Fulfillment *Fulfillment `json:"fulfillment"`
}

// FulfillmentsResource represents the result from the fullfilments.json endpoint
type FulfillmentsResource struct {
	Fulfillments []Fulfillment `json:"fulfillments"`
}

// List fulfillments
func (s *FulfillmentServiceOp) List(options interface{}) ([]Fulfillment, error) {
	prefix := FulfillmentPathPrefix(s.resource, s.resourceID)
	path := fmt.Sprintf("%s.json", prefix)
	resource := new(FulfillmentsResource)
	err := s.client.Get(path, resource, options)
	return resource.Fulfillments, err
}

// Count fulfillments
func (s *FulfillmentServiceOp) Count(options interface{}) (int, error) {
	prefix := FulfillmentPathPrefix(s.resource, s.resourceID)
	path := fmt.Sprintf("%s/count.json", prefix)
	return s.client.Count(path, options)
}

// Get individual fulfillment
func (s *FulfillmentServiceOp) Get(fulfillmentID int, options interface{}) (*Fulfillment, error) {
	prefix := FulfillmentPathPrefix(s.resource, s.resourceID)
	path := fmt.Sprintf("%s/%d.json", prefix, fulfillmentID)
	resource := new(FulfillmentResource)
	err := s.client.Get(path, resource, options)
	return resource.Fulfillment, err
}

// Create a new fulfillment
func (s *FulfillmentServiceOp) Create(fulfillment Fulfillment) (*Fulfillment, error) {
	prefix := FulfillmentPathPrefix(s.resource, s.resourceID)
	path := fmt.Sprintf("%s.json", prefix)
	wrappedData := FulfillmentResource{Fulfillment: &fulfillment}
	resource := new(FulfillmentResource)
	err := s.client.Post(path, wrappedData, resource)
	return resource.Fulfillment, err
}

// Update an existing fulfillment
func (s *FulfillmentServiceOp) Update(fulfillment Fulfillment) (*Fulfillment, error) {
	prefix := FulfillmentPathPrefix(s.resource, s.resourceID)
	path := fmt.Sprintf("%s/%d.json", prefix, fulfillment.ID)
	wrappedData := FulfillmentResource{Fulfillment: &fulfillment}
	resource := new(FulfillmentResource)
	err := s.client.Put(path, wrappedData, resource)
	return resource.Fulfillment, err
}

// Complete an existing fulfillment
func (s *FulfillmentServiceOp) Complete(fulfillmentID int) (*Fulfillment, error) {
	prefix := FulfillmentPathPrefix(s.resource, s.resourceID)
	path := fmt.Sprintf("%s/%d.json", prefix, fulfillmentID)
	resource := new(FulfillmentResource)
	err := s.client.Post(path, nil, resource)
	return resource.Fulfillment, err
}

// Transition an existing fulfillment
func (s *FulfillmentServiceOp) Transition(fulfillmentID int) (*Fulfillment, error) {
	prefix := FulfillmentPathPrefix(s.resource, s.resourceID)
	path := fmt.Sprintf("%s/%d.json", prefix, fulfillmentID)
	resource := new(FulfillmentResource)
	err := s.client.Post(path, nil, resource)
	return resource.Fulfillment, err
}

// Cancel an existing fulfillment
func (s *FulfillmentServiceOp) Cancel(fulfillmentID int) (*Fulfillment, error) {
	prefix := FulfillmentPathPrefix(s.resource, s.resourceID)
	path := fmt.Sprintf("%s/%d.json", prefix, fulfillmentID)
	resource := new(FulfillmentResource)
	err := s.client.Post(path, nil, resource)
	return resource.Fulfillment, err
}
