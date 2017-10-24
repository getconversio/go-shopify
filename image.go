package goshopify

import (
	"fmt"
	"time"
)

// ImageService is an interface for interacting with the image endpoints
// of the Shopify API.
// See https://help.shopify.com/api/reference/product_image
type ImageService interface {
	List(int, interface{}) ([]Image, error)
	Count(int, interface{}) (int, error)
}

// ImageServiceOp handles communication with the image related methods of
// the Shopify API.
type ImageServiceOp struct {
	client *Client
}

// Image represents a Shopify product's image
type Image struct {
	ID         int        `json:"id"`
	ProductID  int        `json:"product_id"`
	Position   int        `json:"position"`
	CreatedAt  *time.Time `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
	Width      int        `json:"width"`
	Height     int        `json:"height"`
	Src        string     `json:"src"`
	VariantIds []int      `json:"variant_ids"`
}

// ImagesResource represents the result from the products/X/images.json endpoint
type ImagesResource struct {
	Images []Image `json:"images"`
}

// List images
func (s *ImageServiceOp) List(productID int, options interface{}) ([]Image, error) {
	path := fmt.Sprintf("%s/%d/images.json", productsBasePath, productID)
	resource := new(ImagesResource)
	err := s.client.Get(path, resource, options)
	return resource.Images, err
}

// Count images
func (s *ImageServiceOp) Count(productID int, options interface{}) (int, error) {
	path := fmt.Sprintf("%s/%d/images/count.json", productsBasePath, productID)
	return s.client.Count(path, options)
}
