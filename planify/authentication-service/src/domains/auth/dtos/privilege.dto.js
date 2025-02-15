/**
 * Data Transfer Objects for Privilege management
 * Provides validation and transformation of privilege-related data
 */

/**
 * DTO for creating a new privilege
 * Validates and transforms input data for privilege creation
 */
class CreatePrivilegeDto {
    /**
     * Constructs a new CreatePrivilegeDto instance
     * @param {Object} data - Raw privilege data
     * @param {string} data.name - Name of the privilege
     * @param {string} [data.description] - Description of the privilege
     * @param {string} data.module - Module the privilege belongs to
     * @param {boolean} [data.isActive=true] - Active status of the privilege
     */
    constructor(data) {
        this.name = data.name;
        this.description = data.description;
        this.module = data.module;
        this.isActive = data.isActive ?? true;
    }

    /**
     * Creates and validates a new CreatePrivilegeDto instance
     * @param {Object} data - Raw privilege data
     * @returns {CreatePrivilegeDto} Validated privilege DTO
     * @throws {Error} If required fields are missing
     */
    static from(data) {
        if (!data.name || !data.module) {
            throw new Error('Privilege name and module are required');
        }

        return new CreatePrivilegeDto({
            name: data.name.trim().toUpperCase(),
            description: data.description?.trim(),
            module: data.module.trim().toUpperCase(),
            isActive: data.isActive
        });
    }
}

/**
 * DTO for updating an existing privilege
 * Validates and transforms input data for privilege updates
 */
class UpdatePrivilegeDto {
    /**
     * Constructs a new UpdatePrivilegeDto instance
     * @param {Object} data - Raw update data
     * @param {string} [data.name] - Updated name
     * @param {string} [data.description] - Updated description
     * @param {string} [data.module] - Updated module
     * @param {boolean} [data.isActive] - Updated active status
     */
    constructor(data) {
        this.name = data.name;
        this.description = data.description;
        this.module = data.module;
        this.isActive = data.isActive;
    }

    /**
     * Creates and validates a new UpdatePrivilegeDto instance
     * @param {Object} data - Raw update data
     * @returns {UpdatePrivilegeDto} Validated update DTO
     * @throws {Error} If no valid update fields are provided
     */
    static from(data) {
        const dto = new UpdatePrivilegeDto({
            name: data.name?.trim()?.toUpperCase(),
            description: data.description?.trim(),
            module: data.module?.trim()?.toUpperCase(),
            isActive: data.isActive
        });

        if (!dto.name && !dto.description && !dto.module && dto.isActive === undefined) {
            throw new Error('At least one field must be provided for update');
        }

        return dto;
    }
}

/**
 * DTO for privilege queries
 * Handles pagination, sorting, and filtering parameters
 */
class PrivilegeQueryDto {
    /**
     * Constructs a new PrivilegeQueryDto instance
     * @param {Object} data - Query parameters
     * @param {string} [data.search] - Search term for name/module
     * @param {string} [data.module] - Filter by module
     * @param {string} [data.isActive] - Filter by active status
     * @param {number} [data.page=1] - Page number
     * @param {number} [data.limit=10] - Items per page
     * @param {string} [data.sortBy='createdAt'] - Sort field
     * @param {string} [data.sortOrder='DESC'] - Sort direction
     */
    constructor(data) {
        this.search = data.search;
        this.module = data.module;
        this.isActive = data.isActive !== undefined ?
            data.isActive === 'true' : undefined;
        this.page = parseInt(data.page) || 1;
        this.limit = parseInt(data.limit) || 10;
        this.sortBy = data.sortBy || 'createdAt';
        this.sortOrder = data.sortOrder?.toUpperCase() === 'ASC' ? 'ASC' : 'DESC';
    }
}

/**
 * DTO for privilege responses
 * Transforms privilege data for API responses
 */
class PrivilegeResponseDto {
    /**
     * Transforms a privilege instance to response format
     * @param {Object} privilege - Privilege instance
     * @returns {Object} Formatted privilege response
     */
    static from(privilege) {
        return {
            id: privilege.id,
            name: privilege.name,
            description: privilege.description,
            module: privilege.module,
            isActive: privilege.isActive,
            roles: privilege.roles?.map(role => ({
                id: role.id,
                name: role.name
            })),
            createdAt: privilege.createdAt,
            updatedAt: privilege.updatedAt
        };
    }

    /**
     * Transforms paginated privilege results to response format
     * @param {Object} paginatedPrivileges - Paginated privilege results
     * @returns {Object} Formatted paginated response
     */
    static fromPaginated(paginatedPrivileges) {
        return {
            items: paginatedPrivileges.rows.map(privilege =>
                PrivilegeResponseDto.from(privilege)
            ),
            meta: {
                totalItems: paginatedPrivileges.count,
                itemsPerPage: paginatedPrivileges.limit,
                totalPages: Math.ceil(paginatedPrivileges.count / paginatedPrivileges.limit),
                currentPage: paginatedPrivileges.page
            }
        };
    }
}

module.exports = {
    CreatePrivilegeDto,
    UpdatePrivilegeDto,
    PrivilegeQueryDto,
    PrivilegeResponseDto
};
