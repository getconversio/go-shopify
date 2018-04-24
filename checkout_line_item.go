package goshopify

// CheckoutLineItem represents a Shopify checkout line item
type CheckoutLineItem struct {
	ID               string      `json:"id"`
	Key              string      `json:"key"`
	ProductID        int         `json:"product_id"`
	VariantID        int         `json:"variant_id"`
	Sku              string      `json:"sku"`
	Vendor           string      `json:"vendor"`
	Title            string      `json:"title"`
	VariantTitle     string      `json:"variant_title"`
	ImageURL         string      `json:"image_url"`
	Taxable          bool        `json:"taxable"`
	RequiresShipping bool        `json:"requires_shipping"`
	Price            string      `json:"price"`
	CompareAtPrice   interface{} `json:"compare_at_price"`
	LinePrice        string      `json:"line_price"`
	Properties       struct {
	} `json:"properties"`
	Quantity           int           `json:"quantity"`
	Grams              int           `json:"grams"`
	FulfillmentService string        `json:"fulfillment_service"`
	AppliedDiscounts   []interface{} `json:"applied_discounts"`
}
