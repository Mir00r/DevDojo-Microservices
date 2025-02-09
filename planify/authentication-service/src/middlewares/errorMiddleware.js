/**
 * Centralized error-handling middlewares.
 */
const errorHandler = (err, req, res, next) => {
    console.error(err.stack);

    const statusCode = err.status || 500;
    res.status(statusCode).json({
        error: err.message || 'Internal Server Error',
    });
};

module.exports = errorHandler;
