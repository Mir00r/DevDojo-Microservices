const {body, param} = require('express-validator');

const roleValidation = {
    createRole: [
        body('name')
            .trim()
            .notEmpty()
            .withMessage('Role name is required')
            .isLength({min: 2, max: 50})
            .withMessage('Role name must be between 2 and 50 characters')
            .matches(/^[A-Za-z_][A-Za-z0-9_]*$/)
            .withMessage('Role name must start with a letter and contain only letters, numbers, and underscores'),

        body('description')
            .optional()
            .trim()
            .isLength({max: 200})
            .withMessage('Description cannot exceed 200 characters')
    ],

    updateRole: [
        param('id')
            .isInt()
            .withMessage('Invalid role ID'),

        body('name')
            .optional()
            .trim()
            .isLength({min: 2, max: 50})
            .withMessage('Role name must be between 2 and 50 characters')
            .matches(/^[A-Za-z_][A-Za-z0-9_]*$/)
            .withMessage('Role name must start with a letter and contain only letters, numbers, and underscores'),

        body('description')
            .optional()
            .trim()
            .isLength({max: 200})
            .withMessage('Description cannot exceed 200 characters')
    ],

    checkId: [
        param('id')
            .isInt()
            .withMessage('Invalid role ID')
    ]
};

module.exports = {roleValidation};
