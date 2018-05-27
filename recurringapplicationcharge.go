package goshopify

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

const recurringApplicationChargesBasePath = "admin/recurring_application_charges"

// RecurringApplicationChargeService is an interface for interacting with the
// RecurringApplicationCharge endpoints of the Shopify API.
// See https://help.shopify.com/api/reference/billing/recurringapplicationcharge
type RecurringApplicationChargeService interface {
	Create(RecurringApplicationCharge) (*RecurringApplicationCharge, error)
	Get(int, interface{}) (*RecurringApplicationCharge, error)
	List(interface{}) ([]RecurringApplicationCharge, error)
	Activate(RecurringApplicationCharge) (*RecurringApplicationCharge, error)
	Delete(int) error
	Update(int, int) (*RecurringApplicationCharge, error)
}

// RecurringApplicationChargeServiceOp handles communication with the
// RecurringApplicationCharge related methods of the Shopify API.
type RecurringApplicationChargeServiceOp struct {
	client *Client
}

// RecurringApplicationCharge represents a Shopify RecurringApplicationCharge.
type RecurringApplicationCharge struct {
	APIClientID           int              `json:"api_client_id"`
	ActivatedOn           *time.Time       `json:"activated_on"`
	BalanceRemaining      *decimal.Decimal `json:"balance_remaining"`
	BalanceUsed           *decimal.Decimal `json:"balance_used"`
	BillingOn             *time.Time       `json:"billing_on"`
	CancelledOn           *time.Time       `json:"cancelled_on"`
	CappedAmount          *decimal.Decimal `json:"capped_amount"`
	ConfirmationURL       string           `json:"confirmation_url"`
	CreatedAt             *time.Time       `json:"created_at"`
	DecoratedReturnURL    string           `json:"decorated_return_url"`
	ID                    int              `json:"id"`
	Name                  string           `json:"name"`
	Price                 *decimal.Decimal `json:"price"`
	ReturnURL             string           `json:"return_url"`
	RiskLevel             *decimal.Decimal `json:"risk_level"`
	Status                string           `json:"status"`
	Terms                 string           `json:"terms"`
	Test                  *bool            `json:"test"`
	TrialDays             int              `json:"trial_days"`
	TrialEndsOn           *time.Time       `json:"trial_ends_on"`
	UpdateCappedAmountURL string           `json:"update_capped_amount_url"`
	UpdatedAt             *time.Time       `json:"updated_at"`
}

// RecurringApplicationChargeResource represents the result from the
// admin/recurring_application_charges{/X{/activate.json}.json}.json endpoints.
type RecurringApplicationChargeResource struct {
	Charge *RecurringApplicationCharge `json:"recurring_application_charge"`
}

// RecurringApplicationChargesResource represents the result from the
// admin/recurring_application_charges.json endpoint.
type RecurringApplicationChargesResource struct {
	Charges []RecurringApplicationCharge `json:"recurring_application_charges"`
}

// Create creates new recurring application charge.
func (r *RecurringApplicationChargeServiceOp) Create(charge RecurringApplicationCharge) (
	*RecurringApplicationCharge, error) {

	path := fmt.Sprintf("%s.json", recurringApplicationChargesBasePath)
	wrappedData := RecurringApplicationChargeResource{Charge: &charge}
	resource := &RecurringApplicationChargeResource{}
	err := r.client.Post(path, wrappedData, resource)
	return resource.Charge, err
}

// Get gets individual recurring application charge.
func (r *RecurringApplicationChargeServiceOp) Get(chargeID int, options interface{}) (
	*RecurringApplicationCharge, error) {

	path := fmt.Sprintf("%s/%d.json", recurringApplicationChargesBasePath, chargeID)
	resource := &RecurringApplicationChargeResource{}
	err := r.client.Get(path, resource, options)
	return resource.Charge, err
}

// List gets all recurring application charges.
func (r *RecurringApplicationChargeServiceOp) List(options interface{}) (
	[]RecurringApplicationCharge, error) {

	path := fmt.Sprintf("%s.json", recurringApplicationChargesBasePath)
	resource := &RecurringApplicationChargesResource{}
	err := r.client.Get(path, resource, options)
	return resource.Charges, err
}

// Activate activates recurring application charge.
func (r *RecurringApplicationChargeServiceOp) Activate(charge RecurringApplicationCharge) (
	*RecurringApplicationCharge, error) {

	path := fmt.Sprintf("%s/%d/activate.json", recurringApplicationChargesBasePath, charge.ID)
	wrappedData := RecurringApplicationChargeResource{Charge: &charge}
	resource := &RecurringApplicationChargeResource{}
	err := r.client.Post(path, wrappedData, resource)
	return resource.Charge, err
}

// Delete deletes recurring application charge.
func (r *RecurringApplicationChargeServiceOp) Delete(chargeID int) error {
	return r.client.Delete(fmt.Sprintf("%s/%d.json", recurringApplicationChargesBasePath, chargeID))
}

// Update updates recurring application charge.
func (r *RecurringApplicationChargeServiceOp) Update(chargeID, newCappedAmount int) (
	*RecurringApplicationCharge, error) {

	path := fmt.Sprintf("%s/%d/customize.json?recurring_application_charge[capped_amount]=%d",
		recurringApplicationChargesBasePath, chargeID, newCappedAmount)
	resource := &RecurringApplicationChargeResource{}
	err := r.client.Put(path, nil, resource)
	return resource.Charge, err
}
