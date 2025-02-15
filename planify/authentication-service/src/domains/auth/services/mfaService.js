const speakeasy = require('speakeasy');
const QRCode = require('qrcode');
const {AppError} = require('../../../utils/errorHandler');
const mfaRepository = require('../repositories/mfa.repository');
const {EnableMfaDto, VerifyMfaDto, MfaResponseDto} = require('../dtos/mfa.dto');
const {generateBackupCodes} = require('../../../utils/mfaUtils');

class MfaService {
    /**
     * Enable MFA for a user
     * @param {Object} data - MFA enable data
     * @returns {Promise<Object>} MFA setup details
     */
    async enableMfa(data) {
        const enableDto = EnableMfaDto.from(data);

        // Check if MFA is already enabled
        const existingMfa = await mfaRepository.getUserWithMfa(enableDto.userId);
        if (existingMfa?.isEnabled) {
            throw new AppError('MFA is already enabled for this user', 400);
        }

        // Generate secret
        const secret = speakeasy.generateSecret({
            name: `YourApp:${enableDto.email || 'user'}`
        });

        // Generate QR code
        const qrCode = await QRCode.toDataURL(secret.otpauth_url);

        // Generate backup codes
        const backupCodes = generateBackupCodes();

        // Save MFA details
        await mfaRepository.createOrUpdateSecret(
            enableDto.userId,
            secret.base32,
            backupCodes
        );

        return MfaResponseDto.forEnable({
            secret: secret.base32,
            qrCode,
            backupCodes
        });
    }

    /**
     * Verify MFA code
     * @param {Object} data - MFA verification data
     * @returns {Promise<Object>} Verification result
     */
    async verifyMfa(data) {
        const verifyDto = VerifyMfaDto.from(data);

        // Get user's MFA secret
        const mfaSecret = await mfaRepository.getMfaSecretForUser(verifyDto.userId);
        if (!mfaSecret) {
            throw new AppError('MFA is not enabled for this user', 400);
        }

        // Verify TOTP
        const isValid = speakeasy.totp.verify({
            secret: mfaSecret.secret,
            encoding: 'base32',
            token: verifyDto.code,
            window: 1 // Allow 30 seconds clock skew
        });

        if (!isValid) {
            // Check if it's a backup code
            const isBackupCode = mfaSecret.backupCodes.includes(verifyDto.code);
            if (!isBackupCode) {
                throw new AppError('Invalid verification code', 400);
            }

            // Remove used backup code
            mfaSecret.backupCodes = mfaSecret.backupCodes.filter(
                code => code !== verifyDto.code
            );
            await mfaSecret.save();
        }

        // Enable MFA if it's the first verification
        if (!mfaSecret.isEnabled) {
            await mfaRepository.enableMfa(verifyDto.userId);
        }

        return MfaResponseDto.forVerify({
            verified: true,
            message: 'MFA verification successful'
        });
    }
}

module.exports = new MfaService();
