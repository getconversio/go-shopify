package goshopify

import (
	"fmt"
	"time"
)

const customCollectionsBasePath = "admin/custom_collections"

// CustomCollectionService is an interface for interacting with the custom
// collection endpoints of the Shopify API.
// See https://help.shopify.com/api/reference/customcollection
type CustomCollectionService interface {
	List(interface{}) ([]CustomCollection, error)
	Count(interface{}) (int, error)
	Get(int, interface{}) (*CustomCollection, error)
	Create(CustomCollection) (*CustomCollection, error)
	Update(CustomCollection) (*CustomCollection, error)
	Delete(int) error
}

// CustomCollectionServiceOp handles communication with the custom collection
// related methods of the Shopify API.
type CustomCollectionServiceOp struct {
	client *Client
}

// CustomCollection represents a Shopify custom collection.
type CustomCollection struct {
	ID             int        `json:"id"`
	Handle         string     `json:"handle"`
	Title          string     `json:"title"`
	UpdatedAt      *time.Time `json:"updated_at"`
	BodyHTML       string     `json:"body_html"`
	SortOrder      string     `json:"sort_order"`
	TemplateSuffix string     `json:"template_suffix"`
	Image          Image      `json:"image"`
	Published      bool       `json:"published"`
	PublishedAt    *time.Time `json:"published_at"`
	PublishedScope string     `json:"published_scope"`
}

// CustomCollectionResource represents the result form the custom_collections/X.json endpoint
type CustomCollectionResource struct {
	Collection *CustomCollection `json:"custom_collection"`
}

// CustomCollectionsResource represents the result from the custom_collections.json endpoint
type CustomCollectionsResource struct {
	Collections []CustomCollection `json:"custom_collections"`
}

// List custom collections
func (s *CustomCollectionServiceOp) List(options interface{}) ([]CustomCollection, error) {
	path := fmt.Sprintf("%s.json", customCollectionsBasePath)
	resource := new(CustomCollectionsResource)
	err := s.client.Get(path, resource, options)
	return resource.Collections, err
}

// Count custom collections
func (s *CustomCollectionServiceOp) Count(options interface{}) (int, error) {
	path := fmt.Sprintf("%s/count.json", customCollectionsBasePath)
	return s.client.Count(path, options)
}

// Get individual custom collection
func (s *CustomCollectionServiceOp) Get(collectionID int, options interface{}) (*CustomCollection, error) {
	path := fmt.Sprintf("%s/%d.json", customCollectionsBasePath, collectionID)
	resource := new(CustomCollectionResource)
	err := s.client.Get(path, resource, options)
	return resource.Collection, err
}

// Create a new custom collection
// See Image for the details of the Image creation for a collection.
func (s *CustomCollectionServiceOp) Create(collection CustomCollection) (*CustomCollection, error) {
	path := fmt.Sprintf("%s.json", customCollectionsBasePath)
	wrappedData := CustomCollectionResource{Collection: &collection}
	resource := new(CustomCollectionResource)
	err := s.client.Post(path, wrappedData, resource)
	return resource.Collection, err
}

// Update an existing custom collection
func (s *CustomCollectionServiceOp) Update(collection CustomCollection) (*CustomCollection, error) {
	path := fmt.Sprintf("%s/%d.json", customCollectionsBasePath, collection.ID)
	wrappedData := CustomCollectionResource{Collection: &collection}
	resource := new(CustomCollectionResource)
	err := s.client.Put(path, wrappedData, resource)
	return resource.Collection, err
}

// Delete an existing custom collection.
func (s *CustomCollectionServiceOp) Delete(collectionID int) error {
	return s.client.Delete(fmt.Sprintf("%s/%d.json", customCollectionsBasePath, collectionID))
}
