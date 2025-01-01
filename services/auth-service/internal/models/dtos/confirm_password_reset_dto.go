package dtos

type ConfirmPasswordResetRequest struct {
	Token       string `json:"token" validate:"required,token"`
	NewPassword string `json:"newPassword" validate:"required,newPassword"`
}
