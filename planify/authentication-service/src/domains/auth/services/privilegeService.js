/**
 * Service class for managing privilege operations
 * Handles business logic for creating, reading, updating, and deleting privileges
 */
const {PrivilegeResponseDto} = require('../dtos/privilege.dto');
const PrivilegeRepository = require('../repositories/privilege.repository');

class PrivilegeService {
    /**
     * Creates a new privilege
     * @param {CreatePrivilegeDto} createDto - The DTO containing privilege creation data
     * @returns {Promise<Object>} The created privilege data transformed through PrivilegeResponseDto
     */
    async createPrivilege(createDto) {
        const privilege = await PrivilegeRepository.create(createDto);
        return PrivilegeResponseDto.from(privilege);
    }

    /**
     * Retrieves all privileges with pagination and filtering
     * @param {PrivilegeQueryDto} queryDto - The DTO containing query parameters for filtering and pagination
     * @returns {Promise<Object>} Paginated list of privileges transformed through PrivilegeResponseDto
     */
    async getAllPrivileges(queryDto) {
        const privileges = await PrivilegeRepository.findAll(queryDto);
        return PrivilegeResponseDto.fromPaginated(privileges);
    }

    /**
     * Retrieves a privilege by its ID
     * @param {string} id - The unique identifier of the privilege
     * @returns {Promise<Object>} The privilege data transformed through PrivilegeResponseDto
     */
    async getPrivilegeById(id) {
        const privilege = await PrivilegeRepository.findById(id);
        return PrivilegeResponseDto.from(privilege);
    }

    /**
     * Updates an existing privilege
     * @param {string} id - The unique identifier of the privilege to update
     * @param {UpdatePrivilegeDto} updateDto - The DTO containing privilege update data
     * @returns {Promise<Object>} The updated privilege data transformed through PrivilegeResponseDto
     */
    async updatePrivilege(id, updateDto) {
        const privilege = await PrivilegeRepository.update(id, updateDto);
        return PrivilegeResponseDto.from(privilege);
    }

    /**
     * Deletes a privilege by its ID
     * @param {string} id - The unique identifier of the privilege to delete
     * @returns {Promise<void>}
     */
    async deletePrivilege(id) {
        await PrivilegeRepository.delete(id);
    }
}

module.exports = new PrivilegeService();
