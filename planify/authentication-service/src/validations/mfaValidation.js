const {body} = require('express-validator');

/**
 * Validation rules for MFA operations
 */
const mfaValidation = {
    /**
     * Validation rules for enabling MFA
     */
    enable: [
        body('phoneNumber')
            .optional()
            .isMobilePhone()
            .withMessage('Invalid phone number format'),

        body('email')
            .optional()
            .isEmail()
            .normalizeEmail()
            .withMessage('Invalid email format')
    ],

    /**
     * Validation rules for verifying MFA
     */
    verify: [
        body('code')
            .exists()
            .withMessage('Verification code is required')
            .notEmpty()
            .withMessage('Verification code cannot be empty')
            .isString()
            .withMessage('Verification code must be a string')
            .trim()
            .custom((value) => {
                // Check for either 6-digit TOTP or 10-char backup code
                const isTOTP = /^[0-9]{6}$/.test(value);
                const isBackup = /^[A-Z0-9]{10}$/.test(value);
                if (!isTOTP && !isBackup) {
                    throw new Error('Invalid code format. Must be 6-digit TOTP or 10-character backup code');
                }
                return true;
            })
    ],

    /**
     * Validation rules for disabling MFA
     */
    disable: [
        body('code')
            .notEmpty()
            .withMessage('Current MFA code is required')
            .isString()
            .withMessage('Code must be a string')
            .matches(/^[0-9]{6}$/)
            .withMessage('Invalid TOTP code format')
            .trim(),

        body('confirm')
            .isBoolean()
            .withMessage('Confirmation is required')
            .equals('true')
            .withMessage('Must confirm MFA disable action')
    ]
};

module.exports = {mfaValidation};
