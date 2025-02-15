const {AppError} = require('../../../utils/errorHandler');


/**
 * DTO for email verification
 */
class VerifyEmailRequestDto {
    constructor(data) {
        this.token = data.token;
    }

    static validate(data) {
        if (!data.token) {
            throw new AppError('Token is required', 400);
        }

        // Basic JWT format validation
        const jwtRegex = /^[A-Za-z0-9-_]+\.[A-Za-z0-9-_]+\.[A-Za-z0-9-_]*$/;
        if (!jwtRegex.test(data.token)) {
            throw new AppError('Invalid token format', 400);
        }

        return new VerifyEmailRequestDto(data);
    }
}

/**
 * DTO for successful email verification response
 */
class EmailVerificationResponseDto {
    constructor(data) {
        this.email = data.email;
        this.verifiedAt = data.updatedAt;
        this.emailVerified = data.emailVerified;
    }

    static fromEntity(user) {
        return new EmailVerificationResponseDto({
            email: user.email,
            updatedAt: user.updatedAt,
            emailVerified: user.emailVerified
        });
    }
}

module.exports = {
    VerifyEmailRequestDto,
    EmailVerificationResponseDto
};
