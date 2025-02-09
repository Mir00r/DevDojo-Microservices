package apiclients

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	config "github.com/Mir00r/auth-service/configs"
	"io/ioutil"
	"net/http"
	"time"
)

// WebClientConfig holds configuration for the HTTP client
type WebClientConfig struct {
	BaseURL        string
	Timeout        time.Duration
	Headers        map[string]string
	AuthMiddleware func(req *http.Request)

	//BaseURL  string
	//Username string
	//Password string
	//HTTPClient *http.Client
}

// WebClient is a generic HTTP client structure
type WebClient struct {
	Config WebClientConfig
	Client *http.Client
}

// NewWebClient initializes a new WebClient
func NewWebClient() WebClient {
	return WebClient{
		Config: WebClientConfig{
			BaseURL: config.AppConfig.InternalSecurity.BaseUrl,
			Timeout: 10 * time.Second,
			Headers: map[string]string{
				"Authorization": "Basic " + basicAuth(config.AppConfig.InternalSecurity.UserName, config.AppConfig.InternalSecurity.Password),
			},
			AuthMiddleware: BasicAuthMiddleware(config.AppConfig.InternalSecurity.UserName, config.AppConfig.InternalSecurity.Password),
			//Username: config.AppConfig.InternalSecurity.UserName,
			//Password: config.AppConfig.InternalSecurity.Password,
			//HTTPClient: &http.Client{
			//	Timeout: 10 * time.Second,
			//},
		},
		Client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// GetEndpoint dynamically resolves endpoints
func (wc *WebClient) GetEndpoint(endpoint string) string {
	return wc.Config.BaseURL + endpoint
}

// Send sends a request and decodes the response
func (wc *WebClient) Send(method, endpoint string, body interface{}, response interface{}) error {
	url := wc.GetEndpoint(endpoint)

	// Serialize body if present
	var requestBody []byte
	var err error
	if body != nil {
		requestBody, err = json.Marshal(body)
		if err != nil {
			return err
		}
	}

	// Create HTTP request
	req, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}

	// Set headers
	for key, value := range wc.Config.Headers {
		req.Header.Set(key, value)
	}
	req.Header.Set("Content-Type", "application/json")

	// Apply authentication middlewares
	if wc.Config.AuthMiddleware != nil {
		wc.Config.AuthMiddleware(req)
	}

	// Perform HTTP request
	resp, err := wc.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check for non-200 status codes
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return errors.New("received non-2xx response: " + resp.Status)
	}

	// Parse response body
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Deserialize response
	if response != nil {
		if err := json.Unmarshal(responseBody, response); err != nil {
			return err
		}
	}

	return nil
}

// Example usage of BasicAuthMiddleware
func BasicAuthMiddleware(username, password string) func(req *http.Request) {
	return func(req *http.Request) {
		req.SetBasicAuth(username, password)
	}
}

// basicAuth generates a Basic Auth header value.
func basicAuth(username, password string) string {
	return base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
}
