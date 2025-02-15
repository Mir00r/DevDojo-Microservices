/**
 * Data Transfer Objects for Privilege management
 */

class CreatePrivilegeDto {
    constructor(data) {
        this.name = data.name;
        this.description = data.description;
        this.module = data.module;
        this.isActive = data.isActive ?? true;
    }

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

class UpdatePrivilegeDto {
    constructor(data) {
        this.name = data.name;
        this.description = data.description;
        this.module = data.module;
        this.isActive = data.isActive;
    }

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

class PrivilegeQueryDto {
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

class PrivilegeResponseDto {
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
