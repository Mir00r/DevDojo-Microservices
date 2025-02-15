/**
 * Data Transfer Objects for Role management
 */

/**
 * DTO for creating a new role
 */
class CreateRoleDto {
    constructor(data) {
        this.name = data.name;
        this.description = data.description;
    }

    /**
     * Validate and sanitize input data
     * @param {Object} data - Raw input data
     * @returns {CreateRoleDto}
     */
    static from(data) {
        if (!data.name) {
            throw new Error('Role name is required');
        }
        return new CreateRoleDto({
            name: data.name.trim(),
            description: data.description?.trim()
        });
    }
}

/**
 * DTO for updating a role
 */
class UpdateRoleDto {
    constructor(data) {
        this.name = data.name;
        this.description = data.description;
    }

    /**
     * Validate and sanitize update data
     * @param {Object} data - Raw update data
     * @returns {UpdateRoleDto}
     */
    static from(data) {
        const dto = new UpdateRoleDto({
            name: data.name?.trim(),
            description: data.description?.trim()
        });

        // Ensure at least one field is being updated
        if (!dto.name && !dto.description) {
            throw new Error('At least one field must be provided for update');
        }

        return dto;
    }
}

/**
 * DTO for role list query parameters
 */
class RoleQueryDto {
    constructor(data) {
        this.search = data.search;
        this.page = parseInt(data.page) || 1;
        this.limit = parseInt(data.limit) || 10;
        this.sortBy = data.sortBy || 'createdAt';
        this.sortOrder = data.sortOrder?.toUpperCase() === 'ASC' ? 'ASC' : 'DESC';
    }
}

/**
 * DTO for role response
 */
class RoleResponseDto {
    /**
     * Transform role entity to DTO
     * @param {Object} role - Role entity
     * @returns {Object} Transformed role data
     */
    static from(role) {
        return {
            id: role.id,
            name: role.name,
            description: role.description,
            userCount: role.userCount,
            users: role.users?.map(user => ({
                id: user.id,
                name: user.name,
                email: user.email
            })),
            createdAt: role.createdAt,
            updatedAt: role.updatedAt
        };
    }

    /**
     * Transform paginated role list to DTO
     * @param {Object} paginatedRoles - Paginated role data
     * @returns {Object} Transformed paginated response
     */
    static fromPaginated(paginatedRoles) {
        return {
            items: paginatedRoles.rows.map(role => RoleResponseDto.from(role)),
            meta: {
                totalItems: paginatedRoles.count,
                itemsPerPage: paginatedRoles.limit,
                totalPages: Math.ceil(paginatedRoles.count / paginatedRoles.limit),
                currentPage: paginatedRoles.page
            }
        };
    }
}

module.exports = {
    CreateRoleDto,
    UpdateRoleDto,
    RoleQueryDto,
    RoleResponseDto
};
