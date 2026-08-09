# BursaPay Go SDK

Official Go SDK for integrating with the **BursaPay Payment Gateway API**.

---

## Installation

```bash
go get github.com/goodyness/bursapay-go-sdk
```

---

## Quick Start

### Initialize Client

```go
package main

import (
    "context"
    "fmt"
    "log"
    "os"

    "github.com/goodyness/bursapay-go-sdk"
)

func main() {
	client := bursapay.NewClient(os.Getenv("BURSA_PAY_SECRET_KEY"))

	ctx := context.Background()

	// 1. Initialize a Payment
	payment, err := client.Payments.Initialize(ctx, &bursapay.PaymentInitRequest{
		Amount:   5000.00,
		Email:    "customer@example.com",
		Currency: "NGN",
		CallbackURL: "https://example.com/callback",
	})
	if err != nil {
		log.Fatalf("Payment init failed: %v", err)
	}

	fmt.Printf("Checkout URL: %s\n", payment.AuthorizationURL)
	fmt.Printf("Reference: %s\n", payment.Reference)
}
```

---

## Features & Usage

### 2. Verify Payment

```go
verification, err := client.Payments.Verify(ctx, "BP_REF_123456789")
if err != nil {
    log.Fatal(err)
}

if verification.Status == "success" {
    fmt.Printf("Payment of ₦%.2f verified!\n", verification.Amount)
}
```

### 3. Dedicated Virtual Accounts

```go
account, err := client.VirtualAccounts.Create(ctx, &bursapay.VirtualAccountRequest{
    CustomerEmail: "merchant@example.com",
    BVN:           "12345678901",
    PreferredBank: "Wema Bank",
})
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Virtual Account: %s (%s)\n", account.AccountNumber, account.BankName)
```

### 4. Verify Webhook Signatures Safely

```go
package main

import (
	"io"
	"net/http"
	"os"

	"github.com/goodyness/bursapay-go-sdk"
)

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	payload, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read body", http.StatusBadRequest)
		return
	}

	sigHeader := r.Header.Get("X-BursaPay-Signature")
	tsHeader := r.Header.Get("X-BursaPay-Timestamp")
	secret := os.Getenv("BURSAPAY_WEBHOOK_SECRET")

	event, err := bursapay.ConstructEvent(payload, sigHeader, tsHeader, secret)
	if err != nil {
		http.Error(w, "Webhook verification failed: "+err.Error(), http.StatusBadRequest)
		return
	}

	switch event.Event {
	case "payment.success":
		// Process payment...
	case "virtual_account.credited":
		// Credit user account...
	}

	w.WriteHeader(http.StatusOK)
}
```

---

## License

MIT License. See `LICENSE` for details.
