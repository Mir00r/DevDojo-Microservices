const request = require('supertest');
const sinon = require('sinon');
const app = require('../../../src/app');
const RoleService = require('../../../src/domains/auth/services/roleService');
const RoleRepository = require('../../../src/domains/auth/repositories/role.repository');
const {CreateRoleDto, UpdateRoleDto, RoleQueryDto, RoleResponseDto} = require('../../../src/domains/auth/dtos/role.dto');
const {AppError} = require('../../../src/utils/errorHandler');

/**
 * Test Configuration and Utilities
 */
const TEST_CONFIG = {
    INTERNAL_API: {
        USERNAME: process.env.INTERNAL_API_USERNAME || 'internal',
        PASSWORD: process.env.INTERNAL_API_PASSWORD || 'internal'
    },
    ADMIN_JWT: process.env.JWT_SECRET || 'test_admin_jwt_token',
    API_PREFIX: '/api/auth'  // Add API prefix to match app.js configuration
};

/**
 * Authentication utility functions for tests
 */
const AuthUtils = {
    /**
     * Generate Basic Auth header for internal API requests
     * @returns {string} Basic Auth header value
     */
    getBasicAuthHeader() {
        const credentials = `${TEST_CONFIG.INTERNAL_API.USERNAME}:${TEST_CONFIG.INTERNAL_API.PASSWORD}`;
        return `Basic ${Buffer.from(credentials).toString('base64')}`;
    },

    /**
     * Generate Admin JWT header for testing
     * @returns {string} Bearer token header value
     */
    getAdminJWTHeader() {
        return `Bearer ${TEST_CONFIG.ADMIN_JWT}`;
    }
};

describe('Role Management System Tests', () => {
    let sandbox;

    beforeEach(() => {
        sandbox = sinon.createSandbox();
    });

    afterEach(() => {
        sandbox.restore();
    });

    /**
     * Authentication Tests
     * Testing different authentication methods for internal APIs
     */
    describe('Authentication Tests', () => {
        describe('Basic Auth', () => {
            it('should accept valid basic auth credentials', async () => {
                const response = await request(app)
                    .get(`${TEST_CONFIG.API_PREFIX}/internal/v1/roles`)
                    .set('Authorization', AuthUtils.getBasicAuthHeader())
                    .expect(200);

                expect(response.body.data.roles).toBeDefined();
            });

            it('should reject invalid basic auth credentials', async () => {
                const invalidCredentials = `Basic ${Buffer.from('invalid:credentials').toString('base64')}`;
                await request(app)
                    .get(`${TEST_CONFIG.API_PREFIX}/internal/v1/roles`)
                    .set('Authorization', invalidCredentials)
                    .expect(401);
            });

            it('should reject requests without auth header', async () => {
                await request(app)
                    .get(`${TEST_CONFIG.API_PREFIX}/internal/v1/roles`)
                    .expect(401);
            });
        });

        describe('Admin JWT Auth', () => {
            it('should accept valid admin JWT', async () => {
                const response = await request(app)
                    .get(`${TEST_CONFIG.API_PREFIX}/internal/v1/roles`)
                    .set('Authorization', AuthUtils.getAdminJWTHeader())
                    .expect(200);

                expect(response.body.data.roles).toBeDefined();
            });

            it('should reject non-admin JWT', async () => {
                await request(app)
                    .get(`${TEST_CONFIG.API_PREFIX}/internal/v1/roles`)
                    .set('Authorization', 'Bearer user_jwt_token')
                    .expect(403);
            });
        });
    });

    /**
     * DTO Tests
     * Testing data transfer objects for proper validation and transformation
     */
    describe('DTO Tests', () => {
        describe('CreateRoleDto', () => {
            it('should create valid DTO with required fields', () => {
                const data = {
                    name: 'TEST_ROLE',
                    description: 'Test role description'
                };
                const dto = CreateRoleDto.from(data);
                expect(dto.name).toBe(data.name.trim());
                expect(dto.description).toBe(data.description.trim());
            });

            it('should throw error when name is missing', () => {
                expect(() => CreateRoleDto.from({})).toThrow('Role name is required');
            });

            it('should trim input fields', () => {
                const dto = CreateRoleDto.from({
                    name: '  TEST_ROLE  ',
                    description: '  description  '
                });
                expect(dto.name).toBe('TEST_ROLE');
                expect(dto.description).toBe('description');
            });
        });

        describe('UpdateRoleDto', () => {
            it('should create valid DTO with partial updates', () => {
                const data = {name: 'UPDATED_ROLE'};
                const dto = UpdateRoleDto.from(data);
                expect(dto.name).toBe('UPDATED_ROLE');
                expect(dto.description).toBeUndefined();
            });

            it('should throw error when no update fields provided', () => {
                expect(() => UpdateRoleDto.from({}))
                    .toThrow('At least one field must be provided for update');
            });
        });
    });

    /**
     * API Integration Tests with Authentication
     */
    describe('Role API Integration Tests', () => {
        describe('POST /api/auth/internal/v1/roles', () => {
            it('should create role with Basic Auth', async () => {
                const roleData = {
                    name: 'TEST_ROLE',
                    description: 'Test role'
                };

                const response = await request(app)
                    .post(`${TEST_CONFIG.API_PREFIX}/internal/v1/roles`)
                    .set('Authorization', AuthUtils.getBasicAuthHeader())
                    .send(roleData)
                    .expect(201);

                expect(response.body.data.role.name).toBe(roleData.name);
            });

            it('should validate role name format with Basic Auth', async () => {
                const response = await request(app)
                    .post(`${TEST_CONFIG.API_PREFIX}/internal/v1/roles`)
                    .set('Authorization', AuthUtils.getBasicAuthHeader())
                    .send({
                        name: '123_INVALID',
                        description: 'Invalid role name'
                    })
                    .expect(400);

                expect(response.body.errors).toBeDefined();
            });
        });

        describe('GET /api/auth/internal/v1/roles', () => {
            it('should return paginated roles list with Basic Auth', async () => {
                const response = await request(app)
                    .get(`${TEST_CONFIG.API_PREFIX}/internal/v1/roles`)
                    .set('Authorization', AuthUtils.getBasicAuthHeader())
                    .query({
                        page: 1,
                        limit: 10,
                        sortBy: 'name',
                        sortOrder: 'ASC'
                    })
                    .expect(200);

                expect(Array.isArray(response.body.data.roles.items)).toBeTruthy();
                expect(response.body.data.roles.meta.totalItems).toBeDefined();
            });
        });

        describe('GET /api/auth/internal/v1/roles/:id', () => {
            it('should get role by ID with Basic Auth', async () => {
                const roleId = 1;
                const response = await request(app)
                    .get(`${TEST_CONFIG.API_PREFIX}/internal/v1/roles/${roleId}`)
                    .set('Authorization', AuthUtils.getBasicAuthHeader())
                    .expect(200);

                expect(response.body.data.role.id).toBe(roleId);
            });

            it('should return 404 for non-existent role', async () => {
                await request(app)
                    .get(`${TEST_CONFIG.API_PREFIX}/internal/v1/roles/999`)
                    .set('Authorization', AuthUtils.getBasicAuthHeader())
                    .expect(404);
            });
        });

        describe('PUT /api/auth/internal/v1/roles/:id', () => {
            it('should update role with Basic Auth', async () => {
                const roleId = 1;
                const updateData = {
                    name: 'UPDATED_ROLE',
                    description: 'Updated description'
                };

                const response = await request(app)
                    .put(`${TEST_CONFIG.API_PREFIX}/internal/v1/roles/${roleId}`)
                    .set('Authorization', AuthUtils.getBasicAuthHeader())
                    .send(updateData)
                    .expect(200);

                expect(response.body.data.role.name).toBe(updateData.name);
            });
        });

        describe('DELETE /api/auth/internal/v1/roles/:id', () => {
            it('should delete role with Basic Auth', async () => {
                await request(app)
                    .delete(`${TEST_CONFIG.API_PREFIX}/internal/v1/roles/2`)
                    .set('Authorization', AuthUtils.getBasicAuthHeader())
                    .expect(204);
            });

            it('should prevent system role deletion', async () => {
                await request(app)
                    .delete(`${TEST_CONFIG.API_PREFIX}/internal/v1/roles/1`)
                    .set('Authorization', AuthUtils.getBasicAuthHeader())
                    .expect(403);
            });
        });
    });

    /**
     * Service Layer Tests
     */
    describe('RoleService Tests', () => {
        describe('createRole', () => {
            it('should create role successfully', async () => {
                const roleData = {
                    name: 'NEW_ROLE',
                    description: 'New role description'
                };

                sandbox.stub(RoleRepository, 'findByName').resolves(null);
                sandbox.stub(RoleRepository, 'create').resolves({
                    id: 1,
                    ...roleData,
                    createdAt: new Date(),
                    updatedAt: new Date()
                });

                const result = await RoleService.createRole(
                    roleData.name,
                    roleData.description
                );

                expect(result.name).toBe(roleData.name);
                expect(result.description).toBe(roleData.description);
            });

            it('should throw error for duplicate role name', async () => {
                sandbox.stub(RoleRepository, 'findByName').resolves({id: 1, name: 'EXISTING_ROLE'});

                await expect(RoleService.createRole('EXISTING_ROLE', 'Description'))
                    .rejects
                    .toThrow('Role with this name already exists');
            });
        });
    });

    /**
     * Error Handling Tests
     */
    describe('Error Handling Tests', () => {
        it('should handle internal server errors gracefully', async () => {
            sandbox.stub(RoleRepository, 'findAll').throws(new Error('Database connection failed'));

            const response = await request(app)
                .get(`${TEST_CONFIG.API_PREFIX}/internal/v1/roles`)
                .set('Authorization', AuthUtils.getBasicAuthHeader())
                .expect(500);

            expect(response.body.error).toBeDefined();
            expect(response.body.error.message).toBeDefined();
        });

        it('should handle validation errors properly', async () => {
            const response = await request(app)
                .post(`${TEST_CONFIG.API_PREFIX}/internal/v1/roles`)
                .set('Authorization', AuthUtils.getBasicAuthHeader())
                .send({
                    name: '',  // Empty name should trigger validation
                    description: 'Test role'
                })
                .expect(400);

            expect(response.body.errors).toBeDefined();
            expect(Array.isArray(response.body.errors)).toBeTruthy();
        });
    });
});
