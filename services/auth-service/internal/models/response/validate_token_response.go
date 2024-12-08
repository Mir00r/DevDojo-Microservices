package services

type ValidateTokenResponse struct {
	IsValid bool   `json:"is_valid"`
	UserID  string `json:"user_id,omitempty"` // Returned only if the token is valid
	Message string `json:"message"`
}
