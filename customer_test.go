package goshopify

import (
	"reflect"
	"testing"
	"time"

	"gopkg.in/jarcoal/httpmock.v1"
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

func TestCustomerGet(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/admin/customers/1.json",
		httpmock.NewStringResponder(200, `{"customer": {"id":1}}`))

	customer, err := client.Customer.Get(1, nil)
	if err != nil {
		t.Errorf("Customer.Get returned error: %v", err)
	}

	expected := &Customer{ID: 1}
	if !reflect.DeepEqual(customer, expected) {
		t.Errorf("Customer.Get returned %+v, expected %+v", customer, expected)
	}
}
