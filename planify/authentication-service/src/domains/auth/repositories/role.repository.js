const {Op} = require('sequelize');
const BaseRepository = require('./base.repository');
const {Role, User} = require('../../../../models');

/**
 * Repository for Role entity operations
 * @extends BaseRepository
 */
class RoleRepository extends BaseRepository {
    constructor() {
        super(Role);
    }

    /**
     * Find role by name
     * @param {string} name - Role name
     * @returns {Promise<Role>}
     */
    async findByName(name) {
        return this.findOne({
            where: {name: name.toUpperCase()}
        });
    }

    /**
     * Find role by ID with optional user inclusion
     * @param {number} id - Role ID
     * @param {boolean} includeUsers - Whether to include associated users
     * @returns {Promise<Role>}
     */
    async findById(id, includeUsers = false) {
        const options = {};

        if (includeUsers) {
            options.include = [{
                model: User,
                as: 'users',
                attributes: ['id', 'name', 'email']
            }];
        }

        return super.findById(id, options);
    }

    /**
     * Find roles with pagination and filters
     * @param {RoleQueryDto} queryDto - Query parameters
     * @returns {Promise<Object>} Paginated role list
     */
    async findAllPaginated(queryDto) {
        const options = {
            where: {},
            include: [{
                model: User,
                as: 'users',
                attributes: ['id', 'name', 'email'],
                separate: true,
                limit: 5
            }],
            order: [[queryDto.sortBy, queryDto.sortOrder]],
            limit: queryDto.limit,
            offset: (queryDto.page - 1) * queryDto.limit
        };

        if (queryDto.search) {
            options.where.name = {
                [Op.iLike]: `%${queryDto.search}%`
            };
        }

        const roles = await this.findAndCountAll(options);

        // Add user count for each role
        const rolesWithCount = await Promise.all(
            roles.rows.map(async (role) => {
                const userCount = await User.count({
                    where: { roleId: role.id }
                });
                const roleJson = role.toJSON();
                roleJson.userCount = userCount;
                return roleJson;
            })
        );

        return {
            count: roles.count,
            rows: rolesWithCount,
            limit: queryDto.limit,
            page: queryDto.page
        };
    }

    /**
     * Count users with specific role
     * @param {number} roleId - Role ID
     * @returns {Promise<number>}
     */
    async countUsers(roleId) {
        return this.count({
            where: {roleId}
        });
    }

    /**
     * Find roles by multiple IDs
     * @param {number[]} ids - Array of role IDs
     * @returns {Promise<Role[]>}
     */
    async findByIds(ids) {
        return this.findAll({
            where: {
                id: {
                    [Op.in]: ids
                }
            }
        });
    }

    /**
     * Find roles by partial name match
     * @param {string} namePattern - Name pattern to match
     * @returns {Promise<Role[]>}
     */
    async findByNamePattern(namePattern) {
        return this.findAll({
            where: {
                name: {
                    [Op.iLike]: `%${namePattern}%`
                }
            }
        });
    }
}

module.exports = new RoleRepository();
