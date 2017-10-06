package goshopify

import (
	"reflect"
	"testing"
	"time"

	httpmock "gopkg.in/jarcoal/httpmock.v1"
)

func productTests(t *testing.T, product Product) {
	// Check that ID is assigned to the returned product
	expectedInt := 1071559748
	if product.ID != expectedInt {
		t.Errorf("Product.ID returned %+v, expected %+v", product.ID, expectedInt)
	}
}

func TestProductList(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/admin/products.json",
		httpmock.NewStringResponder(200, `{"products": [{"id":1},{"id":2}]}`))

	products, err := client.Product.List(nil)
	if err != nil {
		t.Errorf("Product.List returned error: %v", err)
	}

	expected := []Product{{ID: 1}, {ID: 2}}
	if !reflect.DeepEqual(products, expected) {
		t.Errorf("Product.List returned %+v, expected %+v", products, expected)
	}
}

func TestProductCount(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/admin/products/count.json",
		httpmock.NewStringResponder(200, `{"count": 3}`))

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/admin/products/count.json?created_at_min=2016-01-01T00%3A00%3A00Z",
		httpmock.NewStringResponder(200, `{"count": 2}`))

	cnt, err := client.Product.Count(nil)
	if err != nil {
		t.Errorf("Product.Count returned error: %v", err)
	}

	expected := 3
	if cnt != expected {
		t.Errorf("Product.Count returned %d, expected %d", cnt, expected)
	}

	date := time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC)
	cnt, err = client.Product.Count(CountOptions{CreatedAtMin: date})
	if err != nil {
		t.Errorf("Product.Count returned error: %v", err)
	}

	expected = 2
	if cnt != expected {
		t.Errorf("Product.Count returned %d, expected %d", cnt, expected)
	}
}

func TestProductGet(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/admin/products/1.json",
		httpmock.NewStringResponder(200, `{"product": {"id":1}}`))

	product, err := client.Product.Get(1, nil)
	if err != nil {
		t.Errorf("Product.Get returned error: %v", err)
	}

	expected := &Product{ID: 1}
	if !reflect.DeepEqual(product, expected) {
		t.Errorf("Product.Get returned %+v, expected %+v", product, expected)
	}
}

func TestProductCreate(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", "https://fooshop.myshopify.com/admin/products.json",
		httpmock.NewBytesResponder(200, loadFixture("product.json")))

	product := Product{
		Title:       "Burton Custom Freestyle 151",
		BodyHTML:    "<strong>Good snowboard!<\\/strong>",
		Vendor:      "Burton",
		ProductType: "Snowboard",
	}

	returnedProduct, err := client.Product.Create(product)
	if err != nil {
		t.Errorf("Product.Create returned error: %v", err)
	}

	productTests(t, *returnedProduct)
}
