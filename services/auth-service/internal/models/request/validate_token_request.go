package services

type ValidateTokenRequest struct {
	Token string `json:"token" validate:"required"`
}
