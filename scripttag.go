package goshopify

import (
	"fmt"
	"time"
)

const scripttagsBasePath = "admin/script_tags"

// ScriptTagService is an interface for interfacing with the scripttag endpoints
// of the Shopify API.
// See: https://help.shopify.com/api/reference/scripttag
type ScriptTagService interface {
	List(interface{}) ([]ScriptTag, error)
	Count(interface{}) (int, error)
	Get(int, interface{}) (*ScriptTag, error)
	Create(ScriptTag) (*ScriptTag, error)
	Update(ScriptTag) (*ScriptTag, error)
	Delete(int) error
}

// ScriptTagServiceOp handles communication with the scripttag related methods of
// the Shopify API.
type ScriptTagServiceOp struct {
	client *Client
}

// ScriptTag represents a Shopify scripttag
type ScriptTag struct {
	ID                             int             `json:"id"`
	Event                          string          `json:"event"`
	Src                       	   string          `json:"src"`
	DisplayScope                   string          `json:"display_scope"`
	CreatedAt                      *time.Time      `json:"created_at"`
	UpdatedAt                      *time.Time      `json:"updated_at"`
}

// The options provided by Shopify
type ScriptTagOption struct {
	Limit        	int      	`json:"limit"`
	Page 		 	int      	`json:"page"`
	SinceId      	int   	 	`json:"since_id"`
	CreatedAtMin  	string      `json:"created_at_min"`
	CreatedAtMax  	string      `json:"created_at_max"`
	UpdatedAtMin  	string      `json:"updated_at_min"`
	UpdatedAtMax  	string      `json:"updated_at_max"`
	Fields    		[]string 	`json:"fields"`
	Src				string		`json:"src"`
}

// Represents the result from the script_tags/X.json endpoint
type ScriptTagResource struct {
	ScriptTag *ScriptTag `json:"script_tag"`
}

// Represents the result from the script_tags.json endpoint
type ScriptTagsResource struct {
	ScriptTags []ScriptTag `json:"script_tags"`
}

// List scripttags
func (s *ScriptTagServiceOp) List(options interface{}) ([]ScriptTag, error) {
	path := fmt.Sprintf("%s.json", scripttagsBasePath)
	resource := new(ScriptTagsResource)
	err := s.client.Get(path, resource, options)
	return resource.ScriptTags, err
}

// Count scripttags
func (s *ScriptTagServiceOp) Count(options interface{}) (int, error) {
	path := fmt.Sprintf("%s/count.json", scripttagsBasePath)
	return s.client.Count(path, options)
}

// Get individual scripttag
func (s *ScriptTagServiceOp) Get(scripttagID int, options interface{}) (*ScriptTag, error) {
	path := fmt.Sprintf("%s/%d.json", scripttagsBasePath, scripttagID)
	resource := new(ScriptTagResource)
	err := s.client.Get(path, resource, options)
	return resource.ScriptTag, err
}

// Create a new scripttag
func (s *ScriptTagServiceOp) Create(scripttag ScriptTag) (*ScriptTag, error) {
	path := fmt.Sprintf("%s.json", scripttagsBasePath)
	wrappedData := ScriptTagResource{ScriptTag: &scripttag}
	resource := new(ScriptTagResource)
	err := s.client.Post(path, wrappedData, resource)
	return resource.ScriptTag, err
}

// Update an existing scripttag
func (s *ScriptTagServiceOp) Update(scripttag ScriptTag) (*ScriptTag, error) {
	path := fmt.Sprintf("%s/%d.json", scripttagsBasePath, scripttag.ID)
	wrappedData := ScriptTagResource{ScriptTag: &scripttag}
	resource := new(ScriptTagResource)
	err := s.client.Put(path, wrappedData, resource)
	return resource.ScriptTag, err
}

// Delete an existing scripttag
func (s *ScriptTagServiceOp) Delete(scripttagID int) error {
	return s.client.Delete(fmt.Sprintf("%s/%d.json", scripttagsBasePath, scripttagID))
}
