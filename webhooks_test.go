package bursapay_test

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/bursapay/bursapay-go"
)

func TestVerifySignatureValid(t *testing.T) {
	payload := []byte(`{"event":"payment.success","data":{"id":101}}`)
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	secret := "secret_key_9988"

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(fmt.Sprintf("%s.%s", timestamp, string(payload))))
	signature := hex.EncodeToString(mac.Sum(nil))

	err := bursapay.VerifySignature(payload, signature, timestamp, secret, 300)
	if err != nil {
		t.Fatalf("expected valid signature verification, got err: %v", err)
	}

	event, err := bursapay.ConstructEvent(payload, signature, timestamp, secret)
	if err != nil {
		t.Fatalf("expected construct event success, got err: %v", err)
	}

	if event.Event != "payment.success" {
		t.Fatalf("expected event type payment.success, got: %s", event.Event)
	}
}

func TestVerifySignatureInvalid(t *testing.T) {
	payload := []byte(`{"event":"payment.success"}`)
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	secret := "secret_key_9988"
	badSignature := "invalid_signature_hash"

	err := bursapay.VerifySignature(payload, badSignature, timestamp, secret, 300)
	if err == nil {
		t.Fatal("expected error for invalid signature, got nil")
	}
}
