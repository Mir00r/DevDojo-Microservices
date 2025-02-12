/**
 * @swagger
 * /public/v1/login:
 *   post:
 *     tags: [Public Auth]
 *     summary: Authenticate user
 *     description: Login with email and password to receive JWT tokens
 *     requestBody:
 *       required: true
 *       content:
 *         application/json:
 *           schema:
 *             $ref: '#/components/schemas/LoginRequestDto'
 *     responses:
 *       200:
 *         description: Login successful
 *         content:
 *           application/json:
 *             schema:
 *               $ref: '#/components/schemas/LoginResponseDto'
 *       400:
 *         description: Validation error - Missing or invalid input parameters
 *         content:
 *           application/json:
 *             schema:
 *               $ref: '#/components/schemas/ApiResponse'
 *             example:
 *               status: error
 *               message: Validation failed
 *               errors:
 *                 - field: email
 *                   message: Email is required
 *       401:
 *         description: Authentication failed - Invalid credentials or unverified email
 *         content:
 *           application/json:
 *             schema:
 *               $ref: '#/components/schemas/ApiResponse'
 *             example:
 *               status: error
 *               message: Invalid email or password
 *       403:
 *         description: Account disabled or locked
 *         content:
 *           application/json:
 *             schema:
 *               $ref: '#/components/schemas/ApiResponse'
 *             example:
 *               status: error
 *               message: Account is disabled
 *       422:
 *         description: Unverified email
 *         content:
 *           application/json:
 *             schema:
 *               $ref: '#/components/schemas/ApiResponse'
 *             example:
 *               status: error
 *               message: Please verify your email first
 *       429:
 *         description: Too many login attempts
 *         content:
 *           application/json:
 *             schema:
 *               $ref: '#/components/schemas/ApiResponse'
 *             example:
 *               status: error
 *               message: Too many login attempts. Please try again later
 *       500:
 *         description: Internal server error
 *         content:
 *           application/json:
 *             schema:
 *               $ref: '#/components/schemas/ApiResponse'
 *             example:
 *               status: error
 *               message: Internal server error
 */
