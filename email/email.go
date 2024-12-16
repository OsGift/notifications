package email

import (
	"fmt"
	"os"
)

type EmailProvider interface {
	SendEmail(email, subject, template string, variables map[string]interface{}) error
}

func NewEmailProvider() EmailProvider {
	switch os.Getenv("EMAIL_PROVIDER") {
	case "GCP":
		return &GCPProvider{}
	case "AWS":
		return &AWSSESProvider{}
	default:
		return &GCPProvider{} // Default provider
	}
}

type GCPProvider struct{}

func (p *GCPProvider) SendEmail(email, subject, template string, variables map[string]interface{}) error {
	// Example: Use Gmail API (setup similar to your example code)
	fmt.Println("Sending email via GCP...")
	return nil
}

type AWSSESProvider struct{}

func (p *AWSSESProvider) SendEmail(email, subject, template string, variables map[string]interface{}) error {
	// Example: Use AWS SES API
	fmt.Println("Sending email via AWS SES...")
	return nil
}
