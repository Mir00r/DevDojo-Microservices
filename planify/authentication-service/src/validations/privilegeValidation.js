const {body, param, query} = require('express-validator');

const privilegeValidation = {
    createPrivilege: [
        body('name')
            .trim()
            .notEmpty()
            .withMessage('Privilege name is required')
            .isLength({min: 2, max: 50})
            .withMessage('Privilege name must be between 2 and 50 characters')
            .matches(/^[A-Z_][A-Z0-9_]*$/)
            .withMessage('Privilege name must be uppercase and contain only letters, numbers, and underscores'),

        body('description')
            .optional()
            .trim()
            .isLength({max: 200})
            .withMessage('Description cannot exceed 200 characters'),

        body('module')
            .trim()
            .notEmpty()
            .withMessage('Module name is required')
            .isLength({min: 2, max: 50})
            .withMessage('Module name must be between 2 and 50 characters')
            .matches(/^[A-Z_][A-Z0-9_]*$/)
            .withMessage('Module name must be uppercase and contain only letters, numbers, and underscores'),

        body('isActive')
            .optional()
            .isBoolean()
            .withMessage('isActive must be a boolean value')
    ],

    updatePrivilege: [
        param('id')
            .isUUID(4)
            .withMessage('Invalid privilege ID'),

        body('name')
            .optional()
            .trim()
            .isLength({min: 2, max: 50})
            .withMessage('Privilege name must be between 2 and 50 characters')
            .matches(/^[A-Z_][A-Z0-9_]*$/)
            .withMessage('Privilege name must be uppercase and contain only letters, numbers, and underscores'),

        body('description')
            .optional()
            .trim()
            .isLength({max: 200})
            .withMessage('Description cannot exceed 200 characters'),

        body('module')
            .optional()
            .trim()
            .isLength({min: 2, max: 50})
            .withMessage('Module name must be between 2 and 50 characters')
            .matches(/^[A-Z_][A-Z0-9_]*$/)
            .withMessage('Module name must be uppercase and contain only letters, numbers, and underscores'),

        body('isActive')
            .optional()
            .isBoolean()
            .withMessage('isActive must be a boolean value')
    ],

    checkId: [
        param('id')
            .isUUID(4)
            .withMessage('Invalid privilege ID')
    ],

    listPrivileges: [
        query('search')
            .optional()
            .trim()
            .isLength({min: 1})
            .withMessage('Search term cannot be empty'),

        query('module')
            .optional()
            .trim()
            .matches(/^[A-Z_][A-Z0-9_]*$/)
            .withMessage('Module name must be uppercase and contain only letters, numbers, and underscores'),

        query('isActive')
            .optional()
            .isBoolean()
            .withMessage('isActive must be a boolean value'),

        query('page')
            .optional()
            .isInt({min: 1})
            .withMessage('Page must be a positive integer'),

        query('limit')
            .optional()
            .isInt({min: 1, max: 100})
            .withMessage('Limit must be between 1 and 100'),

        query('sortBy')
            .optional()
            .isIn(['name', 'module', 'createdAt', 'updatedAt'])
            .withMessage('Invalid sort field'),

        query('sortOrder')
            .optional()
            .isIn(['ASC', 'DESC'])
            .withMessage('Sort order must be ASC or DESC')
    ]
};

module.exports = {privilegeValidation};
