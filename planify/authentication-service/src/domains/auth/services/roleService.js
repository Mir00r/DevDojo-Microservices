const {Role, User} = require('../../../../models');
const {AppError} = require('../../../utils/errorHandler');

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

    async getAllRoles() {
        const roles = await Role.findAll({
            include: [{
                model: User,
                as: 'users',
                attributes: ['id', 'name', 'email'],
            }]
        });

        return roles;
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
