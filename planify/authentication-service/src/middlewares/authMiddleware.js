const jwt = require('jsonwebtoken');
const {AppError} = require("../utils/errorHandler");
const {User, Role} = require('../../models');
require('dotenv').config();

class AuthMiddleware {
    authenticate = async (req, res, next) => {
        try {
            // Check if token exists
            const authHeader = req.headers.authorization;
            if (!authHeader || !authHeader.startsWith('Bearer ')) {
                throw new AppError('No token provided', 401);
            }

            const token = authHeader.split(' ')[1];

            // Verify token
            const decoded = jwt.verify(token, process.env.JWT_SECRET);

            // Get user
            const user = await User.findByPk(decoded.id, {
                include: [{
                    model: Role,
                    as: 'role',
                    attributes: ['name']
                }]
            });

            if (!user) {
                throw new AppError('User not found', 401);
            }

            // Check if user is still active
            if (!user.isActive) {
                throw new AppError('User account is deactivated', 401);
            }

            // Attach user to request
            req.user = user;
            next();
        } catch (error) {
            if (error.name === 'JsonWebTokenError') {
                next(new AppError('Invalid token', 401));
            } else if (error.name === 'TokenExpiredError') {
                next(new AppError('Token expired', 401));
            } else {
                next(error);
            }
        }
    };

    requireRole = (...roles) => {
        return (req, res, next) => {
            if (!roles.includes(req.user.role.name)) {
                throw new AppError('You do not have permission to perform this action', 403);
            }
            next();
        };
    };

    checkOwnership = (req, res, next) => {
        if (req.user.id !== req.params.id && req.user.role.name !== 'ADMIN') {
            throw new AppError('You can only access your own resources', 403);
        }
        next();
    };

    internalBasicAuth = (req, res, next) => {
        try {
            // Get authorization header
            const authHeader = req.headers.authorization;

            if (!authHeader || !authHeader.startsWith('Basic ')) {
                throw new AppError('Basic authentication required', 401);
            }

            // Extract credentials
            const base64Credentials = authHeader.split(' ')[1];
            const credentials = Buffer.from(base64Credentials, 'base64').toString('utf8');
            const [username, password] = credentials.split(':');

            // Check credentials
            if (username !== process.env.INTERNAL_API_USERNAME ||
                password !== process.env.INTERNAL_API_PASSWORD) {
                throw new AppError('Invalid credentials', 401);
            }

            next();
        } catch (error) {
            next(error);
        }
    };
}

module.exports = new AuthMiddleware();
