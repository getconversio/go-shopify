package goshopify

import (
	"fmt"
	"io/ioutil"
	"net/http"
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

	type MyStruct struct {
		Foo string `json:"foo"`
	}

	cases := []struct {
		url       string
		responder httpmock.Responder
		expected  interface{}
	}{
		{
			"foo/1",
			httpmock.NewStringResponder(200, `{"foo": "bar"}`),
			&MyStruct{Foo: "bar"},
		},
		{
			"foo/2",
			httpmock.NewStringResponder(404, `{"error": "does not exist"}`),
			ResponseError{Status: 404, Message: "does not exist"},
		},
		{
			"foo/3",
			httpmock.NewStringResponder(400, `{"errors": {"title": ["wrong"]}}`),
			ResponseError{Status: 400, Message: "wrong", Errors: []string{"title: wrong"}},
		},
	}

	for _, c := range cases {
		shopUrl := fmt.Sprintf("https://fooshop.myshopify.com/%v", c.url)
		httpmock.RegisterResponder("GET", shopUrl, c.responder)

		body := new(MyStruct)
		req, _ := client.NewRequest("GET", c.url, nil)
		err := client.Do(req, body)

		if err != nil && !reflect.DeepEqual(err, c.expected) {
			t.Errorf("Do(): expected error %#v, actual %#v", c.expected, err)
		} else if err == nil && !reflect.DeepEqual(body, c.expected) {
			t.Errorf("Do(): expected %#v, actual %#v", c.expected, body)
		}
	}
}

func TestResponseErrorError(t *testing.T) {
	cases := []struct {
		err      ResponseError
		expected string
	}{
		{
			ResponseError{Message: "oh no"},
			"oh no",
		},
		{
			ResponseError{},
			"Unknown Error",
		},
		{
			ResponseError{Errors: []string{"title: not a valid title"}},
			"title: not a valid title",
		},
		{
			ResponseError{Errors: []string{
				"not a valid title",
				"not a valid description",
			}},
			// The strings are sorted description comes first
			"not a valid description, not a valid title",
		},
	}

	for _, c := range cases {
		actual := fmt.Sprint(c.err)
		if actual != c.expected {
			t.Errorf("ResponseError.Error(): expected %s, actual %s", c.expected, actual)
		}
	}
}

func TestCheckResponseError(t *testing.T) {
	cases := []struct {
		resp     *http.Response
		expected error
	}{
		{
			httpmock.NewStringResponse(200, `{"foo": "bar"}`),
			nil,
		},
		{
			httpmock.NewStringResponse(299, `{"foo": "bar"}`),
			nil,
		},
		{
			httpmock.NewStringResponse(400, `{"error": "bad request"}`),
			ResponseError{Status: 400, Message: "bad request"},
		},
		{
			httpmock.NewStringResponse(500, `{"error": "terrible error"}`),
			ResponseError{Status: 500, Message: "terrible error"},
		},
		{
			httpmock.NewStringResponse(400, `{"errors": { "order": ["order is wrong"] }}`),
			ResponseError{Status: 400, Message: "order is wrong", Errors: []string{"order: order is wrong"}},
		},
	}

	for _, c := range cases {
		actual := CheckResponseError(c.resp)
		if !reflect.DeepEqual(actual, c.expected) {
			t.Errorf("CheckResponseError(): expected %#v, actual %#v", c.expected, actual)
		}
	}
}
