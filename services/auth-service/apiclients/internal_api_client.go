package apiclients

import (
	config "github.com/Mir00r/auth-service/configs"
	"net/http"
	"time"
)

// InternalServiceClient defines a client for interacting with the user-service
type InternalServiceClient struct {
	BaseURL    string
	Username   string
	Password   string
	HTTPClient *http.Client
}

// NewUserServiceClient creates a new client for the user-service
func NewUserServiceClient() *InternalServiceClient {
	return &InternalServiceClient{
		BaseURL:  config.AppConfig.InternalSecurity.BaseUrl,
		Username: config.AppConfig.InternalSecurity.UserName,
		Password: config.AppConfig.InternalSecurity.Password,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}
