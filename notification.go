package notification

import (
	"notification/email"
	"notification/push"
	"notification/sms"
)

type NotificationService struct {
	SMSProvider   sms.SMSProvider
	EmailProvider email.EmailProvider
	PushProvider  push.PushProvider
}

func NewNotificationService() *NotificationService {
	return &NotificationService{
		SMSProvider:   sms.NewSMSProvider(),
		EmailProvider: email.NewEmailProvider(),
		PushProvider:  push.NewPushProvider(),
	}
}

func (n *NotificationService) SendSMS(phoneNumber, message string) error {
	return n.SMSProvider.SendSMS(phoneNumber, message)
}

func (n *NotificationService) SendEmail(email, subject, template string, variables map[string]interface{}) error {
	return n.EmailProvider.SendEmail(email, subject, template, variables)
}

func (n *NotificationService) SendPush(deviceToken, title, body string, data map[string]interface{}) error {
	return n.PushProvider.SendPush(deviceToken, title, body, data)
}
