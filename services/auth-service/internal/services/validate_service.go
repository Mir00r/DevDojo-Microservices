package services

import "github.com/go-playground/validator/v10"

// Validator instance (singleton)
var validate *validator.Validate

// Initialize the validator
func init() {
	validate = validator.New()
}

// ValidateRequest validates a struct based on validation tags
func ValidateRequest(req interface{}) error {
	return validate.Struct(req)
}
