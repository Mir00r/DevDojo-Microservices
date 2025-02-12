const BaseDto = require("./base.dto");
const {AppError} = require("../../../utils/errorHandler");

class LoginRequestDto extends BaseDto {
    constructor(data) {
        super();
        this.email = data.email;
        this.password = data.password;
    }

    static validate(data) {
        const errors = super.validate(data);
        if (!data.email) errors.push('Email is required');
        if (!data.password) errors.push('Password is required');
        if (errors.length > 0) {
            throw new AppError(errors.join(', '));
        }
        return new LoginRequestDto(data);
    }
}

class LoginResponseDto extends BaseDto {
    static success(user, tokens) {
        return {
            user: UserDto.fromEntity(user),
            accessToken: tokens.accessToken,
            refreshToken: tokens.refreshToken
        };
    }
}

class UserDto {
    static fromEntity(user) {
        return {
            id: user.id,
            name: user.name,
            email: user.email,
            role: user.role ? {
                name: user.role.name
            } : null,
            lastLogin: user.lastLogin,
            emailVerified: user.emailVerified
        };
    }
}

module.exports = {
    LoginRequestDto,
    LoginResponseDto,
    UserDto
};
