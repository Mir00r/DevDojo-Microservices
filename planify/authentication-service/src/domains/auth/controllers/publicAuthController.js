const PublicAuthService = require('../services/publicAuthService');
const {catchAsync, AppError} = require("../../../utils/errorHandler");
const {ApiResponse} = require("../../../utils/apiResponse");
const {LoginRequestDto, LoginResponseDto} = require('../dtos/login.dto');
const {RegisterDto, AuthResponseDto} = require("../dtos/register.dto");
const {VerifyEmailRequestDto, EmailVerificationResponseDto} = require("../dtos/emal.dto");
const {ResetPasswordRequestDto, ForgotPasswordRequestDto} = require("../dtos/password.dto");

class PublicAuthController {

    /**
     * Register a new user
     * @param {Express.Request} req - Express request object
     * @param {Express.Response} res - Express response object
     * @returns {Promise<void>}
     */
    register = catchAsync(async (req, res) => {
        // Validate and transform input data
        const registerDto = await RegisterDto.validate(req.body);

        // Process registration
        const result = await PublicAuthService.registerUser(registerDto);

        // Transform and send response
        return ApiResponse.created(res, {
            message: 'User registered successfully',
            data: AuthResponseDto.toResponse(result)
        });
    });

    login = catchAsync(async (req, res) => {
        // Validate and transform input
        const loginDto = LoginRequestDto.validate(req.body);
        // Process login
        const result = await PublicAuthService.login(loginDto);
        // Transform and send response
        return ApiResponse.success(res, {
            data: result
        });
    });

    /**
     * Handle forgot password request
     * @param {Express.Request} req - Express request object
     * @param {Express.Response} res - Express response object
     */
    forgotPassword = catchAsync(async (req, res) => {
        // Validate request data using DTO
        const forgotPasswordDto = new ForgotPasswordRequestDto(req.body);

        // Process forgot password request
        await PublicAuthService.forgotPassword(forgotPasswordDto.email);

        return ApiResponse.success(res, {
            message: 'Password reset instructions sent to email'
        });
    });

    /**
     * Handle password reset with token
     * @param {Express.Request} req - Express request object
     * @param {Express.Response} res - Express response object
     */
    resetPassword = catchAsync(async (req, res) => {
        // Validate request data using DTO
        const resetPasswordDto = new ResetPasswordRequestDto(req.body);

        // Process password reset
        await PublicAuthService.resetPassword(
            resetPasswordDto.token,
            resetPasswordDto.password
        );

        return ApiResponse.success(res, {
            message: 'Password reset successful'
        });
    });

    /**
     * Handle email verification
     * @param {Express.Request} req - Express request object
     * @param {Express.Response} res - Express response object
     */
    verifyEmail = catchAsync(async (req, res) => {
        // Validate request data using DTO
        const verifyEmailDto = new VerifyEmailRequestDto(req.body);

        // Process email verification
        const result = await PublicAuthService.verifyEmail(verifyEmailDto.token);

        // Transform response using DTO
        const response = EmailVerificationResponseDto.fromEntity(result);

        return ApiResponse.success(res, {
            message: 'Email verified successfully',
            data: response
        });
    });
}

module.exports = new PublicAuthController();
