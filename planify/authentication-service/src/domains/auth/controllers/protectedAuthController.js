const ProtectedAuthService = require('../services/protectedAuthService');
const {catchAsync, AppError} = require("../../../utils/errorHandler");

class ProtectedAuthController {

    // Protected endpoints
    refresh = catchAsync(async (req, res) => {
        const {refreshToken} = req.body;

        if (!refreshToken) {
            throw new AppError('Refresh token is required', 400);
        }

        const tokens = await ProtectedAuthService.refreshToken(refreshToken);

        res.json({
            status: 'success',
            data: tokens
        });
    });

    logout = catchAsync(async (req, res) => {
        const {refreshToken} = req.body;

        if (!refreshToken) {
            throw new AppError('Refresh token is required', 400);
        }

        await ProtectedAuthService.logout(refreshToken);

        res.json({
            status: 'success',
            message: 'Logged out successfully'
        });
    });

    logoutAll = catchAsync(async (req, res) => {
        await ProtectedAuthService.logoutAll(req.user.id);

        res.json({
            status: 'success',
            message: 'Logged out from all devices'
        });
    });

    changePassword = catchAsync(async (req, res) => {
        const {currentPassword, newPassword} = req.body;
        await ProtectedAuthService.changePassword(req.user.id, currentPassword, newPassword);

        res.status(200).json({
            status: 'success',
            message: 'Password changed successfully'
        });
    });

    getCurrentUser = catchAsync(async (req, res) => {
        const user = await ProtectedAuthService.getUserById(req.user.id);

        res.status(200).json({
            status: 'success',
            data: {user}
        });
    });

    updateProfile = catchAsync(async (req, res) => {
        const {name, email} = req.body;
        const user = await ProtectedAuthService.updateProfile(req.user.id, {name, email});

        res.status(200).json({
            status: 'success',
            data: {user}
        });
    });
}

module.exports = new ProtectedAuthController();
