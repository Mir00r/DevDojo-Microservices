/**
 * @swagger
 * components:
 *   securitySchemes:
 *     bearerAuth:
 *       type: http
 *       scheme: bearer
 *       bearerFormat: JWT
 *
 *   schemas:
 *     EnableMfaDto:
 *       type: object
 *       required:
 *         - userId
 *       properties:
 *         userId:
 *           type: string
 *           format: uuid
 *           description: "User's unique identifier"
 *           example: "123e4567-e89b-12d3-a456-426614174000"
 *         phoneNumber:
 *           type: string
 *           description: "Optional: User's phone number for SMS-based MFA"
 *           example: "+1234567890"
 *         email:
 *           type: string
 *           format: email
 *           description: "Optional: User's email for email-based MFA"
 *           example: "user@example.com"
 *
 *     VerifyMfaDto:
 *       type: object
 *       required:
 *         - userId
 *         - code
 *       properties:
 *         userId:
 *           type: string
 *           format: uuid
 *           description: "User's unique identifier"
 *           example: "123e4567-e89b-12d3-a456-426614174000"
 *         code:
 *           type: string
 *           description: "Verification code"
 *           example: "123456"
 *
 *     MfaEnableResponse:
 *       type: object
 *       properties:
 *         status:
 *           type: string
 *           example: "success"
 *         message:
 *           type: string
 *           example: "MFA setup initiated successfully"
 *         data:
 *           type: object
 *           properties:
 *             secret:
 *               type: string
 *               example: "OQQWOURWJQSES5TVMVKHG3JWGRYUQQTXLIYFK4CHNMUVI2J4IM3Q"
 *             qrCode:
 *               type: string
 *               example: "data:image/png;base64,..."
 *             backupCodes:
 *               type: array
 *               items:
 *                 type: string
 *               example: ["F8BB1FDBAE", "1F5ADC4A84"]
 *
 *     MfaVerifyResponse:
 *       type: object
 *       properties:
 *         status:
 *           type: string
 *           example: "success"
 *         message:
 *           type: string
 *           example: "MFA verification successful"
 *         data:
 *           type: object
 *           properties:
 *             verified:
 *               type: boolean
 *               example: true
 *             message:
 *               type: string
 *               example: "Verification successful"
 *
 *     ApiResponse:
 *       type: object
 *       properties:
 *         status:
 *           type: string
 *           example: "error"
 *         message:
 *           type: string
 *           example: "Invalid verification code"
 *         errors:
 *           type: array
 *           items:
 *             type: object
 *             properties:
 *               field:
 *                 type: string
 *               message:
 *                 type: string
 *
 * tags:
 *   name: MFA
 *   description: Multi-Factor Authentication endpoints
 */

/**
 * @swagger
 * /protected/v1/mfa/enable:
 *   post:
 *     summary: Enable MFA for user
 *     tags: [MFA]
 *     security:
 *       - bearerAuth: []
 *     requestBody:
 *       required: true
 *       content:
 *         application/json:
 *           schema:
 *             $ref: '#/components/schemas/EnableMfaDto'
 *     responses:
 *       200:
 *         description: MFA setup initiated successfully
 *         content:
 *           application/json:
 *             schema:
 *               $ref: '#/components/schemas/MfaEnableResponse'
 *       401:
 *         description: Unauthorized - Invalid or missing token
 *         content:
 *           application/json:
 *             schema:
 *               $ref: '#/components/schemas/ApiResponse'
 *       400:
 *         description: Bad Request - Invalid input or MFA already enabled
 *         content:
 *           application/json:
 *             schema:
 *               $ref: '#/components/schemas/ApiResponse'
 *       500:
 *         description: Internal Server Error
 *         content:
 *           application/json:
 *             schema:
 *               $ref: '#/components/schemas/ApiResponse'
 *
 * /protected/v1/mfa/verify:
 *   post:
 *     summary: Verify MFA code
 *     tags: [MFA]
 *     security:
 *       - bearerAuth: []
 *     requestBody:
 *       required: true
 *       content:
 *         application/json:
 *           schema:
 *             $ref: '#/components/schemas/VerifyMfaDto'
 *     responses:
 *       200:
 *         description: MFA verification successful
 *         content:
 *           application/json:
 *             schema:
 *               $ref: '#/components/schemas/MfaVerifyResponse'
 *       400:
 *         description: Bad Request - Invalid code format or missing required fields
 *         content:
 *           application/json:
 *             schema:
 *               $ref: '#/components/schemas/ApiResponse'
 *       401:
 *         description: Unauthorized - Invalid or missing token
 *         content:
 *           application/json:
 *             schema:
 *               $ref: '#/components/schemas/ApiResponse'
 *       500:
 *         description: Internal Server Error
 *         content:
 *           application/json:
 *             schema:
 *               $ref: '#/components/schemas/ApiResponse'
 */
