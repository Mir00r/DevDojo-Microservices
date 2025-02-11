const bcrypt = require('bcryptjs');
const jwt = require('jsonwebtoken');
const {sequelize} = require('../../../configs/database');
const {AppError} = require("../../../utils/errorHandler");
const jwtUtils = require('../../../utils/jwtUtils');
const {User, Role} = require('../../../../models');
const tokenService = require('./tokenService');


class ProtectedAuthService {

    async refreshToken(refreshToken) {
        try {
            // Verify refresh token and get user
            const user = await tokenService.verifyRefreshToken(refreshToken);

            // Revoke the used refresh token
            await tokenService.revokeRefreshToken(refreshToken);

            // Generate new tokens
            const newAccessToken = tokenService.generateAccessToken(user);
            const newRefreshToken = await tokenService.generateRefreshToken(user);

            return {
                accessToken: newAccessToken,
                refreshToken: newRefreshToken
            };
        } catch (error) {
            throw new AppError('Invalid refresh token', 401);
        }
    }

    async logout(refreshToken) {
        await tokenService.revokeRefreshToken(refreshToken);
    }

    async logoutAll(userId) {
        await tokenService.revokeAllUserTokens(userId);
    }

    async changePassword(userId, currentPassword, newPassword) {
        const user = await User.findByPk(userId);

        if (!(await bcrypt.compare(currentPassword, user.password))) {
            throw new AppError('Current password is incorrect', 401);
        }

        user.password = await bcrypt.hash(newPassword, 12);
        await user.save();
    }

    async updateProfile(userId, updateData) {
        const user = await User.findByPk(userId);

        if (!user) {
            throw new AppError('User not found', 404);
        }

        // Only allow certain fields to be updated
        const allowedUpdates = ['name', 'email'];
        const filteredData = Object.keys(updateData)
            .filter(key => allowedUpdates.includes(key))
            .reduce((obj, key) => {
                obj[key] = updateData[key];
                return obj;
            }, {});

        await user.update(filteredData);
        return jwtUtils.sanitizeUser(user);
    }

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
}

module.exports = new ProtectedAuthService();
