const express = require('express');
const PrivilegeController = require('../controllers/privilegeController');
const {authenticate} = require('../../../middlewares/authMiddleware');
const {authorize} = require('../../../middlewares/authorizationMiddleware');

const router = express.Router();

/**
 * @swagger
 * tags:
 *   name: Privileges
 *   description: Privilege management endpoints
 */

/**
 * @swagger
 * /api/v1/auth/privileges:
 *   post:
 *     summary: Create a new privilege
 *     tags: [Privileges]
 *     security:
 *       - bearerAuth: []
 *     requestBody:
 *       required: true
 *       content:
 *         application/json:
 *           schema:
 *             type: object
 *             required:
 *               - name
 *               - module
 *             properties:
 *               name:
 *                 type: string
 *               description:
 *                 type: string
 *               module:
 *                 type: string
 *               isActive:
 *                 type: boolean
 *     responses:
 *       201:
 *         description: Privilege created successfully
 *       400:
 *         description: Invalid request data
 *       401:
 *         description: Unauthorized
 *       403:
 *         description: Forbidden
 */
router.post('/', authenticate, authorize('MANAGE_PRIVILEGES'), PrivilegeController.createPrivilege);

/**
 * @swagger
 * /api/v1/auth/privileges:
 *   get:
 *     summary: Get all privileges
 *     tags: [Privileges]
 *     security:
 *       - bearerAuth: []
 *     parameters:
 *       - in: query
 *         name: search
 *         schema:
 *           type: string
 *       - in: query
 *         name: module
 *         schema:
 *           type: string
 *       - in: query
 *         name: isActive
 *         schema:
 *           type: boolean
 *       - in: query
 *         name: page
 *         schema:
 *           type: integer
 *       - in: query
 *         name: limit
 *         schema:
 *           type: integer
 *       - in: query
 *         name: sortBy
 *         schema:
 *           type: string
 *       - in: query
 *         name: sortOrder
 *         schema:
 *           type: string
 *           enum: [ASC, DESC]
 *     responses:
 *       200:
 *         description: List of privileges
 *       401:
 *         description: Unauthorized
 *       403:
 *         description: Forbidden
 */
router.get('/', authenticate, authorize('VIEW_PRIVILEGES'), PrivilegeController.getAllPrivileges);

/**
 * @swagger
 * /api/v1/auth/privileges/{id}:
 *   get:
 *     summary: Get a privilege by ID
 *     tags: [Privileges]
 *     security:
 *       - bearerAuth: []
 *     parameters:
 *       - in: path
 *         name: id
 *         required: true
 *         schema:
 *           type: string
 *     responses:
 *       200:
 *         description: Privilege details
 *       401:
 *         description: Unauthorized
 *       403:
 *         description: Forbidden
 *       404:
 *         description: Privilege not found
 */
router.get('/:id', authenticate, authorize('VIEW_PRIVILEGES'), PrivilegeController.getPrivilegeById);

/**
 * @swagger
 * /api/v1/auth/privileges/{id}:
 *   patch:
 *     summary: Update a privilege
 *     tags: [Privileges]
 *     security:
 *       - bearerAuth: []
 *     parameters:
 *       - in: path
 *         name: id
 *         required: true
 *         schema:
 *           type: string
 *     requestBody:
 *       required: true
 *       content:
 *         application/json:
 *           schema:
 *             type: object
 *             properties:
 *               name:
 *                 type: string
 *               description:
 *                 type: string
 *               module:
 *                 type: string
 *               isActive:
 *                 type: boolean
 *     responses:
 *       200:
 *         description: Privilege updated successfully
 *       400:
 *         description: Invalid request data
 *       401:
 *         description: Unauthorized
 *       403:
 *         description: Forbidden
 *       404:
 *         description: Privilege not found
 */
router.patch('/:id', authenticate, authorize('MANAGE_PRIVILEGES'), PrivilegeController.updatePrivilege);

/**
 * @swagger
 * /api/v1/auth/privileges/{id}:
 *   delete:
 *     summary: Delete a privilege
 *     tags: [Privileges]
 *     security:
 *       - bearerAuth: []
 *     parameters:
 *       - in: path
 *         name: id
 *         required: true
 *         schema:
 *           type: string
 *     responses:
 *       204:
 *         description: Privilege deleted successfully
 *       401:
 *         description: Unauthorized
 *       403:
 *         description: Forbidden
 *       404:
 *         description: Privilege not found
 */
router.delete('/:id', authenticate, authorize('MANAGE_PRIVILEGES'), PrivilegeController.deletePrivilege);

module.exports = router;