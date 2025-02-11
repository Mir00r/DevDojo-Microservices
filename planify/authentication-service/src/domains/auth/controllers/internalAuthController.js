const InternalAuthService = require('../services/internalAuthService');
const {catchAsync, AppError} = require("../../../utils/errorHandler");
const {ApiResponse} = require("../../../utils/apiResponse");

class InternalAuthController {

    // Internal endpoints (Admin only)
    getAllUsers = catchAsync(async (req, res) => {
        const users = await InternalAuthService.getAllUsers(req.query);

        return ApiResponse.success(res, {
            data: { users }
        });
    });

    getUserById = catchAsync(async (req, res) => {
        const user = await InternalAuthService.getUserById(req.params.id);

        return ApiResponse.success(res, {
            data: { user }
        });
    });

    updateUserRole = catchAsync(async (req, res) => {
        const {roleId} = req.body;
        const user = await InternalAuthService.updateUserRole(req.params.id, roleId);

        return ApiResponse.success(res, {
            data: { user }
        });
    });

    deleteUser = catchAsync(async (req, res) => {
        await InternalAuthService.deleteUser(req.params.id);

        return ApiResponse.noContent(res);
    });
}

module.exports = new InternalAuthController();
