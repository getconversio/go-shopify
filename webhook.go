package goshopify

import (
	"fmt"
	"time"
	"net/http"
	"crypto/hmac"
	"crypto/sha256"
	"io/ioutil"
	"bytes"
	"encoding/base64"
)

const (
	webhooksBasePath = "admin/webhooks"
	shopifyChecksumHeader = "X-Shopify-Hmac-Sha256"
)

// WebhookService is an interface for interfacing with the webhook endpoints of
// the Shopify API.
// See: https://help.shopify.com/api/reference/webhook
type WebhookService interface {
	List(interface{}) ([]Webhook, error)
	Count(interface{}) (int, error)
	Get(int, interface{}) (*Webhook, error)
	Create(Webhook) (*Webhook, error)
	Update(Webhook) (*Webhook, error)
	Delete(int) error
	Verify(*http.Request) bool
}

// WebhookServiceOp handles communication with the webhook-related methods of
// the Shopify API.
type WebhookServiceOp struct {
	client *Client
}

// Webhook represents a Shopify webhook
type Webhook struct {
	ID                  int        `json:"id"`
	Address             string     `json:"address"`
	Topic               string     `json:"topic"`
	Format              string     `json:"format"`
	CreatedAt           *time.Time `json:"created_at,omitempty"`
	UpdatedAt           *time.Time `json:"updated_at,omitempty"`
	Fields              []string   `json:"fields"`
	MetafieldNamespaces []string   `json:"metafield_namespaces"`
}

// WebhookOptions can be used for filtering webhooks on a List request.
type WebhookOptions struct {
	Address string `url:"address,omitempty"`
	Topic   string `url:"topic,omitempty"`
}

// WebhookResource represents the result from the admin/webhooks.json endpoint
type WebhookResource struct {
	Webhook *Webhook `json:"webhook"`
}

// WebhooksResource is the root object for a webhook get request.
type WebhooksResource struct {
	Webhooks []Webhook `json:"webhooks"`
}

// List webhooks
func (s *WebhookServiceOp) List(options interface{}) ([]Webhook, error) {
	path := fmt.Sprintf("%s.json", webhooksBasePath)
	resource := new(WebhooksResource)
	err := s.client.Get(path, resource, options)
	return resource.Webhooks, err
}

// Count webhooks
func (s *WebhookServiceOp) Count(options interface{}) (int, error) {
	path := fmt.Sprintf("%s/count.json", webhooksBasePath)
	return s.client.Count(path, options)
}

// Get individual webhook
func (s *WebhookServiceOp) Get(webhookdID int, options interface{}) (*Webhook, error) {
	path := fmt.Sprintf("%s/%d.json", webhooksBasePath, webhookdID)
	resource := new(WebhookResource)
	err := s.client.Get(path, resource, options)
	return resource.Webhook, err
}

// Create a new webhook
func (s *WebhookServiceOp) Create(webhook Webhook) (*Webhook, error) {
	path := fmt.Sprintf("%s.json", webhooksBasePath)
	wrappedData := WebhookResource{Webhook: &webhook}
	resource := new(WebhookResource)
	err := s.client.Post(path, wrappedData, resource)
	return resource.Webhook, err
}

// Update an existing webhook.
func (s *WebhookServiceOp) Update(webhook Webhook) (*Webhook, error) {
	path := fmt.Sprintf("%s/%d.json", webhooksBasePath, webhook.ID)
	wrappedData := WebhookResource{Webhook: &webhook}
	resource := new(WebhookResource)
	err := s.client.Put(path, wrappedData, resource)
	return resource.Webhook, err
}

// Delete an existing webhooks
func (s *WebhookServiceOp) Delete(ID int) error {
	return s.client.Delete(fmt.Sprintf("%s/%d.json", webhooksBasePath, ID))
}

// Verify a webhook created through the Shopify API
func (s *WebhookServiceOp) Verify(httpRequest *http.Request) bool {
	shopifySha256 := httpRequest.Header.Get(shopifyChecksumHeader)
	actualMac := []byte(shopifySha256)

	shopifySecret := s.client.app.ApiSecret
	mac := hmac.New(sha256.New, []byte(shopifySecret))
	requestBody, _ := ioutil.ReadAll(httpRequest.Body)
	httpRequest.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))
	mac.Write(requestBody)
	macSum := mac.Sum(nil)
	expectedMac := []byte(base64.StdEncoding.EncodeToString(macSum))

	return hmac.Equal(actualMac, expectedMac)
}
