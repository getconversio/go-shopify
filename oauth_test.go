package goshopify

import (
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestAppAuthorizeUrl(t *testing.T) {
	setup()
	defer teardown()

	cases := []struct {
		shopName string
		nonce    string
		expected string
	}{
		{"fooshop", "thenonce", "https://fooshop.myshopify.com/admin/oauth/authorize?client_id=apikey&redirect_uri=https%3A%2F%2Fexample.com%2Fcallback&scope=read_products&state=thenonce"},
	}

	for _, c := range cases {
		actual := app.AuthorizeUrl(c.shopName, c.nonce)
		if actual != c.expected {
			t.Errorf("App.AuthorizeUrl(): expected %s, actual %s", c.expected, actual)
		}
	}
}

func TestAppGetAccessToken(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", "https://fooshop.myshopify.com/admin/oauth/access_token",
		httpmock.NewStringResponder(200, `{"access_token":"footoken"}`))

	token, err := app.GetAccessToken("fooshop", "foocode")

	if err != nil {
		t.Fatalf("GetAccessToken(): %v", err)
	}

	expected := "footoken"
	if token != expected {
		t.Errorf("Token = %v, expected %v", token, expected)
	}
}
