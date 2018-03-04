package goshopify

import (
	"reflect"
	"testing"

	"gopkg.in/jarcoal/httpmock.v1"
)

func TestThemeList(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder(
		"GET",
		"https://fooshop.myshopify.com/admin/themes.json",
		httpmock.NewStringResponder(
			200,
			`{"themes": [{"id":1},{"id":2}]}`,
		),
	)

	httpmock.RegisterResponder(
		"GET",
		"https://fooshop.myshopify.com/admin/themes.json?role=main",
		httpmock.NewStringResponder(
			200,
			`{"themes": [{"id":1}]}`,
		),
	)

	themes, err := client.Theme.List(nil)
	if err != nil {
		t.Errorf("Theme.List returned error: %v", err)
	}

	expected := []Theme{{ID: 1}, {ID: 2}}
	if !reflect.DeepEqual(themes, expected) {
		t.Errorf("Theme.List returned %+v, expected %+v", themes, expected)
	}

	themes, err = client.Theme.List(ThemeListOptions{Role: "main"})
	if err != nil {
		t.Errorf("Theme.List returned error: %v", err)
	}

	expected = []Theme{{ID: 1}}
	if !reflect.DeepEqual(themes, expected) {
		t.Errorf("Theme.List returned %+v, expected %+v", themes, expected)
	}
}
