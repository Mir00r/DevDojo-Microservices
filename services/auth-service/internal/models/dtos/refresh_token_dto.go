package dtos

// RefreshTokenRequest represents the payload for refreshing an access token
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

// RefreshTokenResponse represents the response for refreshing an access token
type RefreshTokenResponse struct {
	AccessToken           string `json:"accessToken"`
	RefreshToken          string `json:"refreshToken"`
	ExpiresIn             int64  `json:"expiresIn"`             // Time in seconds until the new access token expires
	RefreshTokenExpiresIn int64  `json:"refreshTokenExpiresIn"` // Time in seconds until the token expires
}
