package dtos

type ValidateTokenRequest struct {
	Token string `json:"token" validate:"required"`
}

type ValidateTokenResponse struct {
	IsValid bool   `json:"isVid"`
	UserID  string `json:"userId,omitempty"` // Returned only if the token is valid
	Expires string `json:"expires,omitempty"`
	Message string `json:"message"`
}
