package utils

import "log"

func SendOTPEmail(email, otp string) error {
	log.Printf("Sending OTP %s to email %s", otp, email)
	// Placeholder: Integrate an actual email service (e.g., SendGrid, SES)
	return nil
}
