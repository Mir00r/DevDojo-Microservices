const bcrypt = require('bcrypt');

const SALT_ROUNDS = 10;

/**
 * Hashes a plain text password.
 * @param {string} password - User's raw password.
 * @returns {Promise<string>} - Returns the hashed password.
 */
const hashPassword = async (password) => {
    return await bcrypt.hash(password, SALT_ROUNDS);
};

/**
 * Compares a given password with its hashed version.
 * @param {string} password - Plain text password.
 * @param {string} hashedPassword - Stored hashed password.
 * @returns {Promise<boolean>} - Returns true if matched.
 */
const comparePassword = async (password, hashedPassword) => {
    return await bcrypt.compare(password, hashedPassword);
};

module.exports = {hashPassword, comparePassword};
