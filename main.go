package main

import (
	"notification"
	"notification/provider"
)

func main() {
	provider.LoadEnv()

	service := notification.NewNotificationService()

	// Send SMS
	err := service.SendSMS("2347052222152", "Hello from Go Notification!")
	if err != nil {
		panic(err)
	}

	// Send Email
	err = service.SendEmail("gift@renda.co", "Test Email", "template.html", map[string]interface{}{
		"Name": "John Doe",
	})
	if err != nil {
		panic(err)
	}

	// Send Push Notification
	err = service.SendPush("device_token", "Welcome!", "Thanks for joining us.", map[string]interface{}{
		"key": "value",
	})
	if err != nil {
		panic(err)
	}
}
