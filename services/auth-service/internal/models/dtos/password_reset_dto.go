package dtos

type PasswordResetRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type PasswordResetResponse struct {
	Message string `json:"message"`
}
