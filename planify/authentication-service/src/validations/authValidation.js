const {body, param, query} = require('express-validator');

const authValidation = {
    register: [
        body('name')
            .trim()
            .notEmpty()
            .withMessage('Name is required')
            .isLength({min: 2, max: 50})
            .withMessage('Name must be between 2 and 50 characters'),

        body('email')
            .trim()
            .notEmpty()
            .withMessage('Email is required')
            .isEmail()
            .withMessage('Please provide a valid email')
            .normalizeEmail(),

        body('password')
            .notEmpty()
            .withMessage('Password is required')
            .isLength({min: 6})
            .withMessage('Password must be at least 6 characters')
            .matches(/^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)/)
            .withMessage('Password must contain at least one uppercase letter, one lowercase letter, and one number'),

        body('roleId')
            .optional()
            .isInt()
            .withMessage('Role ID must be an integer')
    ],

    login: [
        body('email')
            .trim()
            .notEmpty()
            .withMessage('Email is required')
            .isEmail()
            .withMessage('Please provide a valid email'),

        body('password')
            .notEmpty()
            .withMessage('Password is required')
    ],

    forgotPassword: [
        body('email')
            .trim()
            .notEmpty()
            .withMessage('Email is required')
            .isEmail()
            .withMessage('Please provide a valid email')
    ],

    resetPassword: [
        body('token')
            .notEmpty()
            .withMessage('Token is required'),

        body('password')
            .notEmpty()
            .withMessage('Password is required')
            .isLength({min: 6})
            .withMessage('Password must be at least 6 characters')
            .matches(/^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)/)
            .withMessage('Password must contain at least one uppercase letter, one lowercase letter, and one number'),

        body('passwordConfirm')
            .notEmpty()
            .withMessage('Password confirmation is required')
            .custom((value, {req}) => {
                if (value !== req.body.password) {
                    throw new Error('Password confirmation does not match password');
                }
                return true;
            })
    ],

    changePassword: [
        body('currentPassword')
            .notEmpty()
            .withMessage('Current password is required'),

        body('newPassword')
            .notEmpty()
            .withMessage('New password is required')
            .isLength({min: 6})
            .withMessage('New password must be at least 6 characters')
            .matches(/^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)/)
            .withMessage('New password must contain at least one uppercase letter, one lowercase letter, and one number')
            .custom((value, {req}) => {
                if (value === req.body.currentPassword) {
                    throw new Error('New password cannot be the same as current password');
                }
                return true;
            }),

        body('passwordConfirm')
            .notEmpty()
            .withMessage('Password confirmation is required')
            .custom((value, {req}) => {
                if (value !== req.body.newPassword) {
                    throw new Error('Password confirmation does not match new password');
                }
                return true;
            })
    ],

    updateProfile: [
        body('name')
            .optional()
            .trim()
            .isLength({min: 2, max: 50})
            .withMessage('Name must be between 2 and 50 characters'),

        body('email')
            .optional()
            .trim()
            .isEmail()
            .withMessage('Please provide a valid email')
            .normalizeEmail()
    ],

    updateRole: [
        body('roleId')
            .notEmpty()
            .withMessage('Role ID is required')
            .isInt()
            .withMessage('Role ID must be an integer')
    ]
};

module.exports = {authValidation};
