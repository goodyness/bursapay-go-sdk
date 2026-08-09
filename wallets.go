package bursapay

import (
	"context"
	"fmt"
)

type WalletsService struct {
	client *Client
}

type WalletBalance struct {
	Currency        string  `json:"currency"`
	AvailableBalance float64 `json:"available_balance"`
	LedgerBalance    float64 `json:"ledger_balance"`
	PendingBalance   float64 `json:"pending_balance"`
}

type LedgerEntry struct {
	ID            int     `json:"id"`
	TransactionType string `json:"transaction_type"`
	Amount        float64 `json:"amount"`
	BalanceAfter  float64 `json:"balance_after"`
	Description   string  `json:"description"`
	CreatedAt     string  `json:"created_at"`
}

func (s *WalletsService) GetBalance(ctx context.Context) (*WalletBalance, error) {
	httpReq, err := s.client.NewRequest(ctx, "GET", "/wallets/balance/", nil, nil)
	if err != nil {
		return nil, err
	}
	var res WalletBalance
	if _, err := s.client.Do(httpReq, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (s *WalletsService) GetLedger(ctx context.Context) ([]LedgerEntry, error) {
	httpReq, err := s.client.NewRequest(ctx, "GET", "/wallets/ledger/", nil, nil)
	if err != nil {
		return nil, err
	}
	var res []LedgerEntry
	if _, err := s.client.Do(httpReq, &res); err != nil {
		return nil, err
	}
	return res, nil
}

type DisputesService struct {
	client *Client
}

type Dispute struct {
	ID        int     `json:"id"`
	Reference string  `json:"reference"`
	Amount    float64 `json:"amount"`
	Reason    string  `json:"reason"`
	Status    string  `json:"status"`
	CreatedAt string  `json:"created_at"`
}

func (s *DisputesService) List(ctx context.Context) ([]Dispute, error) {
	httpReq, err := s.client.NewRequest(ctx, "GET", "/disputes/", nil, nil)
	if err != nil {
		return nil, err
	}
	var res []Dispute
	if _, err := s.client.Do(httpReq, &res); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *DisputesService) Get(ctx context.Context, id string) (*Dispute, error) {
	path := fmt.Sprintf("/disputes/%s/", id)
	httpReq, err := s.client.NewRequest(ctx, "GET", path, nil, nil)
	if err != nil {
		return nil, err
	}
	var res Dispute
	if _, err := s.client.Do(httpReq, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
