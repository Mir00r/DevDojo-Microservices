const AuthService = require('../services/authService');
const {catchAsync, AppError} = require("../../../utils/errorHandler");

class AuthController {
    // Public endpoints
    register = catchAsync(async (req, res) => {
        const {name, email, password, roleId} = req.body;
        const {user, token} = await AuthService.registerUser(name, email, password, roleId);

        res.status(201).json({
            status: 'success',
            data: {user, token}
        });
    });

    login = catchAsync(async (req, res) => {
        const {email, password} = req.body;
        const {user, token} = await AuthService.loginUser(email, password);

        res.status(200).json({
            status: 'success',
            data: {user, token}
        });
    });

    forgotPassword = catchAsync(async (req, res) => {
        const {email} = req.body;
        await AuthService.forgotPassword(email);

        res.status(200).json({
            status: 'success',
            message: 'Password reset instructions sent to email'
        });
    });

    resetPassword = catchAsync(async (req, res) => {
        const {token, password} = req.body;
        await AuthService.resetPassword(token, password);

        res.status(200).json({
            status: 'success',
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

        const result = await AuthService.verifyEmail(token);

        res.status(200).json({
            status: 'success',
            message: 'Email verified successfully',
            data: {
                email: result.email,
                verifiedAt: result.verifiedAt
            }
        });
    });

    // Protected endpoints
    logout = catchAsync(async (req, res) => {
        await AuthService.logout(req.user.id);

        res.status(200).json({
            status: 'success',
            message: 'Logged out successfully'
        });
    });

    changePassword = catchAsync(async (req, res) => {
        const {currentPassword, newPassword} = req.body;
        await AuthService.changePassword(req.user.id, currentPassword, newPassword);

        res.status(200).json({
            status: 'success',
            message: 'Password changed successfully'
        });
    });

    getCurrentUser = catchAsync(async (req, res) => {
        const user = await AuthService.getUserById(req.user.id);

        res.status(200).json({
            status: 'success',
            data: {user}
        });
    });

    updateProfile = catchAsync(async (req, res) => {
        const {name, email} = req.body;
        const user = await AuthService.updateProfile(req.user.id, {name, email});

        res.status(200).json({
            status: 'success',
            data: {user}
        });
    });

    // Internal endpoints (Admin only)
    getAllUsers = catchAsync(async (req, res) => {
        const users = await AuthService.getAllUsers(req.query);

        res.status(200).json({
            status: 'success',
            data: {users}
        });
    });

    getUserById = catchAsync(async (req, res) => {
        const user = await AuthService.getUserById(req.params.id);

        res.status(200).json({
            status: 'success',
            data: {user}
        });
    });

    updateUserRole = catchAsync(async (req, res) => {
        const {roleId} = req.body;
        const user = await AuthService.updateUserRole(req.params.id, roleId);

        res.status(200).json({
            status: 'success',
            data: {user}
        });
    });

    deleteUser = catchAsync(async (req, res) => {
        await AuthService.deleteUser(req.params.id);

        res.status(200).json({
            status: 'success',
            message: 'User deleted successfully'
        });
    });
}

module.exports = new AuthController();
