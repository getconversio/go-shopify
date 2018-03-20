package goshopify

import (
	"reflect"
	"testing"

	httpmock "gopkg.in/jarcoal/httpmock.v1"
)

func recurringApplicationChargeTests(t *testing.T, collection RecurringApplicationCharge) {

	// Test a few fields
	cases := []struct {
		field    string
		expected interface{}
		actual   interface{}
	}{
		{"ID", 44433414, collection.ID},
		{"Name", "App charge", collection.Name},
		{"Status", "pending", collection.Status},
		{"Test", true, collection.Test},
	}

	for _, c := range cases {
		if c.expected != c.actual {
			t.Errorf("RecurringApplicationCharge.%v returned %v, expected %v", c.field, c.actual, c.expected)
		}
	}
}

func TestRecurringApplicationChargeList(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/admin/recurring_application_charges.json",
		httpmock.NewStringResponder(200, `{"recurring_application_charges": [{"id":1},{"id":2}]}`))

	collections, err := client.RecurringApplicationCharge.List(nil)
	if err != nil {
		t.Errorf("RecurringApplicationCharge.List returned error: %v", err)
	}

	expected := []RecurringApplicationCharge{{ID: 1}, {ID: 2}}
	if !reflect.DeepEqual(collections, expected) {
		t.Errorf("RecurringApplicationCharge.List returned %+v, expected %+v", collections, expected)
	}
}

func TestRecurringApplicationChargeGet(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/admin/recurring_application_charges/1.json",
		httpmock.NewStringResponder(200, `{"recurring_application_charge": {"id":1}}`))

	collection, err := client.RecurringApplicationCharge.Get(1, nil)
	if err != nil {
		t.Errorf("RecurringApplicationCharge.Get returned error: %v", err)
	}

	expected := &RecurringApplicationCharge{ID: 1}
	if !reflect.DeepEqual(collection, expected) {
		t.Errorf("RecurringApplicationCharge.Get returned %+v, expected %+v", collection, expected)
	}
}

func TestRecurringApplicationChargeCreate(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", "https://fooshop.myshopify.com/admin/recurring_application_charges.json",
		httpmock.NewBytesResponder(200, loadFixture("recurringapplicationcharge.json")))

	collection := RecurringApplicationCharge{
		Name: "App charge",
	}

	returnedCollection, err := client.RecurringApplicationCharge.Create(collection)
	if err != nil {
		t.Errorf("RecurringApplicationCharge.Create returned error: %v", err)
	}

	recurringApplicationChargeTests(t, *returnedCollection)
}

func TestRecurringApplicationChargeActivate(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", "https://fooshop.myshopify.com/admin/recurring_application_charges/1/activate.json",
		httpmock.NewBytesResponder(200, loadFixture("recurringapplicationcharge.json")))

	collection := RecurringApplicationCharge{
		ID:    1,
		Name: "App charge",
	}

	returnedCollection, err := client.RecurringApplicationCharge.Activate(collection)
	if err != nil {
		t.Errorf("RecurringApplicationCharge.Activate returned error: %v", err)
	}

	recurringApplicationChargeTests(t, *returnedCollection)
}

func TestRecurringApplicationChargeDelete(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("DELETE", "https://fooshop.myshopify.com/admin/recurring_application_charges/1.json",
		httpmock.NewStringResponder(200, "{}"))

	err := client.RecurringApplicationCharge.Delete(1)
	if err != nil {
		t.Errorf("RecurringApplicationCharge.Delete returned error: %v", err)
	}
}
