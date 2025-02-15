const {Role, User} = require('../../../../models');
const {AppError} = require('../../../utils/errorHandler');
const {getPaginationParams, formatPaginatedResponse} = require('../../../utils/paginationUtils');
const roleRepository = require('../repositories/role.repository');
const {Op} = require("sequelize");
const {RoleResponseDto, CreateRoleDto, RoleQueryDto, UpdateRoleDto} = require("../dtos/role.dto");

class RoleService {
    /**
     * Create a new role
     * @param {string} name - Role name
     * @param {string} description - Role description
     * @returns {Promise<Object>} Created role
     */
    async createRole(name, description) {
        const createRoleDto = CreateRoleDto.from({name, description});

        // Check if role exists
        const existingRole = await roleRepository.findByName(createRoleDto.name);
        if (existingRole) {
            throw new AppError('Role with this name already exists', 400);
        }

        const role = await roleRepository.create(createRoleDto);
        return RoleResponseDto.from(role);
    }

    /**
     * Get paginated list of roles
     * @param {Object} query - Query parameters
     * @returns {Promise<Object>} Paginated role list
     */
    async getAllRoles(query) {
        const queryDto = new RoleQueryDto(query);
        const paginatedRoles = await roleRepository.findAllPaginated(queryDto);
        return RoleResponseDto.fromPaginated(paginatedRoles);
    }

    /**
     * Get role by ID
     * @param {number} id - Role ID
     * @returns {Promise<Object>} Role data
     */
    async getRoleById(id) {
        const role = await roleRepository.findById(id, true);
        if (!role) {
            throw new AppError('Role not found', 404);
        }
        return RoleResponseDto.from(role);
    }

    /**
     * Update role
     * @param {number} id - Role ID
     * @param {Object} updateData - Role update data
     * @returns {Promise<Object>} Updated role
     */
    async updateRole(id, updateData) {
        const updateRoleDto = UpdateRoleDto.from(updateData);
        const role = await roleRepository.findById(id);

        if (!role) {
            throw new AppError('Role not found', 404);
        }

        // Prevent system role modification
        if (role.name === 'ADMIN' || role.name === 'USER') {
            throw new AppError('Cannot modify system roles', 403);
        }

        // Check name uniqueness if name is being updated
        if (updateRoleDto.name && updateRoleDto.name !== role.name) {
            const existingRole = await roleRepository.findByName(updateRoleDto.name);
            if (existingRole) {
                throw new AppError('Role with this name already exists', 400);
            }
        }

        const updatedRole = await roleRepository.update(role, updateRoleDto);
        return RoleResponseDto.from(updatedRole);
    }

    /**
     * Delete role
     * @param {number} id - Role ID
     * @returns {Promise<boolean>}
     */
    async deleteRole(id) {
        const role = await roleRepository.findById(id);
        if (!role) {
            throw new AppError('Role not found', 404);
        }

        // Prevent system role deletion
        if (role.name === 'ADMIN' || role.name === 'USER') {
            throw new AppError('Cannot delete system roles', 403);
        }

        // Check for assigned users
        const userCount = await roleRepository.countUsers(id);
        if (userCount > 0) {
            throw new AppError('Cannot delete role that is assigned to users', 400);
        }

        return roleRepository.delete(role);
    }
}

module.exports = new RoleService();
