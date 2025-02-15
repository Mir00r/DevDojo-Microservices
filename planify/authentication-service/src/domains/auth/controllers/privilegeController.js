/**
 * Controller class for handling privilege-related HTTP requests
 * Manages the creation, retrieval, update, and deletion of privileges
 */
const PrivilegeService = require('../services/privilegeService');
const {catchAsync} = require('../../../utils/errorHandler');
const {ApiResponse} = require('../../../utils/apiResponse');
const {CreatePrivilegeDto, UpdatePrivilegeDto, PrivilegeQueryDto} = require('../dtos/privilege.dto');

class PrivilegeController {
    /**
     * Creates a new privilege
     * @param {Request} req - Express request object containing privilege data in body
     * @param {Response} res - Express response object
     * @returns {Promise<Response>} Created privilege with 201 status
     */
    createPrivilege = catchAsync(async (req, res) => {
        const createDto = CreatePrivilegeDto.from(req.body);
        const privilege = await PrivilegeService.createPrivilege(createDto);

        return ApiResponse.created(res, {
            message: 'Privilege created successfully',
            data: { privilege }
        });
    });

    /**
     * Retrieves all privileges with pagination and filtering
     * @param {Request} req - Express request object containing query parameters
     * @param {Response} res - Express response object
     * @returns {Promise<Response>} List of privileges with pagination metadata
     */
    getAllPrivileges = catchAsync(async (req, res) => {
        const queryDto = new PrivilegeQueryDto(req.query);
        const privileges = await PrivilegeService.getAllPrivileges(queryDto);

        return ApiResponse.success(res, {
            data: { privileges }
        });
    });

    /**
     * Retrieves a specific privilege by ID
     * @param {Request} req - Express request object containing privilege ID in params
     * @param {Response} res - Express response object
     * @returns {Promise<Response>} Requested privilege data
     */
    getPrivilegeById = catchAsync(async (req, res) => {
        const {id} = req.params;
        const privilege = await PrivilegeService.getPrivilegeById(id);

        return ApiResponse.success(res, {
            data: { privilege }
        });
    });

    /**
     * Updates an existing privilege
     * @param {Request} req - Express request object containing privilege ID in params and update data in body
     * @param {Response} res - Express response object
     * @returns {Promise<Response>} Updated privilege data
     */
    updatePrivilege = catchAsync(async (req, res) => {
        const {id} = req.params;
        const updateDto = UpdatePrivilegeDto.from(req.body);
        const privilege = await PrivilegeService.updatePrivilege(id, updateDto);

        return ApiResponse.success(res, {
            data: { privilege }
        });
    });

    /**
     * Deletes a privilege
     * @param {Request} req - Express request object containing privilege ID in params
     * @param {Response} res - Express response object
     * @returns {Promise<Response>} 204 No Content response
     */
    deletePrivilege = catchAsync(async (req, res) => {
        const {id} = req.params;
        await PrivilegeService.deletePrivilege(id);

        return ApiResponse.noContent(res);
    });
}

module.exports = new PrivilegeController();