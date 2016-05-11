package goshopify

import (
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestCustomerCount(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/admin/customers/count.json",
		httpmock.NewStringResponder(200, `{"count": 5}`))

	cnt, err := client.Customer.Count()
	if err != nil {
		t.Errorf("Customer.Count returned error: %v", err)
	}

	expected := 5
	if cnt != expected {
		t.Errorf("Customer.Count returned %d, expected %d", cnt, expected)
	}
}
