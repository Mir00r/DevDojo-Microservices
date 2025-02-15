const {Op} = require('sequelize');
const BaseRepository = require('./base.repository');
const {PasswordReset, User} = require('../../../../models');
const {sequelize} = require('../../../configs/database');
const bcrypt = require('bcryptjs');

/**
 * Repository for handling password reset operations
 */
class PasswordResetRepository extends BaseRepository {
    constructor() {
        super(PasswordReset);
    }

    /**
     * Invalidate all existing reset tokens for a user
     * @param {string} userId - User ID
     */
    async invalidateExistingTokens(userId) {
        return this.model.update(
            {isUsed: true},
            {where: {userId, isUsed: false}}
        );
    }

    /**
     * Find valid (not used and not expired) reset token
     * @param {string} token - Hashed reset token
     */
    async findValidToken(token) {
        return this.model.findOne({
            where: {
                token,
                isUsed: false,
                expiresAt: {[Op.gt]: new Date()}
            },
            include: [{
                model: User,
                as: 'user',
                attributes: ['id', 'email']
            }],
            order: [['createdAt', 'DESC']]
        });
    }

    /**
     * Execute password reset within transaction
     * @param {Object} passwordReset - Password reset record
     * @param {string} newPassword - New password to set
     */
    async executePasswordReset(passwordReset, newPassword) {
        const transaction = await sequelize.transaction();

        try {
            const hashedPassword = await bcrypt.hash(newPassword, 12);

            // Update user password
            await passwordReset.user.update(
                {password: hashedPassword},
                {transaction}
            );

            // Mark token as used
            await passwordReset.update(
                {isUsed: true},
                {transaction}
            );

            await transaction.commit();
        } catch (error) {
            await transaction.rollback();
            throw error;
        }
    }
}

module.exports = new PasswordResetRepository();
