package goshopify

import (
	"time"

	"github.com/shopspring/decimal"
)

// VariantService is an interface for interacting with the variant endpoints
// of the Shopify API.
// See https://help.shopify.com/api/reference/product_variant
type VariantService interface {
	List(int, interface{}) ([]Variant, error)
}

// VariantServiceOp handles communication with the variant related methods of
// the Shopify API.
type VariantServiceOp struct {
	client *Client
}

// Variant represents a Shopify variant
type Variant struct {
	ID                   int              `json:"id"`
	ProductID            int              `json:"product_id"`
	Title                string           `json:"title"`
	Sku                  string           `json:"sku"`
	Position             int              `json:"position"`
	Grams                int              `json:"grams"`
	InventoryPolicy      string           `json:"inventory_policy"`
	Price                *decimal.Decimal `json:"price"`
	CompareAtPrice       *decimal.Decimal `json:"compare_at_price"`
	FulfillmentService   string           `json:"fulfillment_service"`
	InventoryManagement  string           `json:"inventory_management"`
	Option1              string           `json:"option1"`
	Option2              string           `json:"option2"`
	Option3              string           `json:"option3"`
	CreatedAt            *time.Time       `json:"created_at"`
	UpdatedAt            *time.Time       `json:"updated_at"`
	Taxable              bool             `json:"taxable"`
	Barcode              string           `json:"barcode"`
	ImageID              int              `json:"image_id"`
	InventoryQuantity    int              `json:"inventory_quantity"`
	Weight               *decimal.Decimal `json:"weight"`
	WeightUnit           string           `json:"weight_unit"`
	OldInventoryQuantity int              `json:"old_inventory_quantity"`
	RequireShipping      bool             `json:"requires_shipping"`
}

// VariantsResource represents the result from the products/X/variants.json endpoint
type VariantsResource struct {
	Variants []Variant `json:"variants"`
}
