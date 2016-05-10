package goshopify

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	UserAgent = "goshopify/0.1.0"
)

// Basic app settings such as Api key, secret, scope, and redirect url.
// See oauth.go for OAuth related helper functions.
type App struct {
	ApiKey      string
	ApiSecret   string
	RedirectUrl string
	Scope       string
}

// Client manages communication with the Shopify API.
type Client struct {
	// HTTP client used to communicate with the DO API.
	client *http.Client

	// App settings
	app App

	// Base URL for API requests.
	// This is set on a per-store basis which means that each store must have
	// its own client.
	baseURL *url.URL

	// A permanent access token
	token string

	// Services used for communicating with the API
	Product ProductService
}

// Creates an API request. A relative URL can be provided in urlStr, which will
// be resolved to the BaseURL of the Client. Relative URLS should always be
// specified without a preceding slash. If specified, the value pointed to by
// body is JSON encoded and included in as the request body.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	// Make the full url based on the relative path
	u := c.baseURL.ResolveReference(rel)

	// A bit of JSON ceremony
	var js []byte = nil
	if body != nil {
		js, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), bytes.NewBuffer(js))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", UserAgent)
	if c.token != "" {
		req.Header.Add("X-Shopify-Access-Token", c.token)
	}
	return req, nil
}

// Returns a new Shopify API client with an already authenticated shopname and
// token.
func NewClient(app App, shopName string, token string) *Client {
	httpClient := http.DefaultClient

	baseURL, _ := url.Parse(ShopBaseUrl(shopName))

	c := &Client{client: httpClient, app: app, baseURL: baseURL, token: token}
	c.Product = &ProductServiceOp{client: c}

	return c
}

// Do sends an API request and populates the given interface with the parsed
// response. It does not make sense to call Do without a prepared interface
// instance.
func (c *Client) Do(req *http.Request, v interface{}) error {
	resp, err := c.client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		return err
	}

	// TODO: Error handling
	//if c := resp.StatusCode; c >= 400 {
	//	return nil,
	//}

	//response := newResponse(resp)

	//err = CheckResponse(resp)
	//if err != nil {
	//	return response, err
	//}

	if v != nil {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		err = json.Unmarshal(body, &v)
		if err != nil {
			return err
		}
	}

	return nil
}
