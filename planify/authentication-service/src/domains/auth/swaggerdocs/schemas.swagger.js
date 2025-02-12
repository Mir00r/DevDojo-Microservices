/**
 * @swagger
 * components:
 *   schemas:
 *     User:
 *       type: object
 *       properties:
 *         id:
 *           type: string
 *           format: uuid
 *           description: Unique identifier for the user
 *         name:
 *           type: string
 *           description: User's full name
 *         email:
 *           type: string
 *           format: email
 *           description: User's email address
 *         roleId:
 *           type: integer
 *           description: ID of the user's role
 *         emailVerified:
 *           type: boolean
 *           description: Email verification status
 *         lastLogin:
 *           type: string
 *           format: date-time
 *           description: Last login timestamp
 *         isActive:
 *           type: boolean
 *           description: User's active status
 *         role:
 *           $ref: '#/components/schemas/Role'
 *
 *     Role:
 *       type: object
 *       properties:
 *         id:
 *           type: integer
 *           description: Unique identifier for the role
 *         name:
 *           type: string
 *           description: Role name (e.g., ADMIN, USER)
 *         description:
 *           type: string
 *           description: Role description
 *         userCount:
 *           type: integer
 *           description: Number of users with this role
 *
 *     LoginRequestDto:
 *       type: object
 *       required:
 *         - email
 *         - password
 *       properties:
 *         email:
 *           type: string
 *           format: email
 *           example: john@example.com
 *         password:
 *           type: string
 *           format: password
 *           example: "YourPassword123!"
 *
 *     LoginResponseDto:
 *       type: object
 *       properties:
 *         status:
 *           type: string
 *           example: success
 *         data:
 *           type: object
 *           properties:
 *             user:
 *               $ref: '#/components/schemas/UserDto'
 *             accessToken:
 *               type: string
 *               example: "eyJhbGciOiJIUzI1NiIsInR5..."
 *             refreshToken:
 *               type: string
 *               example: "eyJhbGciOiJIUzI1NiIsInR5..."
 *
 *     UserDto:
 *       type: object
 *       properties:
 *         id:
 *           type: string
 *           format: uuid
 *         name:
 *           type: string
 *         email:
 *           type: string
 *         role:
 *           type: object
 *           properties:
 *             name:
 *               type: string
 *
 *     ApiResponse:
 *       type: object
 *       properties:
 *         status:
 *           type: string
 *           enum: [success, error]
 *         message:
 *           type: string
 *         data:
 *           type: object
 *
 *     PaginationParams:
 *       type: object
 *       properties:
 *         page:
 *           type: integer
 *           minimum: 1
 *           default: 1
 *         limit:
 *           type: integer
 *           minimum: 1
 *           maximum: 100
 *           default: 10
 *         search:
 *           type: string
 *         sortBy:
 *           type: string
 *         sortOrder:
 *           type: string
 *           enum: [ASC, DESC]
 */
