const express = require('express');
const PublicAuthController = require('../domains/auth/controllers/publicAuthController');
const ProtectedAuthController = require('../domains/auth/controllers/protectedAuthController');
const InternalAuthController = require('../domains/auth/controllers/internalAuthController');
const RoleController = require('../domains/auth/controllers/roleController');
const PrivilegeController = require('../domains/auth/controllers/privilegeController');
const MfaController = require('../domains/auth/controllers/mfaController');
const authMiddleware = require('../middlewares/authMiddleware');
const validate = require('../middlewares/validateMiddleware');
const {authValidation} = require('../validations/authValidation');
const {roleValidation} = require('../validations/roleValidation');
const {privilegeValidation} = require('../validations/privilegeValidation');
const {mfaValidation} = require('../validations/mfaValidation');


const router = express.Router();

// Public routes
router.post('/public/v1/register', validate(authValidation.register), PublicAuthController.register);
router.post('/public/v1/login', validate(authValidation.login), PublicAuthController.login);
router.post('/public/v1/forgot-password', validate(authValidation.forgotPassword), PublicAuthController.forgotPassword);
router.post('/public/v1/reset-password', validate(authValidation.resetPassword), PublicAuthController.resetPassword);
router.post('/public/v1/verify-email', PublicAuthController.verifyEmail);


// Internal routes with Basic Auth OR Admin JWT
router.use('/internal/v1', (req, res, next) => {
    // Check if Basic Auth is used
    const authHeader = req.headers.authorization;
    if (authHeader && authHeader.startsWith('Basic ')) {
        return authMiddleware.internalBasicAuth(req, res, next);
    }
    // If not Basic Auth, check for Admin JWT
    return authMiddleware.requireRole('ADMIN')(req, res, next);
});
// Internal routes
router.get('/internal/v1/users', InternalAuthController.getAllUsers);
router.get('/internal/v1/users/:id', InternalAuthController.getUserById);
router.put('/internal/v1/users/:id/role', validate(authValidation.updateRole), InternalAuthController.updateUserRole);
router.delete('/internal/v1/users/:id', InternalAuthController.deleteUser);

// ROLE Routes
router.post('/internal/v1/roles', validate(roleValidation.createRole), RoleController.createRole);
router.get('/internal/v1/roles', RoleController.getAllRoles);
router.get('/internal/v1/roles/:id', validate(roleValidation.checkId), RoleController.getRoleById);
router.put('/internal/v1/roles/:id', validate(roleValidation.updateRole), RoleController.updateRole);
router.delete('/internal/v1/roles/:id', validate(roleValidation.checkId), RoleController.deleteRole);

// PRIVILEGES Routes
router.post('/internal/v1/privileges', validate(privilegeValidation.createPrivilege), PrivilegeController.createPrivilege);
router.get('/internal/v1/privileges', PrivilegeController.getAllPrivileges);
router.get('/internal/v1/privileges/:id', validate(privilegeValidation.checkId), PrivilegeController.getPrivilegeById);
router.put('/internal/v1/privileges/:id', validate(privilegeValidation.updatePrivilege), PrivilegeController.updatePrivilege);
router.delete('/internal/v1/privileges/:id', validate(privilegeValidation.checkId), PrivilegeController.deletePrivilege);

// Protected routes (require authentication)
router.use(authMiddleware.authenticate);
router.post('/protected/v1/logout', ProtectedAuthController.logout);
router.post('/protected/v1/refresh-token', ProtectedAuthController.refresh);
router.post('/protected/v1/change-password', validate(authValidation.changePassword), ProtectedAuthController.changePassword);
router.get('/protected/v1/me', ProtectedAuthController.getCurrentUser);
router.put('/protected/v1/me', validate(authValidation.updateProfile), ProtectedAuthController.updateProfile);

// All MFA routes require authentication
router.use(authMiddleware.authenticate);

router.post('/protected/v1/mfa/enable', validate(mfaValidation.enable), MfaController.enableMfa);
router.post('/protected/v1/mfa/verify', validate(mfaValidation.verify), MfaController.verifyMfa);

module.exports = router;

