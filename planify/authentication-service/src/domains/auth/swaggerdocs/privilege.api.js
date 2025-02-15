/**
 * @swagger
 * components:
 *   schemas:
 *     Privilege:
 *       type: object
 *       properties:
 *         id:
 *           type: string
 *           format: uuid
 *           description: The privilege unique identifier
 *         name:
 *           type: string
 *           description: The name of the privilege (uppercase)
 *         description:
 *           type: string
 *           description: Description of the privilege
 *         module:
 *           type: string
 *           description: The module this privilege belongs to
 *         isActive:
 *           type: boolean
 *           description: Whether the privilege is active
 *         createdAt:
 *           type: string
 *           format: date-time
 *         updatedAt:
 *           type: string
 *           format: date-time
 *       example:
 *         id: "550e8400-e29b-41d4-a716-446655440000"
 *         name: "CREATE_USER"
 *         description: "Allows creating new users"
 *         module: "USER_MANAGEMENT"
 *         isActive: true
 *         createdAt: "2024-02-15T12:00:00Z"
 *         updatedAt: "2024-02-15T12:00:00Z"
 *
 *     CreatePrivilegeRequest:
 *       type: object
 *       required:
 *         - name
 *         - module
 *       properties:
 *         name:
 *           type: string
 *           description: Privilege name (uppercase)
 *         description:
 *           type: string
 *           description: Privilege description
 *         module:
 *           type: string
 *           description: Module name (uppercase)
 *         isActive:
 *           type: boolean
 *           default: true
 *       example:
 *         name: "CREATE_USER"
 *         description: "Allows creating new users"
 *         module: "USER_MANAGEMENT"
 *
 *     UpdatePrivilegeRequest:
 *       type: object
 *       properties:
 *         name:
 *           type: string
 *           description: Updated privilege name
 *         description:
 *           type: string
 *           description: Updated description
 *         module:
 *           type: string
 *           description: Updated module name
 *         isActive:
 *           type: boolean
 *       example:
 *         description: "Updated description for user creation"
 *         isActive: false
 *
 * @swagger
 * /internal/v1/privileges:
 *   post:
 *     summary: Create a new privilege
 *     tags: [Privileges]
 *     security:
 *       - basicAuth: []
 *     requestBody:
 *       required: true
 *       content:
 *         application/json:
 *           schema:
 *             $ref: '#/components/schemas/CreatePrivilegeRequest'
 *     responses:
 *       201:
 *         description: Privilege created successfully
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
 *                     privilege:
 *                       $ref: '#/components/schemas/Privilege'
 *       400:
 *         description: Invalid input data
 *       401:
 *         description: Unauthorized
 *       403:
 *         description: Forbidden - Insufficient permissions
 *
 *   get:
 *     summary: List all privileges
 *     tags: [Privileges]
 *     security:
 *       - basicAuth: []
 *     parameters:
 *       - in: query
 *         name: search
 *         schema:
 *           type: string
 *         description: Search term for name or module
 *       - in: query
 *         name: module
 *         schema:
 *           type: string
 *         description: Filter by module name
 *       - in: query
 *         name: isActive
 *         schema:
 *           type: boolean
 *         description: Filter by active status
 *       - in: query
 *         name: page
 *         schema:
 *           type: integer
 *           minimum: 1
 *           default: 1
 *         description: Page number
 *       - in: query
 *         name: limit
 *         schema:
 *           type: integer
 *           minimum: 1
 *           maximum: 100
 *           default: 10
 *         description: Items per page
 *       - in: query
 *         name: sortBy
 *         schema:
 *           type: string
 *           enum: [name, module, createdAt, updatedAt]
 *         description: Field to sort by
 *       - in: query
 *         name: sortOrder
 *         schema:
 *           type: string
 *           enum: [ASC, DESC]
 *         description: Sort order
 *     responses:
 *       200:
 *         description: List of privileges
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
 *                     items:
 *                       type: array
 *                       items:
 *                         $ref: '#/components/schemas/Privilege'
 *                     meta:
 *                       type: object
 *                       properties:
 *                         totalItems:
 *                           type: integer
 *                         itemsPerPage:
 *                           type: integer
 *                         totalPages:
 *                           type: integer
 *                         currentPage:
 *                           type: integer
 *
 * @swagger
 * /internal/v1/privileges/{id}:
 *   get:
 *     summary: Get privilege by ID
 *     tags: [Privileges]
 *     security:
 *       - basicAuth: []
 *     parameters:
 *       - in: path
 *         name: id
 *         required: true
 *         schema:
 *           type: string
 *           format: uuid
 *         description: Privilege ID
 *     responses:
 *       200:
 *         description: Privilege details
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
 *                     privilege:
 *                       $ref: '#/components/schemas/Privilege'
 *       404:
 *         description: Privilege not found
 *
 *   put:
 *     summary: Update privilege
 *     tags: [Privileges]
 *     security:
 *       - basicAuth: []
 *     parameters:
 *       - in: path
 *         name: id
 *         required: true
 *         schema:
 *           type: string
 *           format: uuid
 *         description: Privilege ID
 *     requestBody:
 *       required: true
 *       content:
 *         application/json:
 *           schema:
 *             $ref: '#/components/schemas/UpdatePrivilegeRequest'
 *     responses:
 *       200:
 *         description: Privilege updated successfully
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
 *                     privilege:
 *                       $ref: '#/components/schemas/Privilege'
 *       404:
 *         description: Privilege not found
 *
 *   delete:
 *     summary: Delete privilege
 *     tags: [Privileges]
 *     security:
 *       - basicAuth: []
 *     parameters:
 *       - in: path
 *         name: id
 *         required: true
 *         schema:
 *           type: string
 *           format: uuid
 *         description: Privilege ID
 *     responses:
 *       204:
 *         description: Privilege deleted successfully
 *       404:
 *         description: Privilege not found
 *       400:
 *         description: Cannot delete privilege that is assigned to roles
 */
