package paystack

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/PaulAjii/go-wallet/pkg/config"
)

type Service struct {
	client  *http.Client
	secret  string
	baseUri string
}

func NewService() *Service {
	return &Service{
		client:  &http.Client{},
		secret:  config.ApplicationConfig.Paystack.PaystackSecret,
		baseUri: "https://api.paystack.co",
	}
}

func (s *Service) InitiateTransaction(email string, amountKobo int64, reference, callbackURI string) (*InitializeResponse, error) {
	payload := map[string]interface{}{
		"email":        email,
		"amount":       amountKobo,
		"reference":    reference,
		"callback_url": callbackURI,
	}

	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", s.baseUri+"/transaction/initialize", bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+s.secret)
	req.Header.Set("Content-Type", "application/json")

	res, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var response InitializeResponse
	if err = json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}

	if !response.Status {
		return nil, fmt.Errorf("paystack error: %w", err)
	}

	return &response, nil
}

func (s *Service) VerifyResponse(reference string) (*VerifyResponse, error) {
	req, _ := http.NewRequest("GET", s.baseUri+"/transaction/verify/"+reference, nil)
	req.Header.Set("Authorization", "Bearer "+s.secret)

	res, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var response VerifyResponse
	if err = json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *Service) VerifyWebhookSignature(body []byte, signature string) bool {
	hash := hmac.New(sha512.New, []byte(s.secret))
	hash.Write(body)
	computed := hex.EncodeToString(hash.Sum(nil))
	return hmac.Equal([]byte(computed), []byte(signature))
}
