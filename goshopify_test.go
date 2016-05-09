package goshopify

import (
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/jarcoal/httpmock"
)

var (
	client *Client
	app    App
)

func setup() {
	app = App{
		ApiKey:      "apikey",
		ApiSecret:   "hush",
		RedirectUrl: "https://example.com/callback",
		Scope:       "read_products",
	}
	client = NewClient(app, "fooshop", "abcd")
	httpmock.Activate()
}

func teardown() {
	httpmock.DeactivateAndReset()
}

func TestNewClient(t *testing.T) {
	c := NewClient(app, "fooshop", "abcd")
	expected := "https://fooshop.myshopify.com"
	if c.baseURL.String() != expected {
		t.Errorf("NewClient BaseURL = %v, expected %v", c.baseURL.String(), expected)
	}
}

func TestNewRequest(t *testing.T) {
	c := NewClient(app, "fooshop", "abcd")

	inURL, outURL := "/foo", "https://fooshop.myshopify.com/foo"
	inBody := struct {
		Hello string `json:"hello"`
	}{Hello: "World"}
	outBody := `{"hello":"World"}`
	req, _ := c.NewRequest("GET", inURL, inBody)

	// Test relative URL was expanded
	if req.URL.String() != outURL {
		t.Errorf("NewRequest(%v) URL = %v, expected %v", inURL, req.URL, outURL)
	}

	// Test body was JSON encoded
	body, _ := ioutil.ReadAll(req.Body)
	if string(body) != outBody {
		t.Errorf("NewRequest(%v)Body = %v, expected %v", inBody, string(body), outBody)
	}

	// Test user-agent is attached to the request
	userAgent := req.Header.Get("User-Agent")
	if userAgent != UserAgent {
		t.Errorf("NewRequest() User-Agent = %v, expected %v", userAgent, UserAgent)
	}

	// Test token is attached to the request
	token := req.Header.Get("X-Shopify-Access-Token")
	expected := "abcd"
	if token != expected {
		t.Errorf("NewRequest() X-Shopify-Access-Token = %v, expected %v", token, expected)
	}
}

func TestNewRequestMissingToken(t *testing.T) {
	c := NewClient(app, "fooshop", "")

	req, _ := c.NewRequest("GET", "/foo", nil)

	// Test token is not attached to the request
	token := req.Header["X-Shopify-Access-Token"]
	if token != nil {
		t.Errorf("NewRequest() X-Shopify-Access-Token = %v, expected %v", token, nil)
	}
}

func TestDo(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com",
		httpmock.NewStringResponder(200, `{"Foo":"bar"}`))

	type foo struct {
		Foo string
	}

	req, _ := client.NewRequest("GET", "", nil)
	body := new(foo)
	err := client.Do(req, body)
	if err != nil {
		t.Fatalf("Do(): %v", err)
	}

	expected := &foo{"bar"}
	if !reflect.DeepEqual(body, expected) {
		t.Errorf("Response body = %v, expected %v", body, expected)
	}
}
