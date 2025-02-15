const {AppError} = require("../../../utils/errorHandler");

/**
 * DTO for user registration request
 * @class RegisterDto
 */
class RegisterDto {
    /**
     * Create a new RegisterDto instance
     * @param {Object} data - Raw request data
     */
    constructor(data) {
        this.name = data.name;
        this.email = data.email?.toLowerCase();
        this.password = data.password;
        this.roleId = data.roleId || 2; // Default to regular user role
    }

    /**
     * Validate DTO data
     * @static
     * @param {Object} data - Raw request data
     * @returns {RegisterDto} Validated DTO instance
     * @throws {AppError} If validation fails
     */
    static validate(data) {
        const errors = [];

        // Name validation
        if (!data.name || typeof data.name !== 'string') {
            errors.push('Name is required and must be a string');
        } else if (data.name.length < 2) {
            errors.push('Name must be at least 2 characters long');
        }

        // Email validation
        const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
        if (!data.email || !emailRegex.test(data.email)) {
            errors.push('Please provide a valid email address');
        }

        // Password validation
        if (!data.password || typeof data.password !== 'string') {
            errors.push('Password is required and must be a string');
        } else if (data.password.length < 6) {
            errors.push('Password must be at least 6 characters long');
        }

        // Role ID validation (optional)
        if (data.roleId && !Number.isInteger(Number(data.roleId))) {
            errors.push('Role ID must be a valid integer');
        }

        if (errors.length > 0) {
            throw new AppError(errors.join('. '), 400);
        }

        return new RegisterDto(data);
    }
}

/**
 * DTO for user response data
 * @class UserResponseDto
 */
class UserResponseDto {
    /**
     * Transform user entity to DTO
     * @param {Object} user - User entity from database
     * @returns {Object} Sanitized user data
     */
    static toResponse(user) {
        return {
            id: user.id,
            name: user.name,
            email: user.email,
            role: user.role ? {
                id: user.role.id,
                name: user.role.name
            } : null,
            emailVerified: user.emailVerified,
            lastLogin: user.lastLogin,
            createdAt: user.createdAt
        };
    }
}

/**
 * DTO for authentication response
 * @class AuthResponseDto
 */
class AuthResponseDto {
    /**
     * Transform authentication result to DTO
     * @param {Object} data - Authentication data
     * @param {Object} data.user - User entity
     * @param {string} data.token - JWT token
     * @returns {Object} Formatted authentication response
     */
    static toResponse({user, token}) {
        return {
            user: UserResponseDto.toResponse(user),
            token
        };
    }
}

module.exports = {
    RegisterDto,
    UserResponseDto,
    AuthResponseDto
};
