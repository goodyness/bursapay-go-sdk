package bursapay

import (
	"context"
	"fmt"
)

type CustomersService struct {
	client *Client
}

type CustomerRequest struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Phone     string `json:"phone,omitempty"`
}

type Customer struct {
	ID        int    `json:"id"`
	CustomerCode string `json:"customer_code"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	CreatedAt string `json:"created_at"`
}

func (s *CustomersService) Create(ctx context.Context, req *CustomerRequest) (*Customer, error) {
	httpReq, err := s.client.NewRequest(ctx, "POST", "/customers/", req, nil)
	if err != nil {
		return nil, err
	}
	var res Customer
	if _, err := s.client.Do(httpReq, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (s *CustomersService) List(ctx context.Context) ([]Customer, error) {
	httpReq, err := s.client.NewRequest(ctx, "GET", "/customers/", nil, nil)
	if err != nil {
		return nil, err
	}
	var res []Customer
	if _, err := s.client.Do(httpReq, &res); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *CustomersService) Get(ctx context.Context, idOrCode string) (*Customer, error) {
	path := fmt.Sprintf("/customers/%s/", idOrCode)
	httpReq, err := s.client.NewRequest(ctx, "GET", path, nil, nil)
	if err != nil {
		return nil, err
	}
	var res Customer
	if _, err := s.client.Do(httpReq, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
