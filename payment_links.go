package bursapay

import (
	"context"
	"fmt"
)

type PaymentLinksService struct {
	client *Client
}

type PaymentLinkRequest struct {
	Title       string  `json:"title"`
	Description string  `json:"description,omitempty"`
	Amount      float64 `json:"amount"`
	Currency    string  `json:"currency,omitempty"`
	RedirectURL string  `json:"redirect_url,omitempty"`
}

type PaymentLink struct {
	ID          int     `json:"id"`
	Slug        string  `json:"slug"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
	Currency    string  `json:"currency"`
	URL         string  `json:"url"`
	IsActive    bool    `json:"is_active"`
	CreatedAt   string  `json:"created_at"`
}

func (s *PaymentLinksService) Create(ctx context.Context, req *PaymentLinkRequest) (*PaymentLink, error) {
	httpReq, err := s.client.NewRequest(ctx, "POST", "/payment_links/", req, nil)
	if err != nil {
		return nil, err
	}
	var res PaymentLink
	if _, err := s.client.Do(httpReq, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (s *PaymentLinksService) List(ctx context.Context) ([]PaymentLink, error) {
	httpReq, err := s.client.NewRequest(ctx, "GET", "/payment_links/", nil, nil)
	if err != nil {
		return nil, err
	}
	var res []PaymentLink
	if _, err := s.client.Do(httpReq, &res); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *PaymentLinksService) Get(ctx context.Context, idOrSlug string) (*PaymentLink, error) {
	path := fmt.Sprintf("/payment_links/%s/", idOrSlug)
	httpReq, err := s.client.NewRequest(ctx, "GET", path, nil, nil)
	if err != nil {
		return nil, err
	}
	var res PaymentLink
	if _, err := s.client.Do(httpReq, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
