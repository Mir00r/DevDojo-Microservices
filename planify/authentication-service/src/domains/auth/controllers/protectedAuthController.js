const ProtectedAuthService = require('../services/protectedAuthService');
const {catchAsync, AppError} = require("../../../utils/errorHandler");
const {ApiResponse} = require("../../../utils/apiResponse");

class ProtectedAuthController {

    // Protected endpoints
    refresh = catchAsync(async (req, res) => {
        const {refreshToken} = req.body;

        if (!refreshToken) {
            throw new AppError('Refresh token is required', 400);
        }

        const tokens = await ProtectedAuthService.refreshToken(refreshToken);

        return ApiResponse.success(res, {
            data: {tokens}
        });
    });

    logout = catchAsync(async (req, res) => {
        const {refreshToken} = req.body;

        if (!refreshToken) {
            throw new AppError('Refresh token is required', 400);
        }

        await ProtectedAuthService.logout(refreshToken);

        return ApiResponse.success(res, {
            message: 'Logged out successfully',
            data: {res}
        });
    });

    logoutAll = catchAsync(async (req, res) => {
        await ProtectedAuthService.logoutAll(req.user.id);

        return ApiResponse.success(res, {
            message: 'Logged out from all devices',
            data: {res}
        });
    });

    changePassword = catchAsync(async (req, res) => {
        const {currentPassword, newPassword} = req.body;
        await ProtectedAuthService.changePassword(req.user.id, currentPassword, newPassword);

        return ApiResponse.success(res, {
            message: 'Password changed successfully',
            data: {res}
        });
    });

    getCurrentUser = catchAsync(async (req, res) => {
        const user = await ProtectedAuthService.getUserById(req.user.id);

        return ApiResponse.success(res, {
            data: {user}
        });
    });

    updateProfile = catchAsync(async (req, res) => {
        const {name, email} = req.body;
        const user = await ProtectedAuthService.updateProfile(req.user.id, {name, email});

        return ApiResponse.success(res, {
            data: {user}
        });
    });
}

module.exports = new ProtectedAuthController();
