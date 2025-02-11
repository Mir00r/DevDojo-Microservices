const {validationResult} = require('express-validator');
const {AppError} = require('../utils/errorHandler');

/**
 * Middleware to validate request data using express-validator rules
 * @param {Array} validations - Array of express-validator validation rules
 */
const validate = (validations) => {
    return async (req, res, next) => {
        // Execute all validations
        await Promise.all(validations.map(validation => validation.run(req)));

        // Check for validation errors
        const errors = validationResult(req);
        if (!errors.isEmpty()) {
            const errorMessages = errors.array().map(err => ({
                field: err.param,
                message: err.msg
            }));

            return next(new AppError('Validation failed', 400, errorMessages));
        }

        next();
    };
};

module.exports = validate;
