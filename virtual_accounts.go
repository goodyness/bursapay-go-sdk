package bursapay

import (
	"context"
	"fmt"
)

type VirtualAccountsService struct {
	client *Client
}

type VirtualAccountRequest struct {
	CustomerEmail string `json:"customer_email"`
	BVN           string `json:"bvn,omitempty"`
	PreferredBank string `json:"preferred_bank,omitempty"`
}

type VirtualAccount struct {
	ID            int    `json:"id"`
	AccountNumber string `json:"account_number"`
	AccountName   string `json:"account_name"`
	BankName      string `json:"bank_name"`
	CustomerEmail string `json:"customer_email"`
	Status        string `json:"status"`
	CreatedAt     string `json:"created_at"`
}

func (s *VirtualAccountsService) Create(ctx context.Context, req *VirtualAccountRequest) (*VirtualAccount, error) {
	httpReq, err := s.client.NewRequest(ctx, "POST", "/virtual_accounts/", req, nil)
	if err != nil {
		return nil, err
	}
	var res VirtualAccount
	if _, err := s.client.Do(httpReq, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (s *VirtualAccountsService) List(ctx context.Context) ([]VirtualAccount, error) {
	httpReq, err := s.client.NewRequest(ctx, "GET", "/virtual_accounts/", nil, nil)
	if err != nil {
		return nil, err
	}
	var res []VirtualAccount
	if _, err := s.client.Do(httpReq, &res); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *VirtualAccountsService) Get(ctx context.Context, id string) (*VirtualAccount, error) {
	path := fmt.Sprintf("/virtual_accounts/%s/", id)
	httpReq, err := s.client.NewRequest(ctx, "GET", path, nil, nil)
	if err != nil {
		return nil, err
	}
	var res VirtualAccount
	if _, err := s.client.Do(httpReq, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
