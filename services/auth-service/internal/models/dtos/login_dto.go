package dtos

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	AccessToken           string `json:"accessToken"`
	RefreshToken          string `json:"refreshToken"`
	ExpiresIn             int64  `json:"expiresIn"`             // Time in seconds until the token expires
	RefreshTokenExpiresIn int64  `json:"refreshTokenExpiresIn"` // Time in seconds until the token expires
}

// LoginAPIResponse represents the entire response structure
type LoginAPIResponse struct {
	Error   bool          `json:"error"`
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Data    LoginResponse `json:"data"`
}
