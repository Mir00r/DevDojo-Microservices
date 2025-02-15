/**
 * @swagger
 * components:
 *   schemas:
 *     Role:
 *       type: object
 *       properties:
 *         id:
 *           type: integer
 *           description: The role ID
 *         name:
 *           type: string
 *           description: The role name
 *         description:
 *           type: string
 *           description: Role description
 *         userCount:
 *           type: integer
 *           description: Number of users with this role
 *         createdAt:
 *           type: string
 *           format: date-time
 *         updatedAt:
 *           type: string
 *           format: date-time
 *       example:
 *         id: 1
 *         name: MANAGER
 *         description: Manager role with elevated permissions
 *         userCount: 5
 *         createdAt: "2024-02-15T10:00:00.000Z"
 *         updatedAt: "2024-02-15T10:00:00.000Z"
 *
 *     CreateRoleRequest:
 *       type: object
 *       required:
 *         - name
 *       properties:
 *         name:
 *           type: string
 *           description: Role name (will be converted to uppercase)
 *         description:
 *           type: string
 *           description: Role description
 *       example:
 *         name: manager
 *         description: Manager role with elevated permissions
 *
 *     UpdateRoleRequest:
 *       type: object
 *       properties:
 *         name:
 *           type: string
 *           description: New role name
 *         description:
 *           type: string
 *           description: New role description
 *       example:
 *         name: senior_manager
 *         description: Updated manager role description
 *
 *     PaginatedRoleResponse:
 *       type: object
 *       properties:
 *         status:
 *           type: string
 *           example: success
 *         data:
 *           type: object
 *           properties:
 *             items:
 *               type: array
 *               items:
 *                 $ref: '#/components/schemas/Role'
 *             meta:
 *               type: object
 *               properties:
 *                 totalItems:
 *                   type: integer
 *                   example: 10
 *                 itemsPerPage:
 *                   type: integer
 *                   example: 10
 *                 totalPages:
 *                   type: integer
 *                   example: 1
 *                 currentPage:
 *                   type: integer
 *                   example: 1
 *
 *   securitySchemes:
 *     bearerAuth:
 *       type: http
 *       scheme: bearer
 *       bearerFormat: JWT
 *
 * @swagger
 * /internal/v1/roles:
 *   post:
 *     summary: Create a new role
 *     tags: [Roles]
 *     security:
 *       - basicAuth: []
 *     requestBody:
 *       required: true
 *       content:
 *         application/json:
 *           schema:
 *             $ref: '#/components/schemas/CreateRoleRequest'
 *     responses:
 *       201:
 *         description: Role created successfully
 *         content:
 *           application/json:
 *             schema:
 *               type: object
 *               properties:
 *                 status:
 *                   type: string
 *                   example: success
 *                 data:
 *                   type: object
 *                   properties:
 *                     role:
 *                       $ref: '#/components/schemas/Role'
 *       400:
 *         description: Bad request (e.g., role already exists)
 *       401:
 *         description: Unauthorized
 *       403:
 *         description: Forbidden (not admin)
 *
 *   get:
 *     summary: Get all roles with pagination
 *     tags: [Roles]
 *     security:
 *       - basicAuth: []
 *     parameters:
 *       - in: query
 *         name: page
 *         schema:
 *           type: integer
 *           default: 1
 *         description: Page number
 *       - in: query
 *         name: limit
 *         schema:
 *           type: integer
 *           default: 10
 *         description: Items per page
 *       - in: query
 *         name: search
 *         schema:
 *           type: string
 *         description: Search by role name
 *       - in: query
 *         name: sortBy
 *         schema:
 *           type: string
 *           default: createdAt
 *         description: Field to sort by
 *       - in: query
 *         name: sortOrder
 *         schema:
 *           type: string
 *           enum: [ASC, DESC]
 *           default: DESC
 *         description: Sort order
 *     responses:
 *       200:
 *         description: List of roles
 *         content:
 *           application/json:
 *             schema:
 *               $ref: '#/components/schemas/PaginatedRoleResponse'
 *       401:
 *         description: Unauthorized
 *       403:
 *         description: Forbidden
 *
 * @swagger
 * /internal/v1/roles/{id}:
 *   get:
 *     summary: Get role by ID
 *     tags: [Roles]
 *     security:
 *       - basicAuth: []
 *     parameters:
 *       - in: path
 *         name: id
 *         required: true
 *         schema:
 *           type: integer
 *         description: Role ID
 *     responses:
 *       200:
 *         description: Role details
 *         content:
 *           application/json:
 *             schema:
 *               type: object
 *               properties:
 *                 status:
 *                   type: string
 *                   example: success
 *                 data:
 *                   type: object
 *                   properties:
 *                     role:
 *                       $ref: '#/components/schemas/Role'
 *       404:
 *         description: Role not found
 *
 *   put:
 *     summary: Update role
 *     tags: [Roles]
 *     security:
 *       - basicAuth: []
 *     parameters:
 *       - in: path
 *         name: id
 *         required: true
 *         schema:
 *           type: integer
 *         description: Role ID
 *     requestBody:
 *       required: true
 *       content:
 *         application/json:
 *           schema:
 *             $ref: '#/components/schemas/UpdateRoleRequest'
 *     responses:
 *       200:
 *         description: Role updated successfully
 *         content:
 *           application/json:
 *             schema:
 *               type: object
 *               properties:
 *                 status:
 *                   type: string
 *                   example: success
 *                 data:
 *                   type: object
 *                   properties:
 *                     role:
 *                       $ref: '#/components/schemas/Role'
 *       400:
 *         description: Bad request
 *       403:
 *         description: Cannot modify system roles
 *       404:
 *         description: Role not found
 *
 *   delete:
 *     summary: Delete role
 *     tags: [Roles]
 *     security:
 *       - basicAuth: []
 *     parameters:
 *       - in: path
 *         name: id
 *         required: true
 *         schema:
 *           type: integer
 *         description: Role ID
 *     responses:
 *       204:
 *         description: Role deleted successfully
 *       400:
 *         description: Cannot delete role with assigned users
 *       403:
 *         description: Cannot delete system roles
 *       404:
 *         description: Role not found
 */
