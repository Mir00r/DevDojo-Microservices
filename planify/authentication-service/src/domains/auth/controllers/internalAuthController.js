const InternalAuthService = require('../services/internalAuthService');
const {catchAsync, AppError} = require("../../../utils/errorHandler");

class InternalAuthController {

    // Internal endpoints (Admin only)
    getAllUsers = catchAsync(async (req, res) => {
        const users = await InternalAuthService.getAllUsers(req.query);

        res.status(200).json({
            status: 'success',
            data: {users}
        });
    });

    getUserById = catchAsync(async (req, res) => {
        const user = await InternalAuthService.getUserById(req.params.id);

        res.status(200).json({
            status: 'success',
            data: {user}
        });
    });

    updateUserRole = catchAsync(async (req, res) => {
        const {roleId} = req.body;
        const user = await InternalAuthService.updateUserRole(req.params.id, roleId);

        res.status(200).json({
            status: 'success',
            data: {user}
        });
    });

    deleteUser = catchAsync(async (req, res) => {
        await InternalAuthService.deleteUser(req.params.id);

        res.status(200).json({
            status: 'success',
            message: 'User deleted successfully'
        });
    });
}

module.exports = new InternalAuthController();
