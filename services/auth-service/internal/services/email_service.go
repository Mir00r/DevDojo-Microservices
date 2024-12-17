package services

import "fmt"

type EmailService struct{}

func NewEmailService() *EmailService {
	return &EmailService{}
}

func (svc *EmailService) SendPasswordResetEmail(email, resetLink string) error {
	// Implement your email sending logic here (e.g., using an SMTP server or a third-party service)
	fmt.Printf("Sending password reset email to %s: %s\n", email, resetLink)
	return nil
}
