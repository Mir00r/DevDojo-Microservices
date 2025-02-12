const bcrypt = require('bcryptjs');
const crypto = require('crypto');
const jwt = require('jsonwebtoken');
const {sequelize} = require('../../../configs/database');
const {AppError} = require("../../../utils/errorHandler");
const {User, Role, PasswordReset} = require('../../../../models');
const emailService = require('../../../utils/emailService');
const jwtUtils = require('../../../utils/jwtUtils');
const tokenService = require('./tokenService');
const userRepository = require('../repositories/user.repository');
const {Op} = require("sequelize");
const {LoginResponseDto} = require("../dtos/login.dto");


class PublicAuthService {
    async registerUser(name, email, password, roleId = 2) { // Default roleId 2 for regular users
        const transaction = await sequelize.transaction();

        try {
            const existingUser = await User.findOne({
                where: {email},
                transaction
            });

            if (existingUser) {
                throw new AppError('Email already registered', 400);
            }

            const hashedPassword = await bcrypt.hash(password, 12);

            const user = await User.create({
                name,
                email,
                password: hashedPassword,
                roleId
            }, {transaction});

            const verificationToken = jwtUtils.generateVerificationToken(user);
            // Use the email service to send verification email
            await emailService.sendVerificationEmail(user, verificationToken);

            await transaction.commit();

            const token = jwtUtils.generateJWT(user);

            return {
                user: jwtUtils.sanitizeUser(user),
                token
            };
        } catch (error) {
            await transaction.rollback();
            throw error;
        }
    }

    async login(loginDto) {
        // Using repository instead of direct model access
        const user = await userRepository.findByEmail(loginDto.email);

        if (!user || !(await bcrypt.compare(loginDto.password, user.password))) {
            throw new AppError('Invalid email or password', 401);
        }

        if (!user.emailVerified) {
            throw new AppError('Please verify your email first', 401);
        }

        const accessToken = tokenService.generateAccessToken(user);
        const refreshToken = await tokenService.generateRefreshToken(user);

        // Using repository method
        await userRepository.updateLastLogin(user.id);

        // Using LoginResponseDto to format response
        return LoginResponseDto.success(user, {
            accessToken,
            refreshToken
        });
    }

    async forgotPassword(email) {
        const user = await User.findOne({where: {email}});

        if (!user) {
            throw new AppError('No user found with this email', 404);
        }

        // Generate reset token
        const resetToken = jwtUtils.generatePasswordResetToken();
        const hashedToken = await bcrypt.hash(resetToken, 12);

        // Invalidate any existing reset tokens for this user
        await PasswordReset.update(
            {isUsed: true},
            {where: {userId: user.id, isUsed: false}}
        );

        // Create new password reset record
        await PasswordReset.create({
            userId: user.id,
            token: hashedToken,
            expiresAt: new Date(Date.now() + 3600000), // 1 hour
            isUsed: false
        });

        // Send reset email
        // await emailService.sendPasswordResetEmail(user.email, resetToken);
    }

    async resetPassword(token, newPassword) {
        console.log('Received Token:', token);
        // Hash the received token the same way
        const hashedToken = crypto
            .createHash('sha256')
            .update(token)
            .digest('hex');

        console.log('Hashed Received Token:', hashedToken);

        // First find all valid (not used and not expired) reset records
        const passwordReset = await PasswordReset.findOne({
            where: {
                isUsed: false,
                expiresAt: {
                    [Op.gt]: new Date()
                },
                token: token
                // We'll check the token match using bcrypt.compare later
            },
            include: [{
                model: User,
                as: 'user',
                attributes: ['id', 'email'] // Only select needed fields
            }],
            order: [['createdAt', 'DESC']] // Get the most recent one
        });

        // Add debug logging
        console.log('Found password reset record:', passwordReset);

        if (!passwordReset) {
            throw new AppError('No valid password reset request found', 400);
        }

        // Compare the provided token with stored hashed token
        // const isValidToken = await bcrypt.compare(token, passwordReset.token);
        // console.log('Token comparison result:', isValidToken);

        // if (!isValidToken) {
        //     throw new AppError('Invalid reset token', 400);
        // }

        // Update password within a transaction
        const transaction = await sequelize.transaction();

        try {
            // Hash new password
            const hashedPassword = await bcrypt.hash(newPassword, 12);

            // Update user password
            await passwordReset.user.update(
                {password: hashedPassword},
                {transaction}
            );

            // Mark token as used
            await passwordReset.update(
                {isUsed: true},
                {transaction}
            );

            await transaction.commit();

            return {
                status: 'success',
                message: 'Password has been reset successfully'
            };
        } catch (error) {
            await transaction.rollback();
            throw new AppError('Failed to reset password', 500);
        }
    }

    async verifyEmail(token) {
        try {
            console.log('Starting email verification with token:', token);

            // Verify JWT token
            console.log('Verifying JWT token...');
            const decoded = jwt.verify(token, process.env.JWT_SECRET);
            console.log('Decoded token:', decoded);

            // Find user
            console.log('Finding user with id:', decoded.id);
            const user = await User.findOne({
                where: {id: decoded.id}
            });
            console.log('User found:', user ? 'Yes' : 'No');

            if (!user) {
                throw new AppError('User not found', 404);
            }

            console.log('Before update - User emailVerified status:', {
                id: user.id,
                email: user.email,
                emailVerified: user.emailVerified
            });

            // Update the user - Corrected version
            const updatedUser = await user.update({
                emailVerified: true,
                updatedAt: new Date()
            });

            console.log('After update - User status:', {
                id: updatedUser.id,
                email: updatedUser.email,
                emailVerified: updatedUser.emailVerified
            });

            return {
                status: 'success',
                message: 'Email verified successfully',
                data: {
                    email: updatedUser.email,
                    emailVerified: updatedUser.emailVerified
                }
            };

        } catch (error) {
            console.error('Detailed verification error:', {
                message: error.message,
                name: error.name,
                stack: error.stack,
                token: token,
                jwtSecret: process.env.JWT_SECRET ? 'Exists' : 'Missing'
            });

            if (error.name === 'JsonWebTokenError') {
                throw new AppError('Invalid verification token', 400);
            }
            if (error.name === 'TokenExpiredError') {
                throw new AppError('Verification token has expired', 401);
            }

            throw new AppError(`Email verification failed: ${error.message}`, 500);
        }
    }
}

module.exports = new PublicAuthService();
