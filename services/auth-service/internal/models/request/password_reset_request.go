package services

type PasswordResetRequest struct {
	Email string `json:"email" validate:"required,email"`
}
