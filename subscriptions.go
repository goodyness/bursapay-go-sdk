package bursapay

import (
	"context"
	"fmt"
)

type SubscriptionsService struct {
	client *Client
}

type SubscriptionRequest struct {
	CustomerEmail string  `json:"customer_email"`
	PlanCode      string  `json:"plan_code"`
	Amount        float64 `json:"amount,omitempty"`
	Interval      string  `json:"interval,omitempty"` // monthly, yearly, weekly
}

type Subscription struct {
	ID               int     `json:"id"`
	SubscriptionCode string  `json:"subscription_code"`
	CustomerEmail    string  `json:"customer_email"`
	PlanCode         string  `json:"plan_code"`
	Amount           float64 `json:"amount"`
	Interval         string  `json:"interval"`
	Status           string  `json:"status"`
	NextBillingDate  string  `json:"next_billing_date"`
	CreatedAt        string  `json:"created_at"`
}

func (s *SubscriptionsService) Create(ctx context.Context, req *SubscriptionRequest) (*Subscription, error) {
	httpReq, err := s.client.NewRequest(ctx, "POST", "/subscriptions/", req, nil)
	if err != nil {
		return nil, err
	}
	var res Subscription
	if _, err := s.client.Do(httpReq, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (s *SubscriptionsService) List(ctx context.Context) ([]Subscription, error) {
	httpReq, err := s.client.NewRequest(ctx, "GET", "/subscriptions/", nil, nil)
	if err != nil {
		return nil, err
	}
	var res []Subscription
	if _, err := s.client.Do(httpReq, &res); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *SubscriptionsService) Get(ctx context.Context, idOrCode string) (*Subscription, error) {
	path := fmt.Sprintf("/subscriptions/%s/", idOrCode)
	httpReq, err := s.client.NewRequest(ctx, "GET", path, nil, nil)
	if err != nil {
		return nil, err
	}
	var res Subscription
	if _, err := s.client.Do(httpReq, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (s *SubscriptionsService) Cancel(ctx context.Context, idOrCode string) (*Subscription, error) {
	path := fmt.Sprintf("/subscriptions/%s/cancel/", idOrCode)
	httpReq, err := s.client.NewRequest(ctx, "POST", path, nil, nil)
	if err != nil {
		return nil, err
	}
	var res Subscription
	if _, err := s.client.Do(httpReq, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
