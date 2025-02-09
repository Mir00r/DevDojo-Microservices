const bcrypt = require('bcryptjs');
const jwt = require('jsonwebtoken');
const {sequelize} = require('../../../configs/database');
const {AppError} = require("../../../utils/errorHandler");
const {User, Role} = require('../../../../models');
const emailService = require('../../../utils/emailService');


class AuthService {
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

            const verificationToken = this.generateVerificationToken(user);
            // Use the email service to send verification email
            await emailService.sendVerificationEmail(user, verificationToken);

            await transaction.commit();

            const token = this.generateJWT(user);

            return {
                user: this.sanitizeUser(user),
                token
            };
        } catch (error) {
            await transaction.rollback();
            throw error;
        }
    }

    async loginUser(email, password) {
        const user = await User.findOne({
            where: {email},
            include: [{
                model: Role,
                as: 'role',
                attributes: ['name']
            }]
        });

        if (!user || !(await bcrypt.compare(password, user.password))) {
            throw new AppError('Invalid email or password', 401);
        }

        if (!user.emailVerified) {
            throw new AppError('Please verify your email first', 401);
        }

        const token = this.generateJWT(user);

        // Update last login
        await user.update({lastLogin: new Date()});

        return {
            user: this.sanitizeUser(user),
            token
        };
    }

    async forgotPassword(email) {
        const user = await User.findOne({where: {email}});

        if (!user) {
            throw new AppError('No user found with this email', 404);
        }

        const resetToken = this.generatePasswordResetToken();
        const resetTokenHashed = await bcrypt.hash(resetToken, 12);

        await user.update({
            passwordResetToken: resetTokenHashed,
            passwordResetExpires: new Date(Date.now() + 3600000) // 1 hour
        });

        await this.sendPasswordResetEmail(user.email, resetToken);
    }

    async resetPassword(token, newPassword) {
        const user = await User.findOne({
            where: {
                passwordResetToken: await bcrypt.hash(token, 12),
                passwordResetExpires: {[Op.gt]: new Date()}
            }
        });

        if (!user) {
            throw new AppError('Invalid or expired password reset token', 400);
        }

        user.password = await bcrypt.hash(newPassword, 12);
        user.passwordResetToken = null;
        user.passwordResetExpires = null;
        await user.save();
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
                where: { id: decoded.id }
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

            // Update user
            // console.log('Updating user email verification status...');
            // await user.update({
            //     emailVerified: true
            // });
            // console.log('User updated successfully');

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

    async logout(userId) {
        // Implement token blacklisting if needed
        return true;
    }

    async changePassword(userId, currentPassword, newPassword) {
        const user = await User.findByPk(userId);

        if (!(await bcrypt.compare(currentPassword, user.password))) {
            throw new AppError('Current password is incorrect', 401);
        }

        user.password = await bcrypt.hash(newPassword, 12);
        await user.save();
    }

    async updateProfile(userId, updateData) {
        const user = await User.findByPk(userId);

        if (!user) {
            throw new AppError('User not found', 404);
        }

        // Only allow certain fields to be updated
        const allowedUpdates = ['name', 'email'];
        const filteredData = Object.keys(updateData)
            .filter(key => allowedUpdates.includes(key))
            .reduce((obj, key) => {
                obj[key] = updateData[key];
                return obj;
            }, {});

        await user.update(filteredData);
        return this.sanitizeUser(user);
    }

    // Admin only methods
    async getAllUsers({page = 1, limit = 10, search = ''}) {
        const offset = (page - 1) * limit;
        const where = search ? {
            [Op.or]: [
                {name: {[Op.iLike]: `%${search}%`}},
                {email: {[Op.iLike]: `%${search}%`}}
            ]
        } : {};

        const users = await User.findAndCountAll({
            where,
            include: [{
                model: Role,
                as: 'role',
                attributes: ['name']
            }],
            limit,
            offset,
            order: [['createdAt', 'DESC']]
        });

        return {
            users: users.rows.map(this.sanitizeUser),
            total: users.count,
            pages: Math.ceil(users.count / limit)
        };
    }

    async updateUserRole(userId, roleId) {
        const user = await User.findByPk(userId);

        if (!user) {
            throw new AppError('User not found', 404);
        }

        const role = await Role.findByPk(roleId);
        if (!role) {
            throw new AppError('Role not found', 404);
        }

        await user.update({roleId});
        return this.sanitizeUser(user);
    }

    async deleteUser(userId) {
        const user = await User.findByPk(userId);

        if (!user) {
            throw new AppError('User not found', 404);
        }

        await user.destroy(); // Soft delete if paranoid is true
    }

    // Helper methods
    generateJWT(user) {
        return jwt.sign(
            {
                id: user.id,
                email: user.email,
                role: user.role?.name || 'USER'
            },
            process.env.JWT_SECRET,
            {expiresIn: '1d'}
        );
    }

    generateVerificationToken(user) {
        return jwt.sign(
            {id: user.id},
            process.env.JWT_SECRET,
            {expiresIn: '24h'}
        );
    }

    generatePasswordResetToken() {
        return crypto.randomBytes(32).toString('hex');
    }

    sanitizeUser(user) {
        const sanitized = user.toJSON();
        delete sanitized.password;
        delete sanitized.passwordResetToken;
        delete sanitized.passwordResetExpires;
        return sanitized;
    }
}

module.exports = new AuthService();
