const PublicAuthService = require('../services/publicAuthService');
const {catchAsync, AppError} = require("../../../utils/errorHandler");
const {ApiResponse} = require("../../../utils/apiResponse");

class PublicAuthController {

    register = catchAsync(async (req, res) => {
        const {name, email, password, roleId} = req.body;
        const {user, token} = await PublicAuthService.registerUser(name, email, password, roleId);

        return ApiResponse.created(res, {
            message: 'User registered successfully',
            data: {user, token}
        });
    });

    login = catchAsync(async (req, res) => {
        const {email, password} = req.body;
        const {user, accessToken, refreshToken} = await PublicAuthService.loginUser(email, password);

        return ApiResponse.success(res, {
            data: {user, accessToken, refreshToken}
        });
    });

    forgotPassword = catchAsync(async (req, res) => {
        const {email} = req.body;
        await PublicAuthService.forgotPassword(email);

        return ApiResponse.success(res, {
            message: 'Password reset instructions sent to email'
        });
    });

    resetPassword = catchAsync(async (req, res) => {
        const {token, password} = req.body;
        await PublicAuthService.resetPassword(token, password);

        return ApiResponse.success(res, {
            message: 'Password reset successful'
        });
    });

    verifyEmail = catchAsync(async (req, res) => {
        const {token} = req.body;

        // Input validation
        if (!token || typeof token !== 'string') {
            throw new AppError('Valid verification token is required', 400);
        }

        // Token length/format validation
        if (token.length < 10) { // Minimum token length check
            throw new AppError('Invalid token format', 400);
        }

        const result = await PublicAuthService.verifyEmail(token);

        return ApiResponse.success(res, {
            message: 'Email verified successfully',
            data: {
                email: result.email,
                verifiedAt: result.verifiedAt
            }
        });
    });
}

module.exports = new PublicAuthController();
