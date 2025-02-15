const {AppError} = require('../../../utils/errorHandler');

/**
 * DTO for password reset request
 */
class ForgotPasswordRequestDto {
    constructor(data) {
        this.email = data.email;
    }

    static validate(data) {
        if (!data.email) {
            throw new AppError('Email is required', 400);
        }

        const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
        if (!emailRegex.test(data.email)) {
            throw new AppError('Invalid email format', 400);
        }

        return new ForgotPasswordRequestDto(data);
    }
}

/**
 * DTO for password reset with token
 */
class ResetPasswordRequestDto {
    constructor(data) {
        this.token = data.token;
        this.password = data.password;
    }

    static validate(data) {
        const errors = [];

        if (!data.token) {
            errors.push('Token is required');
        } else if (data.token.length < 10) {
            errors.push('Invalid token format');
        }

        if (!data.password) {
            errors.push('Password is required');
        } else if (data.password.length < 6) {
            errors.push('Password must be at least 6 characters long');
        }

        if (errors.length > 0) {
            throw new AppError(errors.join('. '), 400);
        }

        return new ResetPasswordRequestDto(data);
    }
}

module.exports = {
    ForgotPasswordRequestDto,
    ResetPasswordRequestDto
};
