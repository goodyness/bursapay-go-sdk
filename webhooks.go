package bursapay

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"time"
)

type WebhooksService struct {
	client *Client
}

type WebhookEvent struct {
	Event string          `json:"event"`
	Data  json.RawMessage `json:"data"`
}

// VerifySignature validates HMAC-SHA256 signature and timestamp header against replay attacks.
func VerifySignature(payload []byte, signatureHeader, timestampHeader, webhookSecret string, toleranceSeconds int64) error {
	if signatureHeader == "" || timestampHeader == "" {
		return &SignatureVerificationError{Message: "missing signature or timestamp header"}
	}

	if toleranceSeconds <= 0 {
		toleranceSeconds = 300 // default 5 minutes
	}

	ts, err := strconv.ParseInt(timestampHeader, 10, 64)
	if err != nil {
		return &SignatureVerificationError{Message: "invalid timestamp header format"}
	}

	now := time.Now().Unix()
	if math.Abs(float64(now-ts)) > float64(toleranceSeconds) {
		return &SignatureVerificationError{Message: "webhook timestamp is outside tolerance limits"}
	}

	expectedPayload := fmt.Sprintf("%s.%s", timestampHeader, string(payload))

	mac := hmac.New(sha256.New, []byte(webhookSecret))
	mac.Write([]byte(expectedPayload))
	computedSignature := hex.EncodeToString(mac.Sum(nil))

	if !hmac.Equal([]byte(computedSignature), []byte(signatureHeader)) {
		return &SignatureVerificationError{Message: "webhook signature verification failed"}
	}

	return nil
}

// ConstructEvent parses and verifies the raw webhook HTTP request payload.
func ConstructEvent(payload []byte, signatureHeader, timestampHeader, webhookSecret string) (*WebhookEvent, error) {
	if err := VerifySignature(payload, signatureHeader, timestampHeader, webhookSecret, 300); err != nil {
		return nil, err
	}

	var event WebhookEvent
	if err := json.Unmarshal(payload, &event); err != nil {
		return nil, fmt.Errorf("bursapay: failed to parse webhook event JSON: %w", err)
	}

	return &event, nil
}
