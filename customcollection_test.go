package goshopify

import (
	"reflect"
	"testing"
	"time"

	httpmock "gopkg.in/jarcoal/httpmock.v1"
)

func collectionTests(t *testing.T, collection CustomCollection) {

	// Test a few fields
	cases := []struct {
		field    string
		expected interface{}
		actual   interface{}
	}{
		{"ID", 30497275952, collection.ID},
		{"Handle", "macbooks", collection.Handle},
		{"Title", "Macbooks", collection.Title},
		{"BodyHTML", "Macbook Body", collection.BodyHTML},
		{"SortOrder", "best-selling", collection.SortOrder},
	}

	for _, c := range cases {
		if c.expected != c.actual {
			t.Errorf("CustomCollection.%v returned %v, expected %v", c.field, c.actual, c.expected)
		}
	}
}

func TestCustomCollectionList(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/admin/custom_collections.json",
		httpmock.NewStringResponder(200, `{"custom_collections": [{"id":1},{"id":2}]}`))

	products, err := client.CustomCollection.List(nil)
	if err != nil {
		t.Errorf("CustomCollection.List returned error: %v", err)
	}

	expected := []CustomCollection{{ID: 1}, {ID: 2}}
	if !reflect.DeepEqual(products, expected) {
		t.Errorf("CustomCollection.List returned %+v, expected %+v", products, expected)
	}
}

func TestCustomCollectionCount(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/admin/custom_collections/count.json",
		httpmock.NewStringResponder(200, `{"count": 5}`))

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/admin/custom_collections/count.json?created_at_min=2016-01-01T00%3A00%3A00Z",
		httpmock.NewStringResponder(200, `{"count": 2}`))

	cnt, err := client.CustomCollection.Count(nil)
	if err != nil {
		t.Errorf("CustomCollection.Count returned error: %v", err)
	}

	expected := 5
	if cnt != expected {
		t.Errorf("CustomCollection.Count returned %d, expected %d", cnt, expected)
	}

	date := time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC)
	cnt, err = client.CustomCollection.Count(CountOptions{CreatedAtMin: date})
	if err != nil {
		t.Errorf("CustomCollection.Count returned error: %v", err)
	}

	expected = 2
	if cnt != expected {
		t.Errorf("CustomCollection.Count returned %d, expected %d", cnt, expected)
	}
}

func TestCustomCollectionGet(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/admin/custom_collections/1.json",
		httpmock.NewStringResponder(200, `{"custom_collection": {"id":1}}`))

	product, err := client.CustomCollection.Get(1, nil)
	if err != nil {
		t.Errorf("CustomCollection.Get returned error: %v", err)
	}

	expected := &CustomCollection{ID: 1}
	if !reflect.DeepEqual(product, expected) {
		t.Errorf("CustomCollection.Get returned %+v, expected %+v", product, expected)
	}
}

func TestCustomCollectionCreate(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", "https://fooshop.myshopify.com/admin/custom_collections.json",
		httpmock.NewBytesResponder(200, loadFixture("customcollection.json")))

	collection := CustomCollection{
		Title: "Macbooks",
	}

	returnedCollection, err := client.CustomCollection.Create(collection)
	if err != nil {
		t.Errorf("CustomCollection.Create returned error: %v", err)
	}

	collectionTests(t, *returnedCollection)
}

func TestCustomCollectionUpdate(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("PUT", "https://fooshop.myshopify.com/admin/custom_collections/1.json",
		httpmock.NewBytesResponder(200, loadFixture("customcollection.json")))

	collection := CustomCollection{
		ID:    1,
		Title: "Macbooks",
	}

	returnedCollection, err := client.CustomCollection.Update(collection.ID, collection)
	if err != nil {
		t.Errorf("CustomCollection.Update returned error: %v", err)
	}

	collectionTests(t, *returnedCollection)
}

func TestCustomCollectionDelete(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("DELETE", "https://fooshop.myshopify.com/admin/custom_collections/1.json",
		httpmock.NewStringResponder(200, "{}"))

	err := client.CustomCollection.Delete(1)
	if err != nil {
		t.Errorf("CustomCollection.Delete returned error: %v", err)
	}
}
