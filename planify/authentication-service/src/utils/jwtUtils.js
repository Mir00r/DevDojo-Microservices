const jwt = require('jsonwebtoken');
const crypto = require('crypto');
require('dotenv').config();

/**
 * Generates a JWT token.
 * @param {object} payload - Data to be stored in the token.
 * @returns {string} - Signed JWT token.
 */
const generateToken = (payload) => {
    return jwt.sign(payload, process.env.JWT_SECRET, {expiresIn: '1h'});
};

/**
 * Verifies and decodes a JWT token.
 * @param {string} token - JWT token.
 * @returns {object} - Decoded token data.
 */
const verifyToken = (token) => {
    return jwt.verify(token, process.env.JWT_SECRET);
};

const generateJWT = (user) => {
    return jwt.sign(
        {
            id: user.id,
            email: user.email,
            role: user.role?.name || 'USER'
        },
        process.env.JWT_SECRET,
        {expiresIn: '1d'}
    );
}

const generateVerificationToken = (user) => {
    return jwt.sign(
        {id: user.id},
        process.env.JWT_SECRET,
        {expiresIn: '24h'}
    );
}

const generatePasswordResetToken = () => {
    return crypto.randomBytes(32).toString('hex');
}

const sanitizeUser = (user) => {
    const sanitized = user.toJSON();
    delete sanitized.password;
    delete sanitized.passwordResetToken;
    delete sanitized.passwordResetExpires;
    return sanitized;
}

module.exports = {
    generateToken,
    verifyToken,
    generateJWT,
    generateVerificationToken,
    generatePasswordResetToken,
    sanitizeUser
};
