package goshopify

import (
	"reflect"
	"testing"
	"time"

	"github.com/shopspring/decimal"
	httpmock "gopkg.in/jarcoal/httpmock.v1"
)

func TestCustomerList(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/admin/customers.json",
		httpmock.NewStringResponder(200, `{"customers": [{"id":1},{"id":2}]}`))

	customers, err := client.Customer.List(nil)
	if err != nil {
		t.Errorf("Customer.List returned error: %v", err)
	}

	expected := []Customer{{ID: 1}, {ID: 2}}
	if !reflect.DeepEqual(customers, expected) {
		t.Errorf("Customer.List returned %+v, expected %+v", customers, expected)
	}
}

func TestCustomerCount(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/admin/customers/count.json",
		httpmock.NewStringResponder(200, `{"count": 5}`))

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/admin/customers/count.json?created_at_min=2016-01-01T00%3A00%3A00Z",
		httpmock.NewStringResponder(200, `{"count": 2}`))

	cnt, err := client.Customer.Count(nil)
	if err != nil {
		t.Errorf("Customer.Count returned error: %v", err)
	}

	expected := 5
	if cnt != expected {
		t.Errorf("Customer.Count returned %d, expected %d", cnt, expected)
	}

	date := time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC)
	cnt, err = client.Customer.Count(CountOptions{CreatedAtMin: date})
	if err != nil {
		t.Errorf("Customer.Count returned error: %v", err)
	}

	expected = 2
	if cnt != expected {
		t.Errorf("Customer.Count returned %d, expected %d", cnt, expected)
	}
}

func TestCustomerSearch(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/admin/customers/search.json",
		httpmock.NewStringResponder(200, `{"customers": [{"id":1},{"id":2}]}`))

	customers, err := client.Customer.Search(nil)
	if err != nil {
		t.Errorf("Customer.Search returned error: %v", err)
	}

	expected := []Customer{{ID: 1}, {ID: 2}}
	if !reflect.DeepEqual(customers, expected) {
		t.Errorf("Customer.Search returned %+v, expected %+v", customers, expected)
	}
}

func TestCustomerGet(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/admin/customers/1.json",
		httpmock.NewBytesResponder(200, loadFixture("customer.json")))

	customer, err := client.Customer.Get(1, nil)
	if err != nil {
		t.Errorf("Customer.Get returned error: %v", err)
	}

	loc := time.FixedZone("AEST", 10)

	createdAt := time.Date(2017, time.September, 23, 18, 15, 47, 0, loc)
	updatedAt := time.Date(2017, time.September, 23, 18, 15, 47, 0, loc)
	totalSpent := decimal.NewFromFloat(278.60)

	expectation := &Customer{
		ID:               1,
		Email:            "test@example.com",
		FirstName:        "Test",
		LastName:         "Citizen",
		AcceptsMarketing: true,
		VerifiedEmail:    true,
		TaxExempt:        false,
		OrdersCount:      4,
		State:            "enabled",
		TotalSpent:       &totalSpent,
		LastOrderId:      123,
		Note:             "",
		Phone:            "",
		CreatedAt:        &createdAt,
		UpdatedAt:        &updatedAt,
	}

	if customer.ID != expectation.ID {
		t.Errorf("Customer.ID returned %+v, expected %+v", customer.ID, expectation.ID)
	}
	if customer.Email != expectation.Email {
		t.Errorf("Customer.Email returned %+v, expected %+v", customer.ID, expectation.Email)
	}
	if customer.FirstName != expectation.FirstName {
		t.Errorf("Customer.FirstName returned %+v, expected %+v", customer.FirstName, expectation.FirstName)
	}
	if customer.LastName != expectation.LastName {
		t.Errorf("Customer.LastName returned %+v, expected %+v", customer.LastName, expectation.LastName)
	}
	if customer.AcceptsMarketing != expectation.AcceptsMarketing {
		t.Errorf("Customer.AcceptsMarketing returned %+v, expected %+v", customer.AcceptsMarketing, expectation.AcceptsMarketing)
	}
	if customer.CreatedAt.Equal(*expectation.CreatedAt) {
		t.Errorf("Customer.CreatedAt returned %+v, expected %+v", customer.CreatedAt, expectation.CreatedAt)
	}
	if customer.UpdatedAt.Equal(*expectation.UpdatedAt) {
		t.Errorf("Customer.UpdatedAt returned %+v, expected %+v", customer.UpdatedAt, expectation.UpdatedAt)
	}
	if customer.OrdersCount != expectation.OrdersCount {
		t.Errorf("Customer.OrdersCount returned %+v, expected %+v", customer.OrdersCount, expectation.OrdersCount)
	}
	if customer.State != expectation.State {
		t.Errorf("Customer.State returned %+v, expected %+v", customer.State, expectation.State)
	}
	if !expectation.TotalSpent.Truncate(2).Equals(customer.TotalSpent.Truncate(2)) {
		t.Errorf("Customer.TotalSpent returned %+v, expected %+v", customer.TotalSpent, expectation.TotalSpent)
	}
	if customer.LastOrderId != expectation.LastOrderId {
		t.Errorf("Customer.LastOrderId returned %+v, expected %+v", customer.LastOrderId, expectation.LastOrderId)
	}
	if customer.Note != expectation.Note {
		t.Errorf("Customer.Note returned %+v, expected %+v", customer.Note, expectation.Note)
	}
	if customer.VerifiedEmail != expectation.VerifiedEmail {
		t.Errorf("Customer.Note returned %+v, expected %+v", customer.VerifiedEmail, expectation.VerifiedEmail)
	}
	if customer.TaxExempt != expectation.TaxExempt {
		t.Errorf("Customer.TaxExempt returned %+v, expected %+v", customer.TaxExempt, expectation.TaxExempt)
	}
	if customer.Phone != expectation.Phone {
		t.Errorf("Customer.Phone returned %+v, expected %+v", customer.Phone, expectation.Phone)
	}
}

func TestCustomerUpdate(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("PUT", "https://fooshop.myshopify.com/admin/customers/1.json",
		httpmock.NewBytesResponder(200, loadFixture("customer.json")))

	customer := Customer{
		ID:   1,
		Tags: "new",
	}

	returnedCustomer, err := client.Customer.Update(customer)
	if err != nil {
		t.Errorf("Customer.Update returned error: %v", err)
	}

	expectedCustomerID := 1
	if returnedCustomer.ID != expectedCustomerID {
		t.Errorf("Customer.ID returned %+v expected %+v", returnedCustomer.ID, expectedCustomerID)
	}
}
