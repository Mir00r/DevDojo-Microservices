const PrivilegeService = require('../services/privilegeService');
const {catchAsync} = require('../../../utils/errorHandler');
const {ApiResponse} = require('../../../utils/apiResponse');
const {CreatePrivilegeDto, UpdatePrivilegeDto, PrivilegeQueryDto} = require('../dtos/privilege.dto');

class PrivilegeController {
    createPrivilege = catchAsync(async (req, res) => {
        const createDto = CreatePrivilegeDto.from(req.body);
        const privilege = await PrivilegeService.createPrivilege(createDto);

        return ApiResponse.created(res, {
            message: 'Privilege created successfully',
            data: { privilege }
        });
    });

    getAllPrivileges = catchAsync(async (req, res) => {
        const queryDto = new PrivilegeQueryDto(req.query);
        const privileges = await PrivilegeService.getAllPrivileges(queryDto);

        return ApiResponse.success(res, {
            data: { privileges }
        });
    });

    getPrivilegeById = catchAsync(async (req, res) => {
        const {id} = req.params;
        const privilege = await PrivilegeService.getPrivilegeById(id);

        return ApiResponse.success(res, {
            data: { privilege }
        });
    });

    updatePrivilege = catchAsync(async (req, res) => {
        const {id} = req.params;
        const updateDto = UpdatePrivilegeDto.from(req.body);
        const privilege = await PrivilegeService.updatePrivilege(id, updateDto);

        return ApiResponse.success(res, {
            data: { privilege }
        });
    });

    deletePrivilege = catchAsync(async (req, res) => {
        const {id} = req.params;
        await PrivilegeService.deletePrivilege(id);

        return ApiResponse.noContent(res);
    });
}

module.exports = new PrivilegeController();