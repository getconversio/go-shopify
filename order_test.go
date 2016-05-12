package goshopify

import (
	"reflect"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/shopspring/decimal"
)

func TestOrderList(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/admin/orders.json",
		httpmock.NewStringResponder(200, `{"orders": [{"id":1, "total_price": "9.99"},{"id":2}]}`))

	orders, err := client.Order.List(nil)
	if err != nil {
		t.Errorf("Order.List returned error: %v", err)
	}

	expected := []Order{{ID: 1, Total: decimal.NewFromFloat(9.99)}, {ID: 2}}
	if !reflect.DeepEqual(orders, expected) {
		t.Errorf("Order.List returned %+v, expected %+v", orders, expected)
	}

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/admin/orders.json?since_id=1",
		httpmock.NewStringResponder(200, `{"orders": [{"id":2}]}`))

	orders, err = client.Order.List(ListOptions{SinceID: 1})
	if err != nil {
		t.Errorf("Order.List returned error: %v", err)
	}

	expected = []Order{{ID: 2}}
	if !reflect.DeepEqual(orders, expected) {
		t.Errorf("Order.List returned %+v, expected %+v", orders, expected)
	}
}

func TestOrderCount(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/admin/orders/count.json",
		httpmock.NewStringResponder(200, `{"count": 7}`))

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/admin/orders/count.json?created_at_min=2016-01-01T00%3A00%3A00Z",
		httpmock.NewStringResponder(200, `{"count": 2}`))

	cnt, err := client.Order.Count(nil)
	if err != nil {
		t.Errorf("Order.Count returned error: %v", err)
	}

	expected := 7
	if cnt != expected {
		t.Errorf("Order.Count returned %d, expected %d", cnt, expected)
	}

	date := time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC)
	cnt, err = client.Order.Count(CountOptions{CreatedAtMin: date})
	if err != nil {
		t.Errorf("Order.Count returned error: %v", err)
	}

	expected = 2
	if cnt != expected {
		t.Errorf("Order.Count returned %d, expected %d", cnt, expected)
	}
}
