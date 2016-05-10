package goshopify

import (
	"reflect"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestProductList(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/admin/products.json",
		httpmock.NewStringResponder(200, `{"products": [{"id":1},{"id":2}]}`))

	products, err := client.Product.List()
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

	cnt, err := client.Product.Count()
	if err != nil {
		t.Errorf("Product.Count returned error: %v", err)
	}

	expected := 3
	if cnt != expected {
		t.Errorf("Product.Count returned %d, expected %d", cnt, expected)
	}
}

func TestProductGet(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/admin/products/1.json",
		httpmock.NewStringResponder(200, `{"product": {"id":1}}`))

	product, err := client.Product.Get(1)
	if err != nil {
		t.Errorf("Product.Get returned error: %v", err)
	}

	expected := &Product{ID: 1}
	if !reflect.DeepEqual(product, expected) {
		t.Errorf("Product.Get returned %+v, expected %+v", product, expected)
	}
}
