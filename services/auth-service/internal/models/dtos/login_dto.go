package dtos

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	AccessToken           string `json:"access_token"`
	RefreshToken          string `json:"refresh_token"`
	ExpiresIn             int64  `json:"expires_in"`               // Time in seconds until the token expires
	RefreshTokenExpiresIn int64  `json:"refresh_token_expires_in"` // Time in seconds until the token expires
}

// LoginAPIResponse represents the entire response structure
type LoginAPIResponse struct {
	Error   bool          `json:"error"`
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Data    LoginResponse `json:"data"`
}
