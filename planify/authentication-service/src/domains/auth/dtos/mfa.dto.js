/**
 * Data Transfer Objects for MFA operations
 */

class EnableMfaDto {
    constructor(data) {
        this.userId = data.userId;
        this.phoneNumber = data.phoneNumber; // Optional: If using SMS
        this.email = data.email;        // Optional: If using email
    }

    static from(data) {
        if (!data.userId) {
            throw new Error('User ID is required');
        }
        return new EnableMfaDto(data);
    }
}

class VerifyMfaDto {
    constructor(data) {
        this.userId = data.userId;
        this.code = data.code;
    }

    static from(data) {
        if (!data.userId || !data.code) {
            throw new Error('User ID and verification code are required');
        }

        // if (data.code.length !== 6) {
        //     throw new Error('Invalid verification code format');
        // }

        return new VerifyMfaDto(data);
    }
}

class MfaResponseDto {
    static forEnable(data) {
        return {
            secret: data.secret,
            qrCode: data.qrCode,
            backupCodes: data.backupCodes
        };
    }

    static forVerify(data) {
        return {
            verified: data.verified,
            message: data.message
        };
    }
}

module.exports = {
    EnableMfaDto,
    VerifyMfaDto,
    MfaResponseDto
};
