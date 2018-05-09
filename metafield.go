package goshopify

import (
	"fmt"
	"time"
)

const metafieldsBasePath = "admin/metafields"

// MetafieldService is an interface for interfacing with the metafield endpoints
// of the Shopify API.
// https://help.shopify.com/api/reference/metafield
type MetafieldService interface {
	List(interface{}) ([]Metafield, error)
	Count(interface{}) (int, error)
	Get(int, interface{}) (*Metafield, error)
	Create(Metafield) (*Metafield, error)
	Update(Metafield) (*Metafield, error)
	Delete(int) error
}

// MetafieldServiceOp handles communication with the metafield
// related methods of the Shopify API.
type MetafieldServiceOp struct {
	client *Client
}

// Metafield represents a Shopify metafield.
type Metafield struct {
	ID            int        `json:"id,omitempty"`
	Key           string     `json:"key,omitempty"`
	Value         string     `json:"value,omitempty"`
	ValueType     string     `json:"value_type,omitempty"`
	Namespace     string     `json:"namespace,omitempty"`
	Description   string     `json:"description,omitempty"`
	OwnerId       int        `json:"owner_id,omitempty"`
	CreatedAt     *time.Time `json:"created_at,omitempty"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty"`
	OwnerResource string     `json:"owner_resource,omitempty"`
}

// MetafieldResource represents the result from the metafields/X.json endpoint
type MetafieldResource struct {
	Metafield *Metafield `json:"metafield"`
}

// MetafieldsResource represents the result from the metafields.json endpoint
type MetafieldsResource struct {
	Metafields []Metafield `json:"metafields"`
}

// List metafields
func (s *MetafieldServiceOp) List(options interface{}) ([]Metafield, error) {
	path := fmt.Sprintf("%s.json", metafieldsBasePath)
	resource := new(MetafieldsResource)
	err := s.client.Get(path, resource, options)
	return resource.Metafields, err
}

// Count metafields
func (s *MetafieldServiceOp) Count(options interface{}) (int, error) {
	path := fmt.Sprintf("%s/count.json", metafieldsBasePath)
	return s.client.Count(path, options)
}

// Get individual metafield
func (s *MetafieldServiceOp) Get(metafieldID int, options interface{}) (*Metafield, error) {
	path := fmt.Sprintf("%s/%d.json", metafieldsBasePath, metafieldID)
	resource := new(MetafieldResource)
	err := s.client.Get(path, resource, options)
	return resource.Metafield, err
}

// Create a new metafield
func (s *MetafieldServiceOp) Create(metafield Metafield) (*Metafield, error) {
	path := fmt.Sprintf("%s.json", metafieldsBasePath)
	wrappedData := MetafieldResource{Metafield: &metafield}
	resource := new(MetafieldResource)
	err := s.client.Post(path, wrappedData, resource)
	return resource.Metafield, err
}

// Update an existing metafield
func (s *MetafieldServiceOp) Update(metafield Metafield) (*Metafield, error) {
	path := fmt.Sprintf("%s/%d.json", metafieldsBasePath, metafield.ID)
	wrappedData := MetafieldResource{Metafield: &metafield}
	resource := new(MetafieldResource)
	err := s.client.Put(path, wrappedData, resource)
	return resource.Metafield, err
}

// Delete an existing metafield
func (s *MetafieldServiceOp) Delete(metafieldID int) error {
	return s.client.Delete(fmt.Sprintf("%s/%d.json", metafieldsBasePath, metafieldID))
}
