package goshopify

// Variant represents a Shopify variant
type Variant struct {
	ID                   int    `json:"id"`
	ProductID            int    `json:"product_id"`
	Title                string `json:"title"`
	Sku                  string `json:"sku"`
	Position             int    `json:"position"`
	Grams                int    `json:"grams"`
	InventoryPolicy      string `json:"inventory_policy"`
	CompareAtPrice       string `json:"compare_at_price"`
	FulfillmentService   string `json:"fulfillment_service"`
	InventoryManagement  string `json:"inventory_management"`
	Option1              string `json:"option1"`
	Option2              string `json:"option2"`
	Option3              string `json:"option3"`
	CreatedAt            string `json:"created_at"`
	UpdatedAt            string `json:"updated_at"`
	Taxable              bool   `json:"taxable"`
	Barcode              string `json:"barcode"`
	ImageID              int    `json:"image_id"`
	InventoryQuantity    int    `json:"inventory_quantity"`
	Weight               int    `json:"weight"`
	WeightUnit           string `json:"weight_unit"`
	OldInventoryQuantity int    `json:"old_inventory_quantity"`
	RequireShipping      bool   `json:"requires_shipping"`
}
