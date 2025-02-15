const express = require('express');
const router = express.Router();
const privilegeController = require('../controllers/privilegeController');
const {authenticate} = require('../../../middlewares/authMiddleware');
const {authorize} = require('../../../middlewares/authorizationMiddleware');

router
    .route('/')
    .post(
        authenticate,
        authorize('CREATE_PRIVILEGE'),
        privilegeController.createPrivilege
    )
    .get(
        authenticate,
        authorize('VIEW_PRIVILEGES'),
        privilegeController.getAllPrivileges
    );

router
    .route('/:id')
    .get(
        authenticate,
        authorize('VIEW_PRIVILEGE'),
        privilegeController.getPrivilegeById
    )
    .put(
        authenticate,
        authorize('UPDATE_PRIVILEGE'),
        privilegeController.updatePrivilege
    )
    .delete(
        authenticate,
        authorize('DELETE_PRIVILEGE'),
        privilegeController.deletePrivilege
    );

module.exports = router;