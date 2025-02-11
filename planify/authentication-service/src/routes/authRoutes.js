const express = require('express');
const PublicAuthController = require('../domains/auth/controllers/publicAuthController');
const ProtectedAuthController = require('../domains/auth/controllers/protectedAuthController');
const InternalAuthController = require('../domains/auth/controllers/internalAuthController');
const authMiddleware = require('../middlewares/authMiddleware');
const validate = require('../middlewares/validateMiddleware');
const { authValidation } = require('../validations/authValidation');

const router = express.Router();

// Public routes
router.post('/public/v1/register', validate(authValidation.register), PublicAuthController.register);
router.post('/public/v1/login', validate(authValidation.login), PublicAuthController.login);
router.post('/public/v1/forgot-password', validate(authValidation.forgotPassword), PublicAuthController.forgotPassword);
router.post('/public/v1/reset-password', validate(authValidation.resetPassword), PublicAuthController.resetPassword);
router.post('/public/v1/verify-email', PublicAuthController.verifyEmail);

// Protected routes (require authentication)
router.use(authMiddleware.authenticate);
router.post('/protected/v1/logout', ProtectedAuthController.logout);
router.post('/protected/v1/refresh-token', ProtectedAuthController.refresh);
router.post('/protected/v1/change-password', validate(authValidation.changePassword), ProtectedAuthController.changePassword);
router.get('/protected/v1/me', ProtectedAuthController.getCurrentUser);
router.put('/protected/v1/me', validate(authValidation.updateProfile), ProtectedAuthController.updateProfile);

// Internal routes (require admin role)
router.use(authMiddleware.requireRole('ADMIN'));
router.get('/internal/v1/users', InternalAuthController.getAllUsers);
router.get('/internal/v1/users/:id', InternalAuthController.getUserById);
router.put('/internal/v1/users/:id/role', validate(authValidation.updateRole), InternalAuthController.updateUserRole);
router.delete('/internal/v1/users/:id', InternalAuthController.deleteUser);

module.exports = router;

