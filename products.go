package goshopify

import (
	"fmt"
)

const productsBasePath = "admin/products"

// ProductsService is an interface for interfacing with the product endpoints
// of the Shopify API.
// See: https://help.shopify.com/api/reference/product
type ProductsService interface {
	List() ([]Product, error)
	Count() (int, error)
	Get(int) (*Product, error)
}

// ProductsServiceOp handles communication with the product related methods of
// the Shopify API.
type ProductsServiceOp struct {
	client *Client
}

// Product represents a Shopify product
type Product struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// productRoot represents a product root
type productRoot struct {
	Product *Product `json:"product"`
}

type productsRoot struct {
	Products []Product `json:"products"`
}

type countRoot struct {
	Count int `json:"count"`
}

// Performs a list request given a path
func (s *ProductsServiceOp) List() ([]Product, error) {
	path := fmt.Sprintf("%s.json", productsBasePath)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	root := new(productsRoot)
	err = s.client.Do(req, root)
	if err != nil {
		return nil, err
	}

	return root.Products, err
}

// Count products
func (s *ProductsServiceOp) Count() (int, error) {
	path := fmt.Sprintf("%s/count.json", productsBasePath)

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return 0, err
	}

	root := new(countRoot)
	err = s.client.Do(req, root)
	if err != nil {
		return 0, err
	}

	return root.Count, err
}

// Get individual product
func (s *ProductsServiceOp) Get(productID int) (*Product, error) {
	path := fmt.Sprintf("%s/%d.json", productsBasePath, productID)

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	root := new(productRoot)
	err = s.client.Do(req, root)
	if err != nil {
		return nil, err
	}

	return root.Product, err
}
