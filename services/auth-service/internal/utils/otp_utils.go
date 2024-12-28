package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// GenerateOTP generates a 6-digit numeric OTP
func GenerateOTP() string {
	maxIn := big.NewInt(1000000) // Maximum value for 6-digit OTP (exclusive)
	n, err := rand.Int(rand.Reader, maxIn)
	if err != nil {
		// Handle the error appropriately (log, return a default value, etc.)
		fmt.Println("Error generating OTP:", err)
		return "000000"
	}
	return fmt.Sprintf("%06d", n) // Format the OTP to always have 6 digits
}
