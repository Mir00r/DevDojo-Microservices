const swaggerJsdoc = require('swagger-jsdoc');
const swaggerUi = require('swagger-ui-express');

const swaggerDefinition = {
    openapi: '3.0.0',
    info: {
        title: 'Authentication Service API Documentation',
        version: '1.0.0',
        description: 'API documentation for the Authentication and User Management Service',
        contact: {
            name: 'Development Team',
            email: 'devdojo@gmail.com'
        }
    },
    servers: [
        {
            url: 'http://localhost:5000/api/auth',
            description: 'Development server'
        }
    ],
    tags: [
        {
            name: 'Public Auth',
            description: 'Authentication endpoints accessible to the public'
        },
        {
            name: 'Protected Auth',
            description: 'Authentication endpoints requiring JWT authentication'
        },
        {
            name: 'Internal Auth',
            description: 'Internal authentication endpoints requiring admin access'
        },
        {
            name: 'Role Management',
            description: 'Role management endpoints'
        }
    ],
    components: {
        securitySchemes: {
            bearerAuth: {
                type: 'http',
                scheme: 'bearer',
                bearerFormat: 'JWT',
                description: 'Enter JWT token'
            },
            basicAuth: {
                type: 'http',
                scheme: 'basic',
                description: 'Basic authentication for internal APIs'
            }
        },
        schemas: {
            Error: {
                type: 'object',
                properties: {
                    status: {
                        type: 'string',
                        example: 'error'
                    },
                    message: {
                        type: 'string',
                        example: 'Error message'
                    },
                    errors: {
                        type: 'array',
                        items: {
                            type: 'object',
                            properties: {
                                field: {
                                    type: 'string'
                                },
                                message: {
                                    type: 'string'
                                }
                            }
                        }
                    }
                }
            }
        }
    }
};

const options = {
    swaggerDefinition,
    apis: ['./src/domains/auth/swaggerdocs/*.js']
};

const swaggerSpec = swaggerJsdoc(options);

const swaggerUiOptions = {
    explorer: true,
    customCss: '.swaggerdocs-ui .topbar { display: none }',
    swaggerOptions: {
        docExpansion: 'none',
        persistAuthorization: true,
        displayRequestDuration: true
    }
};

module.exports = {
    serve: swaggerUi.serve,
    setup: swaggerUi.setup(swaggerSpec, swaggerUiOptions)
};
