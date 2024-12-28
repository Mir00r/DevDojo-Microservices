package services

type VerifyMFARequest struct {
	OTP string `json:"otp" validate:"required,len=6"`
}
