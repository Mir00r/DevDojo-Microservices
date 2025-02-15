'use strict';

module.exports = {
    async up(queryInterface, Sequelize) {
        // First create the auth schema if it doesn't exist
        await queryInterface.sequelize.query('CREATE SCHEMA IF NOT EXISTS auth;');

        // Create UUID extension if it doesn't exist
        try {
            await queryInterface.sequelize.query('CREATE EXTENSION IF NOT EXISTS "uuid-ossp";');
        } catch (error) {
            console.warn('Warning: Could not create UUID extension. Using Sequelize UUID generation instead.');
        }

        await queryInterface.createTable({
            tableName: 'privileges',
            schema: 'auth'
        }, {
            id: {
                type: Sequelize.UUID,
                defaultValue: Sequelize.UUIDV4,
                primaryKey: true,
                allowNull: false
            },
            name: {
                type: Sequelize.STRING,
                allowNull: false,
                unique: true
            },
            description: {
                type: Sequelize.STRING,
                allowNull: true
            },
            module: {
                type: Sequelize.STRING,
                allowNull: false
            },
            isActive: {
                type: Sequelize.BOOLEAN,
                defaultValue: true
            },
            createdAt: {
                allowNull: false,
                type: Sequelize.DATE,
                defaultValue: Sequelize.literal('CURRENT_TIMESTAMP')
            },
            updatedAt: {
                allowNull: false,
                type: Sequelize.DATE,
                defaultValue: Sequelize.literal('CURRENT_TIMESTAMP')
            }
        });

        // Add default privileges using Sequelize's UUID generation
        await queryInterface.bulkInsert({
            tableName: 'privileges',
            schema: 'auth'
        }, [
            {
                id: Sequelize.literal('gen_random_uuid()'), // PostgreSQL's built-in UUID generation
                name: 'MANAGE_PRIVILEGES',
                description: 'Can manage privileges',
                module: 'PRIVILEGE',
                isActive: true,
                createdAt: new Date(),
                updatedAt: new Date()
            },
            {
                id: Sequelize.literal('gen_random_uuid()'),
                name: 'VIEW_PRIVILEGES',
                description: 'Can view privileges',
                module: 'PRIVILEGE',
                isActive: true,
                createdAt: new Date(),
                updatedAt: new Date()
            }
        ]);
    },

    async down(queryInterface, Sequelize) {
        await queryInterface.dropTable({
            tableName: 'privileges',
            schema: 'auth'
        });
    }
};
