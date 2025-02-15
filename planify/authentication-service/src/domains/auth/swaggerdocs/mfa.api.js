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
 *     MfaVerifyRequest:
 *       type: object
 *       required:
 *         - code
 *       properties:
 *         code:
 *           type: string
 *           description: "6-digit TOTP code or 10-character backup code"
 *           example: "123456"
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
 *
 *     ErrorResponse:
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
 * @swagger
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
 *               $ref: '#/components/schemas/ErrorResponse'
 *       400:
 *         description: Bad Request - MFA already enabled
 *         content:
 *           application/json:
 *             schema:
 *               $ref: '#/components/schemas/ErrorResponse'
 *       500:
 *         description: Internal Server Error
 *         content:
 *           application/json:
 *             schema:
 *               $ref: '#/components/schemas/ErrorResponse'
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
 *             $ref: '#/components/schemas/MfaVerifyRequest'
 *     responses:
 *       200:
 *         description: MFA verification successful
 *         content:
 *           application/json:
 *             schema:
 *               $ref: '#/components/schemas/MfaVerifyResponse'
 *       400:
 *         description: Bad Request - Invalid code format
 *         content:
 *           application/json:
 *             schema:
 *               $ref: '#/components/schemas/ErrorResponse'
 *       401:
 *         description: Unauthorized - Invalid or missing token
 *         content:
 *           application/json:
 *             schema:
 *               $ref: '#/components/schemas/ErrorResponse'
 *       500:
 *         description: Internal Server Error
 *         content:
 *           application/json:
 *             schema:
 *               $ref: '#/components/schemas/ErrorResponse'
 */
