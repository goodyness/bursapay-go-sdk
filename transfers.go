package bursapay

import (
	"context"
	"fmt"
)

type TransfersService struct {
	client *Client
}

type TransferRequest struct {
	Amount        float64 `json:"amount"`
	RecipientBank string  `json:"recipient_bank"`
	AccountNumber string  `json:"account_number"`
	AccountName   string  `json:"account_name"`
	Narration     string  `json:"narration,omitempty"`
	Currency      string  `json:"currency,omitempty"`
}

type Transfer struct {
	ID               int     `json:"id"`
	TransferReference string  `json:"transfer_reference"`
	Amount           float64 `json:"amount"`
	RecipientBank    string  `json:"recipient_bank"`
	AccountNumber    string  `json:"account_number"`
	AccountName      string  `json:"account_name"`
	Status           string  `json:"status"`
	CreatedAt        string  `json:"created_at"`
}

func (s *TransfersService) Initiate(ctx context.Context, req *TransferRequest) (*Transfer, error) {
	httpReq, err := s.client.NewRequest(ctx, "POST", "/transfers/", req, nil)
	if err != nil {
		return nil, err
	}
	var res Transfer
	if _, err := s.client.Do(httpReq, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (s *TransfersService) List(ctx context.Context) ([]Transfer, error) {
	httpReq, err := s.client.NewRequest(ctx, "GET", "/transfers/", nil, nil)
	if err != nil {
		return nil, err
	}
	var res []Transfer
	if _, err := s.client.Do(httpReq, &res); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *TransfersService) Get(ctx context.Context, id string) (*Transfer, error) {
	path := fmt.Sprintf("/transfers/%s/", id)
	httpReq, err := s.client.NewRequest(ctx, "GET", path, nil, nil)
	if err != nil {
		return nil, err
	}
	var res Transfer
	if _, err := s.client.Do(httpReq, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
