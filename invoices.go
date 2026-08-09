package bursapay

import (
	"context"
	"fmt"
)

type InvoicesService struct {
	client *Client
}

type InvoiceItem struct {
	Name     string  `json:"name"`
	Amount   float64 `json:"amount"`
	Quantity int     `json:"quantity"`
}

type InvoiceRequest struct {
	CustomerEmail string        `json:"customer_email"`
	DueDate       string        `json:"due_date"`
	Items         []InvoiceItem `json:"items"`
	Description   string        `json:"description,omitempty"`
}

type Invoice struct {
	ID            int           `json:"id"`
	InvoiceCode   string        `json:"invoice_code"`
	CustomerEmail string        `json:"customer_email"`
	Amount        float64       `json:"amount"`
	Status        string        `json:"status"`
	DueDate       string        `json:"due_date"`
	Items         []InvoiceItem `json:"items"`
	CreatedAt     string        `json:"created_at"`
}

func (s *InvoicesService) Create(ctx context.Context, req *InvoiceRequest) (*Invoice, error) {
	httpReq, err := s.client.NewRequest(ctx, "POST", "/invoices/", req, nil)
	if err != nil {
		return nil, err
	}
	var res Invoice
	if _, err := s.client.Do(httpReq, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (s *InvoicesService) List(ctx context.Context) ([]Invoice, error) {
	httpReq, err := s.client.NewRequest(ctx, "GET", "/invoices/", nil, nil)
	if err != nil {
		return nil, err
	}
	var res []Invoice
	if _, err := s.client.Do(httpReq, &res); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *InvoicesService) Get(ctx context.Context, idOrCode string) (*Invoice, error) {
	path := fmt.Sprintf("/invoices/%s/", idOrCode)
	httpReq, err := s.client.NewRequest(ctx, "GET", path, nil, nil)
	if err != nil {
		return nil, err
	}
	var res Invoice
	if _, err := s.client.Do(httpReq, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (s *InvoicesService) Send(ctx context.Context, idOrCode string) (*Invoice, error) {
	path := fmt.Sprintf("/invoices/%s/send/", idOrCode)
	httpReq, err := s.client.NewRequest(ctx, "POST", path, nil, nil)
	if err != nil {
		return nil, err
	}
	var res Invoice
	if _, err := s.client.Do(httpReq, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
