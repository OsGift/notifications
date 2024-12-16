package push

import (
	"fmt"
	"os"
)

type PushProvider interface {
	SendPush(deviceToken, title, body string, data map[string]interface{}) error
}

func NewPushProvider() PushProvider {
	switch os.Getenv("PUSH_PROVIDER") {
	case "FIREBASE":
		return &FirebaseProvider{}
	default:
		return &FirebaseProvider{} // Default provider
	}
}

type FirebaseProvider struct{}

func (p *FirebaseProvider) SendPush(deviceToken, title, body string, data map[string]interface{}) error {
	// Example: Use Firebase Cloud Messaging
	fmt.Println("Sending push notification via Firebase...")
	return nil
}
