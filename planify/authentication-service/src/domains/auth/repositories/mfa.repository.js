const BaseRepository = require('../repositories/base.repository');
const {MfaSecret, User} = require('../../../../models');
const {Op} = require('sequelize');

/**
 * Repository for MFA operations
 * @extends BaseRepository
 */
class MfaRepository extends BaseRepository {
    constructor() {
        super(MfaSecret);
    }

    /**
     * Find MFA secret by user ID
     * @param {string} userId - User ID
     * @returns {Promise<MfaSecret>}
     */
    async findByUserId(userId) {
        return this.findOne({
            where: {userId}
        });
    }

    /**
     * Create or update MFA secret
     * @param {string} userId - User ID
     * @param {string} secret - TOTP secret
     * @param {string[]} backupCodes - Backup codes
     * @param {Object} transaction - Sequelize transaction
     * @returns {Promise<MfaSecret>}
     */
    async createOrUpdateSecret(userId, secret, backupCodes, transaction = null) {
        const existingSecret = await this.findOne({
            where: {userId}
        });

        if (existingSecret) {
            return this.update(existingSecret.id, {
                secret,
                backupCodes,
                isEnabled: false
            }, {transaction});
        }

        return this.create({
            userId,
            secret,
            backupCodes,
            isEnabled: false
        }, {transaction});
    }

    /**
     * Enable MFA for user
     * @param {string} userId - User ID
     * @param {Object} transaction - Sequelize transaction
     * @returns {Promise<[MfaSecret, number]>}
     */
    async enableMfa(userId, transaction = null) {
        const mfaSecret = await this.findOne({where: {userId}});
        if (!mfaSecret) return [null, 0];

        const options = transaction ? {transaction} : {};

        const [updatedMfa, updatedUser] = await Promise.all([
            this.update(mfaSecret.id,
                {isEnabled: true},
                options
            ),
            User.update(
                {isMfaEnabled: true},
                {
                    where: {id: userId},
                    ...options
                }
            )
        ]);

        return [updatedMfa, updatedUser[0]];
    }

    /**
     * Disable MFA for user
     * @param {string} userId - User ID
     * @param {Object} transaction - Sequelize transaction
     * @returns {Promise<boolean>}
     */
    async disableMfa(userId, transaction = null) {
        const mfaSecret = await this.findOne({where: {userId}});
        if (!mfaSecret) return false;

        const options = transaction ? {transaction} : {};

        await Promise.all([
            this.delete(mfaSecret.id, options),
            User.update(
                {isMfaEnabled: false},
                {
                    where: {id: userId},
                    ...options
                }
            )
        ]);

        return true;
    }

    /**
     * Remove used backup code
     * @param {string} userId - User ID
     * @param {string} usedCode - Used backup code
     * @param {Object} transaction - Sequelize transaction
     * @returns {Promise<MfaSecret>}
     */
    async removeBackupCode(userId, usedCode, transaction = null) {
        const mfaSecret = await this.findByUserId(userId);
        if (!mfaSecret) return null;

        const updatedBackupCodes = mfaSecret.backupCodes.filter(
            code => code !== usedCode
        );

        return this.update(mfaSecret.id,
            {backupCodes: updatedBackupCodes},
            {transaction}
        );
    }

    /**
     * Verify backup code exists
     * @param {string} userId - User ID
     * @param {string} code - Backup code to verify
     * @returns {Promise<boolean>}
     */
    async verifyBackupCode(userId, code) {
        const mfaSecret = await this.findByUserId(userId);
        return mfaSecret?.backupCodes.includes(code) || false;
    }

    /**
     * Get user with MFA details
     * @param {string} userId - User ID
     * @returns {Promise<User>}
     */
    async getUserWithMfa(userId) {
        return User.findOne({
            where: {id: userId},
            include: [{
                model: this.model,
                as: 'mfaSecret'
            }]
        });
    }

    async getMfaSecretForUser(userId) {
        const mfaSecret = await this.model.findOne({
            where: { userId },
            include: [{
                model: User,
                as: 'user'
            }]
        });

        if (!mfaSecret) {
            return null;
        }

        // Ensure backupCodes is initialized
        if (!Array.isArray(mfaSecret.backupCodes)) {
            mfaSecret.backupCodes = [];
            await mfaSecret.save();
        }

        return mfaSecret;
    }
}

module.exports = new MfaRepository();
