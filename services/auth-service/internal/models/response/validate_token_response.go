package services

type ValidateTokenResponse struct {
	IsValid bool   `json:"is_valid"`
	UserID  string `json:"user_id,omitempty"` // Returned only if the token is valid
	Expires string `json:"expires,omitempty"`
	Message string `json:"message"`
}
