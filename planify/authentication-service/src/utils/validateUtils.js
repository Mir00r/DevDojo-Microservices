/**
 * Checks if a string is a valid email.
 * @param {string} email - User email input.
 * @returns {boolean} - Returns true if valid.
 */
const isValidEmail = (email) => {
    const regex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    return regex.test(email);
};

/**
 * Checks if a password meets security criteria.
 * @param {string} password - User password.
 * @returns {boolean} - Returns true if password is strong.
 */
const isStrongPassword = (password) => {
    return password.length >= 8;
};

module.exports = {isValidEmail, isStrongPassword};
