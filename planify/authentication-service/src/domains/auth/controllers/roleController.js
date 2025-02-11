const RoleService = require('../services/roleService');
const {catchAsync} = require('../../../utils/errorHandler');
const {ApiResponse} = require("../../../utils/apiResponse");

class RoleController {
    createRole = catchAsync(async (req, res) => {
        const {name, description} = req.body;
        const role = await RoleService.createRole(name, description);

        return ApiResponse.created(res, {
            message: 'Role created successfully',
            data: { role }
        });
    });

    getAllRoles = catchAsync(async (req, res) => {
        const roles = await RoleService.getAllRoles(req.query);

        return ApiResponse.success(res, {
            data: { roles }
        });
    });

    getRoleById = catchAsync(async (req, res) => {
        const {id} = req.params;
        const role = await RoleService.getRoleById(id);

        return ApiResponse.success(res, {
            data: { role }
        });
    });

    updateRole = catchAsync(async (req, res) => {
        const {id} = req.params;
        const {name, description} = req.body;
        const role = await RoleService.updateRole(id, {name, description});

        return ApiResponse.success(res, {
            data: { role }
        });
    });

    deleteRole = catchAsync(async (req, res) => {
        const {id} = req.params;
        await RoleService.deleteRole(id);

        return ApiResponse.noContent(res);
    });
}

module.exports = new RoleController();
