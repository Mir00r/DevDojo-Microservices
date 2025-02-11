const {Role, User} = require('../../../../models');
const {AppError} = require('../../../utils/errorHandler');
const {getPaginationParams, formatPaginatedResponse} = require('../../../utils/paginationUtils');
const {Op} = require("sequelize");

class RoleService {
    async createRole(name, description) {
        // Check if role already exists
        const existingRole = await Role.findOne({where: {name}});
        if (existingRole) {
            throw new AppError('Role with this name already exists', 400);
        }

        return Role.create({
            name: name.toUpperCase(),
            description
        });
    }

    async getAllRoles(query) {
        const {limit, offset, page} = getPaginationParams(query);

        // Build filter conditions
        const whereClause = {};

        // Name search
        if (query.search) {
            whereClause.name = {
                [Op.iLike]: `%${query.search}%`
            };
        }

        // Get roles with pagination and filters
        const roles = await Role.findAndCountAll({
            where: whereClause,
            include: [{
                model: User,
                as: 'users',
                attributes: ['id', 'name', 'email'],
                separate: true, // Perform separate query for better performance
                limit: 5 // Limit the number of users shown per role
            }],
            order: [
                [query.sortBy || 'createdAt', query.sortOrder || 'DESC']
            ],
            limit,
            offset
        });

        // Add user count for each role
        const rolesWithCount = {
            count: roles.count,
            rows: await Promise.all(roles.rows.map(async (role) => {
                const userCount = await User.count({
                    where: {roleId: role.id}
                });
                const roleJson = role.toJSON();
                roleJson.userCount = userCount;
                return roleJson;
            }))
        };

        return formatPaginatedResponse(rolesWithCount, page, limit);
    }

    async getRoleById(id) {
        const role = await Role.findByPk(id, {
            include: [{
                model: User,
                as: 'users',
                attributes: ['id', 'name', 'email'],
            }]
        });

        if (!role) {
            throw new AppError('Role not found', 404);
        }

        return role;
    }

    async updateRole(id, updateData) {
        const role = await Role.findByPk(id);

        if (!role) {
            throw new AppError('Role not found', 404);
        }

        // Prevent updating if it's a system role (optional)
        if (role.name === 'ADMIN' || role.name === 'USER') {
            throw new AppError('Cannot modify system roles', 403);
        }

        // If name is being updated, check for duplicates
        if (updateData.name && updateData.name !== role.name) {
            const existingRole = await Role.findOne({
                where: {name: updateData.name.toUpperCase()}
            });

            if (existingRole) {
                throw new AppError('Role with this name already exists', 400);
            }

            updateData.name = updateData.name.toUpperCase();
        }

        await role.update(updateData);
        return role;
    }

    async deleteRole(id) {
        const role = await Role.findByPk(id);

        if (!role) {
            throw new AppError('Role not found', 404);
        }

        // Prevent deleting if it's a system role
        if (role.name === 'ADMIN' || role.name === 'USER') {
            throw new AppError('Cannot delete system roles', 403);
        }

        // Check if role is assigned to any users
        const usersWithRole = await User.count({where: {roleId: id}});
        if (usersWithRole > 0) {
            throw new AppError('Cannot delete role that is assigned to users', 400);
        }

        await role.destroy();
        return true;
    }
}

module.exports = new RoleService();
