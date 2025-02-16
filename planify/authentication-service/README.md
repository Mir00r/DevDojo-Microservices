# Authentication Service

A robust authentication and user management service built with Node.js and Express.js, providing secure user authentication, authorization, and account management functionalities.

## 🛠️ Technologies Used

| Technology | Purpose |
|------------|----------|
| **Node.js + Express.js** | Backend framework for building RESTful APIs |
| **PostgreSQL** | Primary database for storing user data and authentication records |
| **Sequelize ORM** | Object-Relational Mapping for database interactions |
| **JWT (jsonwebtoken)** | Token-based authentication and authorization |
| **bcryptjs** | Secure password hashing |
| **Swagger UI Express** | API documentation and testing interface |
| **Express Validator** | Request validation and sanitization |
| **Nodemailer** | Email service integration for notifications |
| **Morgan** | HTTP request logging |
| **Helmet** | Security middleware for HTTP headers |
| **Jest** | Unit and integration testing |
| **Cors** | Cross-Origin Resource Sharing middleware |
| **dotenv** | Environment variable management |

## 📁 Folder Structure

```
├── config/                 # Configuration files
│   ├── config.json        # Database configuration
│   └── swagger.js         # Swagger documentation setup
├── migrations/            # Database migrations
├── models/                # Sequelize models
├── src/
│   ├── app.js            # Express application setup
│   ├── server.js         # Server initialization
│   ├── configs/          # Application configurations
│   ├── domains/          # Business logic modules
│   │   └── auth/         # Authentication domain
│   ├── middlewares/      # Custom middleware
│   ├── routes/           # API routes
│   ├── utils/            # Utility functions
│   └── validations/      # Request validators
├── tests/                # Test files
└── scripts/              # Utility scripts
```

## 🚀 Installation & Setup

1. Clone the repository:
```bash
git clone <repository-url>
cd authentication-service
```

2. Install dependencies:
```bash
npm install
```

3. Configure environment variables:
```bash
cp .env.example .env
# Edit .env with your configuration
```

4. Set up the database:
```bash
# Update config/config.json with your database credentials
npx sequelize-cli db:migrate
npx sequelize-cli db:seed:all
```

5. Start the service:
```bash
npm start
```

## 📦 Key Dependencies

- **express**: Web framework for Node.js
- **sequelize**: ORM for database operations
- **jsonwebtoken**: JWT implementation for secure authentication
- **bcryptjs**: Password hashing library
- **nodemailer**: Email sending functionality
- **swagger-jsdoc & swagger-ui-express**: API documentation
- **class-validator**: Input validation
- **morgan**: HTTP request logger
- **helmet**: Security headers middleware
- **cors**: Cross-Origin Resource Sharing
- **jest & supertest**: Testing framework

## 🔐 API Endpoints

### Public Routes

- **POST /api/auth/register**
  - Register a new user
  - Request: `{ name, email, password }`

- **POST /api/auth/login**
  - Authenticate user
  - Request: `{ email, password }`

- **POST /api/auth/forgot-password**
  - Request password reset
  - Request: `{ email }`

### Protected Routes

- **GET /api/auth/protected/v1/me**
  - Get current user profile
  - Requires: Bearer token

- **PUT /api/auth/protected/v1/profile**
  - Update user profile
  - Requires: Bearer token
  - Request: `{ name, email }`

- **POST /api/auth/protected/v1/change-password**
  - Change password
  - Requires: Bearer token
  - Request: `{ currentPassword, newPassword }`

- **POST /api/auth/protected/v1/logout**
  - Logout current session
  - Requires: Bearer token

### Internal Routes

- **GET /api/auth/internal/users**
  - List all users (admin only)
  - Requires: Bearer token + Admin role

- **PUT /api/auth/internal/users/:userId/role**
  - Update user role (admin only)
  - Requires: Bearer token + Admin role

## 📚 API Documentation

Swagger documentation is available at:
```
http://localhost:5000/api-docs
```

## 🧪 Testing

Run tests:
```bash
npm test
```

Run tests in watch mode:
```bash
npm run test:watch
```

## 🔒 Security

- JWT-based authentication
- Password hashing with bcrypt
- Protected routes with middleware
- Request validation and sanitization
- Security headers with helmet
- CORS configuration

## 📝 License

ISC
