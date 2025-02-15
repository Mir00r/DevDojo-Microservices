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
const PasswordResetRepository = require('../repositories/password.repository');
const {Op} = require("sequelize");
const {LoginResponseDto} = require("../dtos/login.dto");


class PublicAuthService {
    /**
     * Register a new user
     * @param {RegisterDto} registerDto - Validated registration data
     * @returns {Promise<Object>} User and token data
     * @throws {AppError} If email is already registered
     */
    async registerUser(registerDto) {
        const transaction = await sequelize.transaction();

        try {
            // Check for existing user
            const existingUser = await userRepository.findByEmail(registerDto.email);
            if (existingUser) {
                throw new AppError('Email already registered', 400);
            }

            // Hash password
            const hashedPassword = await bcrypt.hash(registerDto.password, 12);

            // Create user
            const user = await userRepository.create({
                name: registerDto.name,
                email: registerDto.email,
                password: hashedPassword,
                roleId: registerDto.roleId
            }, {transaction});

            // Generate and send verification token
            const verificationToken = jwtUtils.generateVerificationToken(user);
            await emailService.sendVerificationEmail(user, verificationToken);

            await transaction.commit();

            // Generate JWT
            const token = jwtUtils.generateJWT(user);

            return {
                user,
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

    /**
     * Process forgot password request
     * @param {string} email - User's email address
     * @throws {AppError} When user is not found
     */
    async forgotPassword(email) {
        const user = await User.findOne({where: {email}});

        if (!user) {
            throw new AppError('No user found with this email', 404);
        }

        // Generate reset token
        const resetToken = jwtUtils.generatePasswordResetToken();
        const hashedToken = await bcrypt.hash(resetToken, 12);

        // Invalidate existing reset tokens
        await PasswordResetRepository.invalidateExistingTokens(user.id);

        // Create new password reset record
        await PasswordResetRepository.create({
            userId: user.id,
            token: hashedToken,
            expiresAt: new Date(Date.now() + 3600000), // 1 hour
            isUsed: false
        });

        // Send reset email
        // await emailService.sendPasswordResetEmail(user.email, resetToken);
    }

    /**
     * Reset password using token
     * @param {string} token - Reset token
     * @param {string} newPassword - New password
     * @throws {AppError} When token is invalid or expired
     */
    async resetPassword(token, newPassword) {
        // const hashedToken = crypto
        //     .createHash('sha256')
        //     .update(token)
        //     .digest('hex');

        // Find valid reset token
        const passwordReset = await PasswordResetRepository.findValidToken(
            token
        );

        if (!passwordReset) {
            throw new AppError('No valid password reset request found', 400);
        }

        // Update password within transaction
        await PasswordResetRepository.executePasswordReset(
            passwordReset,
            newPassword
        );

        return {
            status: 'success',
            message: 'Password has been reset successfully'
        };
    }

    /**
     * Verify user's email address
     * @param {string} token - Verification token
     * @throws {AppError} When token is invalid or user not found
     */
    async verifyEmail(token) {
        try {
            // Verify JWT token
            const decoded = jwt.verify(token, process.env.JWT_SECRET);

            // Find and update user
            const user = await userRepository.findById(decoded.id);
            if (!user) {
                throw new AppError('User not found', 404);
            }

            // Update email verification status
            const updatedUser = await userRepository.update(user.id, {
                emailVerified: true,
                emailVerifiedAt: new Date()
            });

            return updatedUser;

        } catch (error) {
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
