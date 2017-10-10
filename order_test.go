package goshopify

import (
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"gopkg.in/jarcoal/httpmock.v1"
)

func orderTests(t *testing.T, order Order) {
	// Check that dates are parsed
	d := time.Date(2016, time.May, 17, 4, 14, 36, 0, time.UTC)
	if !d.Equal(*order.CreatedAt) {
		t.Errorf("Order.CreatedAt returned %+v, expected %+v", order.CreatedAt, d)
	}

	// Check null dates
	if order.ProcessedAt != nil {
		t.Errorf("Order.ProcessedAt returned %+v, expected %+v", order.ProcessedAt, nil)
	}

	// Check prices
	p := decimal.NewFromFloat(10)
	if !p.Equals(*order.TotalPrice) {
		t.Errorf("Order.TotalPrice returned %+v, expected %+v", order.TotalPrice, p)
	}

	// Check null prices, notice that prices are usually not empty.
	if order.TotalTax != nil {
		t.Errorf("Order.TotalTax returned %+v, expected %+v", order.TotalTax, nil)
	}

	// Check customer
	if order.Customer == nil {
		t.Error("Expected Customer to not be nil")
	}
	if order.Customer.Email != "john@test.com" {
		t.Errorf("Customer.Email, expected %v, actual %v", "john@test.com", order.Customer.Email)
	}
}

func transactionTest(t *testing.T, transaction Transaction) {
	// Check that dates are parsed
	d := time.Date(2017, time.October, 9, 19, 26, 23, 0, time.UTC)
	if !d.Equal(*transaction.CreatedAt) {
		t.Errorf("Transaction.CreatedAt returned %+v, expected %+v", transaction.CreatedAt, d)
	}

	// Check null value
	if transaction.LocationID != nil {
		t.Error("Expected Transaction.LocationID to be nil")
	}

	if transaction.PaymentDetails == nil {
		t.Error("Expected Transaction.PaymentDetails to not be nil")
	}
}

func TestOrderList(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/admin/orders.json",
		httpmock.NewBytesResponder(200, loadFixture("orders.json")))

	orders, err := client.Order.List(nil)
	if err != nil {
		t.Errorf("Order.List returned error: %v", err)
	}

	// Check that orders were parsed
	if len(orders) != 1 {
		t.Errorf("Order.List got %v orders, expected: 1", len(orders))
	}

	order := orders[0]
	orderTests(t, order)
}

func TestOrderListOptions(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/admin/orders.json?limit=250&page=10&status=any",
		httpmock.NewBytesResponder(200, loadFixture("orders.json")))

	options := OrderListOptions{
		Page:   10,
		Limit:  250,
		Status: "any"}

	orders, err := client.Order.List(options)
	if err != nil {
		t.Errorf("Order.List returned error: %v", err)
	}

	// Check that orders were parsed
	if len(orders) != 1 {
		t.Errorf("Order.List got %v orders, expected: 1", len(orders))
	}

	order := orders[0]
	orderTests(t, order)
}

func TestOrderGet(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/admin/orders/123456.json",
		httpmock.NewBytesResponder(200, loadFixture("order.json")))

	order, err := client.Order.Get(123456, nil)
	if err != nil {
		t.Errorf("Order.List returned error: %v", err)
	}

	orderTests(t, *order)
}

func TestOrderGetWithTransactions(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/admin/orders/123456.json",
		httpmock.NewBytesResponder(200, loadFixture("order_with_transaction.json")))

	options := struct {
		ApiFeatures string `url:"_apiFeatures"`
	}{"include-transactions"}

	order, err := client.Order.Get(123456, options)
	if err != nil {
		t.Errorf("Order.List returned error: %v", err)
	}

	orderTests(t, *order)

	// Check transactions is not nil
	if order.Transactions == nil {
		t.Error("Expected Transactions to not be nil")
	}
	if len(order.Transactions) != 1 {
		t.Errorf("Expected Transactions to have 1 transaction but received %v", len(order.Transactions))
	}

	transactionTest(t, order.Transactions[0])
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
