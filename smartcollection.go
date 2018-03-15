package goshopify

import (
	"fmt"
	"time"
)

const smartCollectionsBasePath = "admin/smart_collections"

// SmartCollectionService is an interface for interacting with the smart
// collection endpoints of the Shopify API.
// See https://help.shopify.com/api/reference/smartcollection
type SmartCollectionService interface {
	List(interface{}) ([]SmartCollection, error)
	Count(interface{}) (int, error)
	Get(int, interface{}) (*SmartCollection, error)
	Create(SmartCollection) (*SmartCollection, error)
	Update(SmartCollection) (*SmartCollection, error)
	Delete(int) error
}

// SmartCollectionServiceOp handles communication with the smart collection
// related methods of the Shopify API.
type SmartCollectionServiceOp struct {
	client *Client
}

type Rule struct {
	Column    string `json:"column"`
	Relation  string `json:"relation"`
	Condition string `json:"condition"`
}

// SmartCollection represents a Shopify smart collection.
type SmartCollection struct {
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
	Rules          []Rule     `json:"rules"`
	Disjunctive    bool       `json:"disjunctive"`
}

// SmartCollectionResource represents the result form the smart_collections/X.json endpoint
type SmartCollectionResource struct {
	Collection *SmartCollection `json:"smart_collection"`
}

// SmartCollectionsResource represents the result from the smart_collections.json endpoint
type SmartCollectionsResource struct {
	Collections []SmartCollection `json:"smart_collections"`
}

// List smart collections
func (s *SmartCollectionServiceOp) List(options interface{}) ([]SmartCollection, error) {
	path := fmt.Sprintf("%s.json", smartCollectionsBasePath)
	resource := new(SmartCollectionsResource)
	err := s.client.Get(path, resource, options)
	return resource.Collections, err
}

// Count smart collections
func (s *SmartCollectionServiceOp) Count(options interface{}) (int, error) {
	path := fmt.Sprintf("%s/count.json", smartCollectionsBasePath)
	return s.client.Count(path, options)
}

// Get individual smart collection
func (s *SmartCollectionServiceOp) Get(collectionID int, options interface{}) (*SmartCollection, error) {
	path := fmt.Sprintf("%s/%d.json", smartCollectionsBasePath, collectionID)
	resource := new(SmartCollectionResource)
	err := s.client.Get(path, resource, options)
	return resource.Collection, err
}

// Create a new smart collection
// See Image for the details of the Image creation for a collection.
func (s *SmartCollectionServiceOp) Create(collection SmartCollection) (*SmartCollection, error) {
	path := fmt.Sprintf("%s.json", smartCollectionsBasePath)
	wrappedData := SmartCollectionResource{Collection: &collection}
	resource := new(SmartCollectionResource)
	err := s.client.Post(path, wrappedData, resource)
	return resource.Collection, err
}

// Update an existing smart collection
func (s *SmartCollectionServiceOp) Update(collection SmartCollection) (*SmartCollection, error) {
	path := fmt.Sprintf("%s/%d.json", smartCollectionsBasePath, collection.ID)
	wrappedData := SmartCollectionResource{Collection: &collection}
	resource := new(SmartCollectionResource)
	err := s.client.Put(path, wrappedData, resource)
	return resource.Collection, err
}

// Delete an existing smart collection.
func (s *SmartCollectionServiceOp) Delete(collectionID int) error {
	return s.client.Delete(fmt.Sprintf("%s/%d.json", smartCollectionsBasePath, collectionID))
}
