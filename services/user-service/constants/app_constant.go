package constants

import "errors"

// Error message strings
const (
	ErrMissingAuthHeader              = "Authorization header is missing"
	ErrInvalidAuthHeader              = "Invalid Authorization header"
	ErrInvalidToken                   = "Invalid token"
	ErrUserNotFound                   = "User not found"
	InvalidCredentials                = "Invalid email or password"
	Unauthorized                      = "Unauthorized"
	Forbidden                         = "Forbidden"
	ResourceNotFound                  = "Resource not found"
	InternalServerError               = "Internal server error"
	ErrHashPassword                   = "Failed to hash password"
	ErrGenerateToken                  = "Failed to generate token"
	ErrInvalidCredential              = "Invalid credentials"
	ErrTokenInvalidOrReset            = "Reset token already used or invalid"
	ErrTokenInvalidated               = "Token is already invalidated"
	ErrOTPExpired                     = "OTP has expired"
	ErrOTPNotFound                    = "OTP not found"
	ErrInvalidOTP                     = "Invalid OTP"
	ErrInvalidRqPayload               = "Invalid request payload"
	ErrFailedToConfirmPasswordReset   = "Failed to confirm password reset"
	ErrFailedToInitiatePasswordReset  = "Failed to initiate password reset"
	ErrFailedToFetchProfile           = "Failed to fetch user profile"
	ErrFailedToRegisterUser           = "Failed to register user"
	ErrInvalidOrExpiredRefreshToken   = "Invalid or expired refresh token"
	ErrSaveToken                      = "Failed to save token"
	ErrFailedToFindToken              = "Failed to find refresh token"
	ErrFailedToRetrieveUser           = "Failed to retrieve user"
	ErrFailedToGenerateAccessToken    = "Failed to generate access token"
	ErrFailedToGenerateNewAccessToken = "Failed to generate new refresh token"
	ErrFailedToUpdateToken            = "Failed to update token"
	ErrFailedToGenerateResetToken     = "Failed to generate reset token"
	ErrFailedToSaveResetToken         = "Failed to save reset token"
	ErrFailedToEnableMFA              = "Failed to enable MFA"
	ErrFailedToVerifyMFA              = "Failed to verify MFA"
	ErrFailedToSendOTPEmail           = "Failed to send OTP email"
	ErrFailedToMarkMFA                = "Failed to mark MFA as used"
)

// Error variables for use throughout the project
var (
	ErrInvalidCredentials                = errors.New(InvalidCredentials)
	ErrUnauthorized                      = errors.New(Unauthorized)
	ErrForbidden                         = errors.New(Forbidden)
	ErrNotFound                          = errors.New(ResourceNotFound)
	ErrInternalServer                    = errors.New(InternalServerError)
	ErrMissingAuthHeaderVar              = errors.New(ErrMissingAuthHeader)
	ErrInvalidAuthHeaderVar              = errors.New(ErrInvalidAuthHeader)
	ErrInvalidTokenVar                   = errors.New(ErrInvalidToken)
	ErrUserNotFoundVar                   = errors.New(ErrUserNotFound)
	ErrHashPasswordVar                   = errors.New(ErrHashPassword)
	ErrGenerateTokenVar                  = errors.New(ErrGenerateToken)
	ErrInvalidCredentialVar              = errors.New(ErrInvalidCredential)
	ErrTokenInvalidOrResetVar            = errors.New(ErrTokenInvalidOrReset)
	ErrTokenInvalidatedVar               = errors.New(ErrTokenInvalidated)
	ErrInvalidOTPVar                     = errors.New(ErrInvalidOTP)
	ErrOTPExpiredVar                     = errors.New(ErrOTPExpired)
	ErrFailedToConfirmPasswordResetVar   = errors.New(ErrFailedToConfirmPasswordReset)
	ErrFailedToInitiatePasswordResetVar  = errors.New(ErrFailedToInitiatePasswordReset)
	ErrFailedToFetchProfileVar           = errors.New(ErrFailedToFetchProfile)
	ErrOTPNotFoundVar                    = errors.New(ErrOTPNotFound)
	ErrInvalidOrExpiredRefreshTokenVar   = errors.New(ErrInvalidOrExpiredRefreshToken)
	ErrSaveTokenVar                      = errors.New(ErrSaveToken)
	ErrFailedToRegisterUserVar           = errors.New(ErrFailedToRegisterUser)
	ErrFailedToFindTokenVar              = errors.New(ErrFailedToFindToken)
	ErrFailedToGenerateAccessTokenVar    = errors.New(ErrFailedToGenerateAccessToken)
	ErrFailedToGenerateNewAccessTokenVar = errors.New(ErrFailedToGenerateNewAccessToken)
	ErrFailedToRetrieveUserVar           = errors.New(ErrFailedToRetrieveUser)
	ErrFailedToUpdateTokenVar            = errors.New(ErrFailedToUpdateToken)
	ErrFailedToGenerateResetTokenVar     = errors.New(ErrFailedToGenerateResetToken)
	ErrFailedToSaveResetTokenVar         = errors.New(ErrFailedToSaveResetToken)
	ErrFailedToEnableMFAVar              = errors.New(ErrFailedToEnableMFA)
	ErrFailedToVerifyMFAVar              = errors.New(ErrFailedToVerifyMFA)
	ErrFailedToSendOTPEmailVar           = errors.New(ErrFailedToSendOTPEmail)
	ErrFailedToMarkMFAVar                = errors.New(ErrFailedToMarkMFA)
)

// Success messages
const (
	MsgLogoutSuccessful             = "Logout successful"
	MsgUserRegSuccessful            = "User registration successful"
	MFAVerifySuccessful             = "MFA verification successful"
	PasswordResetLinkSentSuccessful = "Password reset link sent successfully"
)

// Api Header
const (
	Authorization = "Authorization"
	Bearer        = "Bearer "
)
