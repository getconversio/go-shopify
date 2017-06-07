package goshopify

import "time"

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
