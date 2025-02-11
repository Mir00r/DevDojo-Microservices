const bcrypt = require('bcryptjs');
const jwt = require('jsonwebtoken');
const {sequelize} = require('../../../configs/database');
const {AppError} = require("../../../utils/errorHandler");
const jwtUtils = require('../../../utils/jwtUtils');
const {getPaginationParams, formatPaginatedResponse} = require('../../../utils/paginationUtils');
const {User, Role} = require('../../../../models');
const {Op} = require("sequelize");


class InternalAuthService {

    async getUserById(userId) {
        // Find user with their role information
        const user = await User.findOne({
            where: {id: userId},
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

    async getAllUsers(query) {
        const {limit, offset, page} = getPaginationParams(query);

        // Build filter conditions
        const whereClause = {};

        // Name search
        if (query.search) {
            whereClause[Op.or] = [
                {name: {[Op.iLike]: `%${query.search}%`}},
                {email: {[Op.iLike]: `%${query.search}%`}}
            ];
        }

        // Role filter
        if (query.roleId) {
            whereClause.roleId = query.roleId;
        }

        // Active status filter
        if (query.isActive !== undefined) {
            whereClause.isActive = query.isActive === 'true';
        }

        // Date range filter
        if (query.startDate && query.endDate) {
            whereClause.createdAt = {
                [Op.between]: [new Date(query.startDate), new Date(query.endDate)]
            };
        }

        // Get users with pagination and filters
        const users = await User.findAndCountAll({
            where: whereClause,
            include: [{
                model: Role,
                as: 'role',
                attributes: ['id', 'name']
            }],
            attributes: {
                exclude: ['password'] // Exclude sensitive data
            },
            order: [
                [query.sortBy || 'createdAt', query.sortOrder || 'DESC']
            ],
            limit,
            offset
        });

        return formatPaginatedResponse(users, page, limit);
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
