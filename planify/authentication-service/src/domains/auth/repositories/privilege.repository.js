const {Op} = require('sequelize');
const BaseRepository = require('../repositories/base.repository');
const {Privilege, Role} = require('../../../../models');

/**
 * Repository for Privilege entity operations
 * @extends BaseRepository
 */
/**
 * Repository class for handling privilege database operations
 * Extends BaseRepository to provide common CRUD operations
 */
class PrivilegeRepository extends BaseRepository {
    constructor() {
        super(Privilege);
    }

    /**
     * Find privilege by name
     * @param {string} name - Privilege name
     * @returns {Promise<Privilege>} Found privilege or null
     * @throws {Error} If database operation fails
     */
    async findByName(name) {
        return this.findOne({
            where: {name: name.toUpperCase()}
        });
    }

    /**
     * Find privilege by ID with optional role inclusion
     * @param {string} id - Privilege ID
     * @param {boolean} includeRoles - Whether to include associated roles
     * @returns {Promise<Privilege>} Found privilege with optional roles or null
     * @throws {Error} If database operation fails
     */
    async findById(id, includeRoles = false) {
        const options = {
            where: {id}
        };

        if (includeRoles) {
            options.include = [{
                model: Role,
                as: 'roles',
                attributes: ['id', 'name']
            }];
        }

        return super.findById(id, options);
    }

    /**
     * Find all privileges with pagination and filters
     * @param {PrivilegeQueryDto} queryDto - Query parameters for filtering and pagination
     * @returns {Promise<Object>} Paginated privilege list with count
     * @throws {Error} If database operation fails
     */
    async findAll(queryDto) {
        const whereClause = {};

        if (queryDto.search) {
            whereClause[Op.or] = [
                {name: {[Op.iLike]: `%${queryDto.search}%`}},
                {module: {[Op.iLike]: `%${queryDto.search}%`}}
            ];
        }

        if (queryDto.module) {
            whereClause.module = queryDto.module.toUpperCase();
        }

        if (queryDto.isActive !== undefined) {
            whereClause.isActive = queryDto.isActive;
        }

        const options = {
            where: whereClause,
            include: [{
                model: Role,
                as: 'roles',
                attributes: ['id', 'name']
            }],
            order: [[queryDto.sortBy, queryDto.sortOrder]],
            limit: queryDto.limit,
            offset: (queryDto.page - 1) * queryDto.limit
        };

        return super.findAndCountAll(options);
    }

    /**
     * Create new privilege
     * @param {CreatePrivilegeDto} createPrivilegeDto - Privilege creation data
     * @param {Object} options - Additional options (transaction, etc.)
     * @returns {Promise<Privilege>} Created privilege instance
     * @throws {Error} If database operation fails
     */
    async create(createPrivilegeDto, options = {}) {
        const privilegeData = {
            name: createPrivilegeDto.name,
            description: createPrivilegeDto.description,
            module: createPrivilegeDto.module,
            isActive: createPrivilegeDto.isActive
        };

        return super.create(privilegeData, options);
    }

    /**
     * Update privilege
     * @param {string} id - Privilege ID
     * @param {UpdatePrivilegeDto} updatePrivilegeDto - Privilege update data
     * @param {Object} options - Additional options (transaction, etc.)
     * @returns {Promise<Privilege>} Updated privilege instance
     * @throws {Error} If database operation fails
     */
    async update(id, updatePrivilegeDto, options = {}) {
        const updateData = {
            ...(updatePrivilegeDto.name && {name: updatePrivilegeDto.name}),
            ...(updatePrivilegeDto.description !== undefined && {description: updatePrivilegeDto.description}),
            ...(updatePrivilegeDto.module && {module: updatePrivilegeDto.module}),
            ...(updatePrivilegeDto.isActive !== undefined && {isActive: updatePrivilegeDto.isActive})
        };

        return super.update(id, updateData, options);
    }

    /**
     * Check if privilege is associated with any roles
     * @param {string} id - Privilege ID
     * @returns {Promise<boolean>} True if privilege has associated roles
     * @throws {Error} If database operation fails
     */
    async hasAssociatedRoles(id) {
        const privilege = await this.findOne({
            where: {id},
            include: [{
                model: Role,
                as: 'roles',
                attributes: ['id']
            }]
        });

        return privilege?.roles?.length > 0;
    }

    /**
     * Get privileges by module
     * @param {string} module - Module name
     * @returns {Promise<Privilege[]>} Array of privileges for the specified module
     * @throws {Error} If database operation fails
     */
    async findByModule(module) {
        return this.findAll({
            where: {
                module: module.toUpperCase(),
                isActive: true
            }
        });
    }

    /**
     * Bulk create privileges
     * @param {CreatePrivilegeDto[]} privileges - Array of privilege data
     * @param {Object} options - Additional options (transaction, etc.)
     * @returns {Promise<Privilege[]>} Array of created privilege instances
     * @throws {Error} If database operation fails
     */
    async bulkCreatePrivileges(privileges, options = {}) {
        const privilegesData = privileges.map(dto => ({
            name: dto.name,
            description: dto.description,
            module: dto.module,
            isActive: dto.isActive ?? true
        }));

        return super.bulkCreate(privilegesData, {
            ...options,
            returning: true
        });
    }
}

module.exports = new PrivilegeRepository();
