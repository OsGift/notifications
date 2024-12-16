package sms

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// SMSProvider defines the interface for SMS providers.
type SMSProvider interface {
	SendSMS(phoneNumber, message string) error
}

// NewSMSProvider selects the SMS provider based on configuration.
func NewSMSProvider() SMSProvider {
	switch os.Getenv("SMS_PROVIDER") {
	case "TERMII":
		return &TermiiProvider{}
	case "YOUR_NOTIFY":
		return &YourNotifyProvider{}
	default:
		return &TermiiProvider{} // Default provider
	}
}

// TermiiProvider implements SMSProvider using Termii.
type TermiiProvider struct{}

func (p *TermiiProvider) SendSMS(phoneNumber, message string) error {
	url := os.Getenv("TERMII_URL")
	apiKey := os.Getenv("TERMII_KEY")

	payload := map[string]interface{}{
		"to":      phoneNumber,
		"from":    os.Getenv("TERMII_FROM"),
		"sms":     message,
		"type":    os.Getenv("TERMII_TYPE"),
		"channel": os.Getenv("TERMII_CHANNEL"),
		"api_key": apiKey,
	}

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	response, err := MakeAPICall("POST", url, payload, nil, headers)
	if err != nil {
		return fmt.Errorf("failed to send SMS with Termii: %v", err)
	}

	fmt.Printf("SMS sent successfully with Termii: %+v\n", response)
	return nil
}

// YourNotifyProvider implements SMSProvider using YourNotify.
type YourNotifyProvider struct{}

func (p *YourNotifyProvider) SendSMS(phoneNumber, message string) error {
	url := os.Getenv("YOUR_NOTIFY_URL")
	apiKey := os.Getenv("YOUR_NOTIFY_API_KEY")

	if apiKey == "" || url == "" {
		return fmt.Errorf("YOUR_NOTIFY_API_KEY or YOUR_NOTIFY_URL is not set")
	}

	payload := map[string]interface{}{
		"name":    "Renda",
		"subject": "Notification",
		"from":    "Renda",
		"text":    message,
		"status":  "running",
		"lists":   []string{phoneNumber},
		"channel": "sms",
	}

	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + apiKey,
	}

	response, err := MakeAPICall("POST", url, payload, nil, headers)
	if err != nil {
		return fmt.Errorf("failed to send SMS with YourNotify: %v", err)
	}

	fmt.Printf("SMS sent successfully with YourNotify: %+v\n", response)
	return nil
}

// MakeAPICall performs a generic HTTP request.
func MakeAPICall(method, endpoint string, payload interface{}, queryParams map[string]string, headers map[string]string) (map[string]interface{}, error) {
	var body []byte
	if payload != nil {
		jsonPayload, err := json.Marshal(payload)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal payload: %v", err)
		}
		body = jsonPayload
	}

	// If it's a GET request and there are query params, append them to the URL.
	if method == "GET" && queryParams != nil {
		req, err := http.NewRequest(method, endpoint, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create HTTP request: %v", err)
		}
		for k, v := range queryParams {
			req.URL.Query().Add(k, v)
		}
		endpoint = req.URL.String()
	}

	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %v", err)
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("received non-2xx status code: %d", resp.StatusCode)
	}

	var response map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return response, nil
}
