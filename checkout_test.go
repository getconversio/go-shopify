package goshopify

import (
	"reflect"
	"testing"

	httpmock "gopkg.in/jarcoal/httpmock.v1"
)

func checkoutTests(t *testing.T, checkout Checkout) {
	// Check that ID is assigned to the returned checkout
	expectedToken := "exuw7apwoycchjuwtiqg8nytfhphr62a"
	if checkout.Token != expectedToken {
		t.Errorf("Checkout.Token returned %+v, expected %+v", checkout.Token, expectedToken)
	}
}

func TestCheckoutGet(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/admin/checkouts/exuw7apwoycchjuwtiqg8nytfhphr62a.json",
		httpmock.NewStringResponder(200, `{"checkout": {"token": "exuw7apwoycchjuwtiqg8nytfhphr62a"}}`))

	checkout, err := client.Checkout.Get("exuw7apwoycchjuwtiqg8nytfhphr62a", nil)
	if err != nil {
		t.Errorf("Checkout.Get returned error: %v", err)
	}

	expected := &Checkout{Token: "exuw7apwoycchjuwtiqg8nytfhphr62a"}
	if !reflect.DeepEqual(checkout, expected) {
		t.Errorf("Checkout.Get returned %+v, expected %+v", checkout, expected)
	}
}

func TestCheckoutCreate(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", "https://fooshop.myshopify.com/admin/checkouts.json",
		httpmock.NewBytesResponder(200, loadFixture("checkout.json")))

	address1 := &Address{ID: 1, FirstName: "Test", LastName: "Citizen", Company: "",
		Address1: "1 Smith St", Address2: "", City: "BRISBANE", Province: "Queensland",
		Zip: "4000", Phone: "1111 111 111", Name: "Test Citizen", ProvinceCode: "QLD", CountryCode: "AU",
		Country: "Australia",
	}
	checkout := Checkout{
		Email:           "test@example.com",
		LineItems:       []CheckoutLineItem{{ID: "c9bac2aa7d06dfa9"}, {ID: "c9bac2aa7d06dfa9"}},
		ShippingAddress: address1,
	}

	returnedCheckout, err := client.Checkout.Create(checkout)
	if err != nil {
		t.Errorf("Checkout.Create returned error: %v", err)
	}

	checkoutTests(t, *returnedCheckout)
}
