package goshopify

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

const recurringApplicationChargesBasePath = "admin/recurring_application_charges"

// RecurringApplicationChargeService is an interface for interfacing with the recurringApplicationCharges endpoints
// of the Shopify API.
// See: https://help.shopify.com/api/reference/recurringapplicationcharge
type RecurringApplicationChargeService interface {
	List(interface{}) ([]RecurringApplicationCharge, error)
	Get(int, interface{}) (*RecurringApplicationCharge, error)
	Create(RecurringApplicationCharge) (*RecurringApplicationCharge, error)
	Activate(RecurringApplicationCharge) (*RecurringApplicationCharge, error)
	Delete(int) error
}

// RecurringApplicationChargeServiceOp handles communication with the product related methods of
// the Shopify API.
type RecurringApplicationChargeServiceOp struct {
	client *Client
}

// RecurringApplicationCharge represents a Shopify recurringApplicationCharge
type RecurringApplicationCharge struct {
	ID                  int              `json:"id"`
	ActivatedOn         *time.Time       `json:"activated_on"`
	BillingOn         	*time.Time       `json:"billing_on"`
	CancelledOn         *time.Time       `json:"cancelled_on"`
	CappedAmount        *decimal.Decimal `json:"capped_amount"`
	ConfirmationUrl     string           `json:"confirmation_url"`
	Name     			string           `json:"name"`
	Price		        *decimal.Decimal `json:"price"`
	ReturnUrl		    string           `json:"return_url"`
	Status			    string           `json:"status"`
	Terms			    string           `json:"terms"`
	Test	            bool             `json:"test"`
	TrialDays           int              `json:"trial_days"`
	TrialEndsOn         *time.Time       `json:"trial_ends_on"`
	CreatedAt           *time.Time       `json:"created_at"`
	UpdatedAt           *time.Time       `json:"updated_at"`
}

// Represents the result from the recurring_application_charges/X.json endpoint
type RecurringApplicationChargeResource struct {
	RecurringApplicationCharge *RecurringApplicationCharge `json:"recurring_application_charge"`
}

// Represents the result from the recurring_application_charges.json endpoint
type RecurringApplicationChargesResource struct {
	RecurringApplicationCharges []RecurringApplicationCharge `json:"recurring_application_charges"`
}

// Represents the options available when searching for a RecurringApplicationCharge
type RecurringApplicationChargeSearchOptions struct {
	SinceId   	int    `url:"since_id,omitempty"`
	Fields 		string `url:"fields,omitempty"`
}

// List recurringApplicationCharges
func (s *RecurringApplicationChargeServiceOp) List(options interface{}) ([]RecurringApplicationCharge, error) {
	path := fmt.Sprintf("%s.json", recurringApplicationChargesBasePath)
	resource := new(RecurringApplicationChargesResource)
	err := s.client.Get(path, resource, options)
	return resource.RecurringApplicationCharges, err
}

// Get recurringApplicationCharge
func (s *RecurringApplicationChargeServiceOp) Get(recurringApplicationChargeID int, options interface{}) (*RecurringApplicationCharge, error) {
	path := fmt.Sprintf("%s/%v.json", recurringApplicationChargesBasePath, recurringApplicationChargeID)
	resource := new(RecurringApplicationChargeResource)
	err := s.client.Get(path, resource, options)
	return resource.RecurringApplicationCharge, err
}

// Create a new RecurringApplicationCharge
func (s *RecurringApplicationChargeServiceOp) Create(recurringApplicationCharge RecurringApplicationCharge) (*RecurringApplicationCharge, error) {
	path := fmt.Sprintf("%s.json", recurringApplicationChargesBasePath)
	wrappedData := RecurringApplicationChargeResource{RecurringApplicationCharge: &recurringApplicationCharge}
	resource := new(RecurringApplicationChargeResource)
	err := s.client.Post(path, wrappedData, resource)
	return resource.RecurringApplicationCharge, err
}

// Activate an existing RecurringApplicationCharge
func (s *RecurringApplicationChargeServiceOp) Activate(recurringApplicationCharge RecurringApplicationCharge) (*RecurringApplicationCharge, error) {
	path := fmt.Sprintf("%s/%d/activate.json", recurringApplicationChargesBasePath, recurringApplicationCharge.ID)
	wrappedData := RecurringApplicationChargeResource{RecurringApplicationCharge: &recurringApplicationCharge}
	resource := new(RecurringApplicationChargeResource)
	err := s.client.Post(path, wrappedData, resource)
	return resource.RecurringApplicationCharge, err
}

// Delete an existing RecurringApplicationCharge
func (s *RecurringApplicationChargeServiceOp) Delete(recurringApplicationChargeID int) error {
	return s.client.Delete(fmt.Sprintf("%s/%d.json", recurringApplicationChargesBasePath, recurringApplicationChargeID))
}