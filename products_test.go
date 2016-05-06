package goshopify

import (
	"reflect"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestProductsList(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/admin/products.json",
		httpmock.NewStringResponder(200, `{"products": [{"id":1},{"id":2}]}`))

	products, err := client.Products.List()
	if err != nil {
		t.Errorf("Products.List returned error: %v", err)
	}

	expected := []Product{{ID: 1}, {ID: 2}}
	if !reflect.DeepEqual(products, expected) {
		t.Errorf("Products.List returned %+v, expected %+v", products, expected)
	}
}

func TestProductsCount(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/admin/products/count.json",
		httpmock.NewStringResponder(200, `{"count": 3}`))

	cnt, err := client.Products.Count()
	if err != nil {
		t.Errorf("Products.Count returned error: %v", err)
	}

	expected := 3
	if cnt != expected {
		t.Errorf("Products.Count returned %d, expected %d", cnt, expected)
	}
}

func TestProductsGet(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/admin/products/1.json",
		httpmock.NewStringResponder(200, `{"product": {"id":1}}`))

	product, err := client.Products.Get(1)
	if err != nil {
		t.Errorf("Products.Get returned error: %v", err)
	}

	expected := &Product{ID: 1}
	if !reflect.DeepEqual(product, expected) {
		t.Errorf("Products.Get returned %+v, expected %+v", product, expected)
	}
}
