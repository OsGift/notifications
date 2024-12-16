# Notification Package 📬

A Go package for handling **SMS**, **Email**, and **Push Notifications** with multiple providers. This package allows you to dynamically choose the notification provider based on environment variables.

## Features
- 📱 **SMS Notifications** (Supports Termii, Twilio, or custom providers)
- 📧 **Email Notifications** (Supports Google Gmail API, AWS SES)
- 🔔 **Push Notifications** (Supports Firebase)

## Folder Structure
```plaintext
.
├── .env                    # Environment variables for providers
├── go.mod                  # Go module file
├── main.go                 # Example usage entry point
├── notification.go         # Core package logic
├── email/
│   └── email.go            # Email-specific logic
├── sms/
│   └── sms.go              # SMS-specific logic
├── push/
│   └── push.go             # Push notification logic
├── provider/
│   ├── config.go           # Config loader for environment variables
│   └── provider.go         # Interfaces for providers
└── templates/              # Email templates
