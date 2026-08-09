package bursapay

import (
	"context"
	"fmt"
	"net/url"
)

type PaymentsService struct {
	client *Client
}

type PaymentInitRequest struct {
	Amount      float64                `json:"amount"`
	Email       string                 `json:"email"`
	Currency    string                 `json:"currency,omitempty"`
	Reference   string                 `json:"reference,omitempty"`
	CallbackURL string                 `json:"callback_url,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

type PaymentInitResponse struct {
	AuthorizationURL string `json:"authorization_url"`
	AccessCode       string `json:"access_code"`
	Reference        string `json:"reference"`
}

type PaymentVerification struct {
	ID        int                    `json:"id"`
	Reference string                 `json:"reference"`
	Amount    float64                `json:"amount"`
	Status    string                 `json:"status"`
	Currency  string                 `json:"currency"`
	Email     string                 `json:"customer_email"`
	PaidAt    string                 `json:"paid_at"`
	Metadata  map[string]interface{} `json:"metadata"`
}

func (s *PaymentsService) Initialize(ctx context.Context, req *PaymentInitRequest) (*PaymentInitResponse, error) {
	httpReq, err := s.client.NewRequest(ctx, "POST", "/payments/initialize/", req, nil)
	if err != nil {
		return nil, err
	}

	var res PaymentInitResponse
	if _, err := s.client.Do(httpReq, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (s *PaymentsService) Verify(ctx context.Context, reference string) (*PaymentVerification, error) {
	path := fmt.Sprintf("/payments/verify/%s/", reference)
	httpReq, err := s.client.NewRequest(ctx, "GET", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var res PaymentVerification
	if _, err := s.client.Do(httpReq, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (s *PaymentsService) List(ctx context.Context, status string, page int) ([]PaymentVerification, error) {
	query := url.Values{}
	if status != "" {
		query.Set("status", status)
	}
	if page > 0 {
		query.Set("page", fmt.Sprintf("%d", page))
	}

	httpReq, err := s.client.NewRequest(ctx, "GET", "/payments/", nil, query)
	if err != nil {
		return nil, err
	}

	var res []PaymentVerification
	if _, err := s.client.Do(httpReq, &res); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *PaymentsService) Get(ctx context.Context, idOrRef string) (*PaymentVerification, error) {
	path := fmt.Sprintf("/payments/%s/", idOrRef)
	httpReq, err := s.client.NewRequest(ctx, "GET", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var res PaymentVerification
	if _, err := s.client.Do(httpReq, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
