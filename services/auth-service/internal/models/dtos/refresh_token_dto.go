package dtos

// RefreshTokenRequest represents the payload for refreshing an access token
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

// RefreshTokenResponse represents the response for refreshing an access token
type RefreshTokenResponse struct {
	AccessToken           string `json:"access_token"`
	RefreshToken          string `json:"refresh_token"`
	ExpiresIn             int64  `json:"expires_in"`               // Time in seconds until the new access token expires
	RefreshTokenExpiresIn int64  `json:"refresh_token_expires_in"` // Time in seconds until the token expires
}
