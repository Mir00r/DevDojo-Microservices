const jwt = require('jsonwebtoken');
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

module.exports = {generateToken, verifyToken};
