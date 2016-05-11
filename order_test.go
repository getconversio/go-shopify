package goshopify

import (
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestOrderCount(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/admin/orders/count.json",
		httpmock.NewStringResponder(200, `{"count": 7}`))

	cnt, err := client.Order.Count()
	if err != nil {
		t.Errorf("Order.Count returned error: %v", err)
	}

	expected := 7
	if cnt != expected {
		t.Errorf("Order.Count returned %d, expected %d", cnt, expected)
	}
}
