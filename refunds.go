package bursapay

import (
	"context"
	"fmt"
)

type RefundsService struct {
	client *Client
}

type RefundRequest struct {
	PaymentReference string  `json:"payment_reference"`
	Amount           float64 `json:"amount,omitempty"`
	Reason           string  `json:"reason,omitempty"`
}

type Refund struct {
	ID               int     `json:"id"`
	RefundReference string  `json:"refund_reference"`
	PaymentReference string  `json:"payment_reference"`
	Amount           float64 `json:"amount"`
	Status           string  `json:"status"`
	CreatedAt        string  `json:"created_at"`
}

func (s *RefundsService) Create(ctx context.Context, req *RefundRequest) (*Refund, error) {
	httpReq, err := s.client.NewRequest(ctx, "POST", "/refunds/", req, nil)
	if err != nil {
		return nil, err
	}
	var res Refund
	if _, err := s.client.Do(httpReq, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (s *RefundsService) List(ctx context.Context) ([]Refund, error) {
	httpReq, err := s.client.NewRequest(ctx, "GET", "/refunds/", nil, nil)
	if err != nil {
		return nil, err
	}
	var res []Refund
	if _, err := s.client.Do(httpReq, &res); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *RefundsService) Get(ctx context.Context, id string) (*Refund, error) {
	path := fmt.Sprintf("/refunds/%s/", id)
	httpReq, err := s.client.NewRequest(ctx, "GET", path, nil, nil)
	if err != nil {
		return nil, err
	}
	var res Refund
	if _, err := s.client.Do(httpReq, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
