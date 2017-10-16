package goshopify

import (
	"reflect"
	"testing"

	httpmock "gopkg.in/jarcoal/httpmock.v1"
)

func TestVariantList(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/admin/products/1/variants.json",
		httpmock.NewStringResponder(200, `{"variants": [{"id":1},{"id":2}]}`))

	variants, err := client.Variant.List(1, nil)
	if err != nil {
		t.Errorf("Variant.List returned error: %v", err)
	}

	expected := []Variant{{ID: 1}, {ID: 2}}
	if !reflect.DeepEqual(variants, expected) {
		t.Errorf("Variant.List returned %+v, expected %+v", variants, expected)
	}
}
