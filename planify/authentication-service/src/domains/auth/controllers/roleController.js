const RoleService = require('../services/roleService');
const {catchAsync} = require('../../../utils/errorHandler');

class RoleController {
    createRole = catchAsync(async (req, res) => {
        const {name, description} = req.body;
        const role = await RoleService.createRole(name, description);

        res.status(201).json({
            status: 'success',
            data: {role}
        });
    });

    getAllRoles = catchAsync(async (req, res) => {
        const roles = await RoleService.getAllRoles();

        res.status(200).json({
            status: 'success',
            data: {roles}
        });
    });

    getRoleById = catchAsync(async (req, res) => {
        const {id} = req.params;
        const role = await RoleService.getRoleById(id);

        res.status(200).json({
            status: 'success',
            data: {role}
        });
    });

    updateRole = catchAsync(async (req, res) => {
        const {id} = req.params;
        const {name, description} = req.body;
        const role = await RoleService.updateRole(id, {name, description});

        res.status(200).json({
            status: 'success',
            data: {role}
        });
    });

    deleteRole = catchAsync(async (req, res) => {
        const {id} = req.params;
        await RoleService.deleteRole(id);

        res.status(204).json({
            status: 'success',
            data: null
        });
    });
}

module.exports = new RoleController();
