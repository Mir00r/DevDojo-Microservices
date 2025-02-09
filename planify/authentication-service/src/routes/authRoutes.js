const express = require('express');
const AuthController = require('../domains/auth/controllers/authController');
const authMiddleware = require('../middlewares/authMiddleware');
const validate = require('../middlewares/validateMiddleware');
const { authValidation } = require('../validations/authValidation');

const router = express.Router();

// Public routes
router.post('/register', validate(authValidation.register), AuthController.register);
router.post('/login', validate(authValidation.login), AuthController.login);
router.post('/forgot-password', validate(authValidation.forgotPassword), AuthController.forgotPassword);
router.post('/reset-password', validate(authValidation.resetPassword), AuthController.resetPassword);
router.post('/verify-email', AuthController.verifyEmail);

// Protected routes (require authentication)
router.use(authMiddleware.authenticate);
router.post('/logout', AuthController.logout);
router.post('/change-password', validate(authValidation.changePassword), AuthController.changePassword);
router.get('/me', AuthController.getCurrentUser);
router.put('/me', validate(authValidation.updateProfile), AuthController.updateProfile);

// Internal routes (require admin role)
router.use(authMiddleware.requireRole('ADMIN'));
router.get('/users', AuthController.getAllUsers);
router.get('/users/:id', AuthController.getUserById);
router.put('/users/:id/role', validate(authValidation.updateRole), AuthController.updateUserRole);
router.delete('/users/:id', AuthController.deleteUser);

module.exports = router;

