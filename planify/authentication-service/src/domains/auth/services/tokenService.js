const jwt = require('jsonwebtoken');
const crypto = require('crypto');
const {User, RefreshToken} = require('../../../../models');
const {AppError} = require('../../../utils/errorHandler');

class TokenService {
    generateAccessToken(user) {
        return jwt.sign(
            {
                id: user.id,
                email: user.email,
                role: user.role?.name || 'USER'
            },
            process.env.JWT_SECRET,
            {expiresIn: '15m'} // Short lived access token
        );
    }

    async generateRefreshToken(user) {
        // Generate a random token
        const refreshToken = crypto.randomBytes(40).toString('hex');

        // Save refresh token in database
        await RefreshToken.create({
            token: refreshToken,
            userId: user.id,
            expiresAt: new Date(Date.now() + 7 * 24 * 60 * 60 * 1000) // 7 days
        });

        return refreshToken;
    }

    async verifyRefreshToken(refreshToken) {
        const token = await RefreshToken.findOne({
            where: {
                token: refreshToken,
                isRevoked: false
            },
            include: [{
                model: User,
                as: 'user',
                attributes: ['id', 'email']
            }]
        });

        if (!token) {
            throw new AppError('Invalid refresh token', 401);
        }

        // Check if token is expired
        if (new Date() > token.expiresAt) {
            await token.update({isRevoked: true});
            throw new AppError('Refresh token has expired', 401);
        }

        return token.user;
    }

    async revokeRefreshToken(refreshToken) {
        const token = await RefreshToken.findOne({
            where: {token: refreshToken}
        });

        if (token) {
            await token.update({isRevoked: true});
        }
    }

    async revokeAllUserTokens(userId) {
        await RefreshToken.update(
            {isRevoked: true},
            {where: {userId}}
        );
    }
}

module.exports = new TokenService();
