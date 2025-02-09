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

// const validate = (schema) => {
//     return async (req, res, next) => {
//         try {
//             const validatedBody = await schema.validateAsync(req.body, {
//                 abortEarly: false,
//                 stripUnknown: true
//             });
//
//             // Replace req.body with validated data
//             req.body = validatedBody;
//             next();
//         } catch (error) {
//             if (error.isJoi) {
//                 // Format Joi validation errors
//                 const errors = error.details.map(detail => ({
//                     field: detail.context.key,
//                     message: detail.message
//                 }));
//
//                 return res.status(400).json({
//                     status: 'error',
//                     message: 'Validation failed',
//                     errors
//                 });
//             }
//
//             // Pass other errors to error handler
//             next(error);
//         }
//     };
// };

module.exports = validate;
