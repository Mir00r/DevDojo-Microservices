/**
 * Custom error class for application-specific errors
 */
class AppError extends Error {
    constructor(message, statusCode) {
        super(message);
        this.statusCode = statusCode;
        this.status = `${statusCode}`.startsWith('4') ? 'fail' : 'error';
        this.isOperational = true;

        Error.captureStackTrace(this, this.constructor);
    }
}

/**
 * Wrapper function to catch async errors
 * @param {Function} fn - Async function to be wrapped
 * @returns {Function} Express middlewares function
 */
const catchAsync = (fn) => {
    return (req, res, next) => {
        Promise.resolve(fn(req, res, next)).catch(next);
    };
};

/**
 * Global error handling middlewares
 */
const errorHandler = (err, req, res, next) => {
    err.statusCode = err.statusCode || 500;
    err.status = err.status || 'error';

    if (process.env.NODE_ENV === 'development') {
        sendErrorDev(err, res);
    } else {
        let error = {...err};
        error.message = err.message;

        // Handle specific types of errors
        if (error.name === 'SequelizeUniqueConstraintError') {
            error = handleSequelizeUniqueConstraintError(error);
        }
        if (error.name === 'SequelizeValidationError') {
            error = handleSequelizeValidationError(error);
        }
        if (error.name === 'JsonWebTokenError') {
            error = handleJWTError();
        }
        if (error.name === 'TokenExpiredError') {
            error = handleJWTExpiredError();
        }

        sendErrorProd(error, res);
    }
};

/**
 * Send detailed error response in development
 */
const sendErrorDev = (err, res) => {
    res.status(err.statusCode).json({
        status: err.status,
        error: err,
        message: err.message,
        stack: err.stack
    });
};

/**
 * Send clean error response in production
 */
const sendErrorProd = (err, res) => {
    // Operational, trusted error: send message to client
    if (err.isOperational) {
        res.status(err.statusCode).json({
            status: err.status,
            message: err.message
        });
    }
    // Programming or other unknown error: don't leak error details
    else {
        // Log error for debugging
        console.error('ERROR 💥', err);

        // Send generic message
        res.status(500).json({
            status: 'error',
            message: 'Something went wrong!'
        });
    }
};

/**
 * Handle Sequelize unique constraint errors
 */
const handleSequelizeUniqueConstraintError = (err) => {
    const message = `Duplicate field value: ${err.errors[0].path}. Please use another value!`;
    return new AppError(message, 400);
};

/**
 * Handle Sequelize validation errors
 */
const handleSequelizeValidationError = (err) => {
    const errors = err.errors.map(error => error.message);
    const message = `Invalid input data. ${errors.join('. ')}`;
    return new AppError(message, 400);
};

/**
 * Handle JWT verification errors
 */
const handleJWTError = () => {
    return new AppError('Invalid token. Please log in again!', 401);
};

/**
 * Handle JWT expiration errors
 */
const handleJWTExpiredError = () => {
    return new AppError('Your token has expired! Please log in again.', 401);
};

/**
 * Handle unhandled promise rejections
 */
const handleUncaughtExceptions = () => {
    process.on('uncaughtException', (err) => {
        console.error('UNCAUGHT EXCEPTION! 💥 Shutting down...');
        console.error(err.name, err.message);
        process.exit(1);
    });
};

/**
 * Handle unhandled promise rejections
 */
const handleUnhandledRejections = (server) => {
    process.on('unhandledRejection', (err) => {
        console.error('UNHANDLED REJECTION! 💥 Shutting down...');
        console.error(err.name, err.message);
        server.close(() => {
            process.exit(1);
        });
    });
};

module.exports = {
    AppError,
    catchAsync,
    errorHandler,
    handleUncaughtExceptions,
    handleUnhandledRejections
};
