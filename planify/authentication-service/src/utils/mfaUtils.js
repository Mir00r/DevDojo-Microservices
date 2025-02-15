const crypto = require('crypto');

/**
 * Generate backup codes for MFA
 * @param {number} count - Number of backup codes to generate
 * @param {number} length - Length of each backup code
 * @returns {string[]} Array of backup codes
 */
const generateBackupCodes = (count = 8, length = 10) => {
    const codes = [];
    for (let i = 0; i < count; i++) {
        const code = crypto.randomBytes(length)
            .toString('hex')
            .toUpperCase()
            .slice(0, length);
        codes.push(code);
    }
    return codes;
};

module.exports = {
    generateBackupCodes
};
