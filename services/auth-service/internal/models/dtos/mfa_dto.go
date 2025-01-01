package dtos

type VerifyMFARequest struct {
	OTP string `json:"otp" validate:"required,len=6"`
}

type VerifyMFAResponse struct {
	OTP string `json:"otp" validate:"required,len=6"`
}

type EnableMFAResponse struct {
	Message string `json:"message"`
	OTP     string `json:"otp"`
}
