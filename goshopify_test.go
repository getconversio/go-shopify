package goshopify

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
	"time"

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

	inURL, outURL := "foo?page=1", "https://fooshop.myshopify.com/foo?limit=10&page=1"
	inBody := struct {
		Hello string `json:"hello"`
	}{Hello: "World"}
	outBody := `{"hello":"World"}`

	type extraOptions struct {
		Limit int `url:"limit"`
	}

	req, _ := c.NewRequest("GET", inURL, inBody, extraOptions{Limit: 10})

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

	req, _ := c.NewRequest("GET", "/foo", nil, nil)

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
		req, _ := client.NewRequest("GET", c.url, nil, nil)
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
			httpmock.NewStringResponse(500, `{"errors": "This action requires read_customers scope"}`),
			ResponseError{Status: 500, Message: "This action requires read_customers scope"},
		},
		{
			httpmock.NewStringResponse(500, `{"errors": ["not", "very good"]}`),
			ResponseError{Status: 500, Message: "not, very good", Errors: []string{"not", "very good"}},
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

func TestCount(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/foocount",
		httpmock.NewStringResponder(200, `{"count": 5}`))

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/foocount?created_at_min=2016-01-01T00%3A00%3A00Z",
		httpmock.NewStringResponder(200, `{"count": 2}`))

	// Test without options
	cnt, err := client.Count("foocount", nil)
	if err != nil {
		t.Errorf("Client.Count returned error: %v", err)
	}

	expected := 5
	if cnt != expected {
		t.Errorf("Client.Count returned %d, expected %d", cnt, expected)
	}

	// Test with options
	date := time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC)
	cnt, err = client.Count("foocount", CountOptions{CreatedAtMin: date})
	if err != nil {
		t.Errorf("Client.Count returned error: %v", err)
	}

	expected = 2
	if cnt != expected {
		t.Errorf("Client.Count returned %d, expected %d", cnt, expected)
	}
}
