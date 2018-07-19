package goshopify

// CustomerAddress represents a Shopify customer address
type CustomerAddress struct {
	ID           int    `json:"id,omitempty"`
	CustomerID   int    `json:"customer_id,omitempty"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Company      string `json:"company"`
	Address1     string `json:"address1"`
	Address2     string `json:"address2"`
	City         string `json:"city"`
	Province     string `json:"province"`
	Country      string `json:"country"`
	Zip          string `json:"zip"`
	Phone        string `json:"phone"`
	Name         string `json:"name"`
	ProvinceCode string `json:"province_code"`
	CountryCode  string `json:"country_code"`
	CountryName  string `json:"country_name"`
	Default      bool   `json:"default"`
}
