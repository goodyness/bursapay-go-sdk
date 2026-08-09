package bursapay

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	DefaultBaseURL = "https://api.bursapay.dev/api/v1"
	UserAgent      = "BursaPay-Go-SDK/1.0.0"
)

type Response struct {
	Status  bool            `json:"status"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data,omitempty"`
}

type Client struct {
	HTTPClient *http.Client
	BaseURL    string
	APIKey     string
	UserAgent  string

	Payments        *PaymentsService
	Customers       *CustomersService
	Refunds         *RefundsService
	Transfers       *TransfersService
	VirtualAccounts *VirtualAccountsService
	PaymentLinks    *PaymentLinksService
	Subscriptions   *SubscriptionsService
	Invoices        *InvoicesService
	Wallets         *WalletsService
	Disputes        *DisputesService
	Webhooks        *WebhooksService
}

func NewClient(apiKey string, opts ...Option) *Client {
	c := &Client{
		HTTPClient: &http.Client{Timeout: 30 * time.Second},
		BaseURL:    DefaultBaseURL,
		APIKey:     apiKey,
		UserAgent:  UserAgent,
	}

	for _, opt := range opts {
		opt(c)
	}

	c.Payments = &PaymentsService{client: c}
	c.Customers = &CustomersService{client: c}
	c.Refunds = &RefundsService{client: c}
	c.Transfers = &TransfersService{client: c}
	c.VirtualAccounts = &VirtualAccountsService{client: c}
	c.PaymentLinks = &PaymentLinksService{client: c}
	c.Subscriptions = &SubscriptionsService{client: c}
	c.Invoices = &InvoicesService{client: c}
	c.Wallets = &WalletsService{client: c}
	c.Disputes = &DisputesService{client: c}
	c.Webhooks = &WebhooksService{client: c}

	return c
}

type Option func(*Client)

func WithBaseURL(url string) Option {
	return func(c *Client) {
		c.BaseURL = strings.TrimSuffix(url, "/")
	}
}

func WithHTTPClient(client *http.Client) Option {
	return func(c *Client) {
		c.HTTPClient = client
	}
}

func (c *Client) NewRequest(ctx context.Context, method, path string, body interface{}, query url.Values) (*http.Request, error) {
	relPath := strings.TrimPrefix(path, "/")
	u := fmt.Sprintf("%s/%s", strings.TrimSuffix(c.BaseURL, "/"), relPath)

	if len(query) > 0 {
		u = fmt.Sprintf("%s?%s", u, query.Encode())
	}

	var buf io.Reader
	if body != nil {
		jsonBytes, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("bursapay: failed to marshal request body: %w", err)
		}
		buf = bytes.NewBuffer(jsonBytes)
	}

	req, err := http.NewRequestWithContext(ctx, method, u, buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)

	return req, nil
}

func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("bursapay: network error: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("bursapay: failed to read response body: %w", err)
	}

	if resp.StatusCode >= 400 {
		apiErr := APIError{StatusCode: resp.StatusCode}
		_ = json.Unmarshal(respBody, &apiErr)

		switch resp.StatusCode {
		case 401, 403:
			return nil, &AuthenticationError{APIError: apiErr}
		case 400, 422:
			return nil, &InvalidRequestError{APIError: apiErr}
		case 429:
			return nil, &RateLimitError{APIError: apiErr}
		default:
			return nil, &apiErr
		}
	}

	var apiResp Response
	if len(respBody) > 0 {
		if err := json.Unmarshal(respBody, &apiResp); err != nil {
			return nil, fmt.Errorf("bursapay: failed to parse response JSON: %w", err)
		}
	}

	if v != nil && len(apiResp.Data) > 0 {
		if err := json.Unmarshal(apiResp.Data, v); err != nil {
			return nil, fmt.Errorf("bursapay: failed to unmarshal response data: %w", err)
		}
	}

	return &apiResp, nil
}
