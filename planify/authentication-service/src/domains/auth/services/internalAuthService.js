const bcrypt = require('bcryptjs');
const jwt = require('jsonwebtoken');
const {sequelize} = require('../../../configs/database');
const {AppError} = require("../../../utils/errorHandler");
const jwtUtils = require('../../../utils/jwtUtils');
const {User, Role} = require('../../../../models');


class InternalAuthService {

    async getUserById(userId) {
        // Find user with their role information
        const user = await User.findOne({
            where: { id: userId },
            include: [{
                model: Role,
                as: 'role',
                attributes: ['name', 'description']
            }],
            attributes: [
                'id',
                'name',
                'email',
                'lastLogin',
                'isActive',
                'createdAt',
                'updatedAt'
            ] // Excluding sensitive fields like password
        });

        if (!user) {
            throw new AppError('User not found', 404);
        }

        // Return sanitized user data
        return jwtUtils.sanitizeUser(user);
    }

    async getAllUsers({page = 1, limit = 10, search = ''}) {
        const offset = (page - 1) * limit;
        const where = search ? {
            [Op.or]: [
                {name: {[Op.iLike]: `%${search}%`}},
                {email: {[Op.iLike]: `%${search}%`}}
            ]
        } : {};

        const users = await User.findAndCountAll({
            where,
            include: [{
                model: Role,
                as: 'role',
                attributes: ['name']
            }],
            limit,
            offset,
            order: [['createdAt', 'DESC']]
        });

        return {
            users: users.rows.map(jwtUtils.sanitizeUser),
            total: users.count,
            pages: Math.ceil(users.count / limit)
        };
    }

    async updateUserRole(userId, roleId) {
        const user = await User.findByPk(userId);

        if (!user) {
            throw new AppError('User not found', 404);
        }

        const role = await Role.findByPk(roleId);
        if (!role) {
            throw new AppError('Role not found', 404);
        }

        await user.update({roleId});
        return jwtUtils.sanitizeUser(user);
    }

    async deleteUser(userId) {
        const user = await User.findByPk(userId);

        if (!user) {
            throw new AppError('User not found', 404);
        }

        await user.destroy(); // Soft delete if paranoid is true
    }
}

module.exports = new InternalAuthService();
