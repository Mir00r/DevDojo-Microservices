'use strict';

/** @type {import('sequelize-cli').Migration} */
module.exports = {
    async up(queryInterface, Sequelize) {
        // First create the auth schema if it doesn't exist
        await queryInterface.sequelize.query('CREATE SCHEMA IF NOT EXISTS auth;');

        await queryInterface.createTable({
            tableName: 'roles',
            schema: 'auth'
        }, {
            id: {
                type: Sequelize.INTEGER,
                allowNull: false,
                autoIncrement: true,
                primaryKey: true
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

        // Add default roles
        await queryInterface.bulkInsert({
            tableName: 'roles',
            schema: 'auth'
        }, [
            {
                name: 'ADMIN',
                description: 'Administrator role',
                createdAt: new Date(),
                updatedAt: new Date()
            },
            {
                name: 'USER',
                description: 'Regular user role',
                createdAt: new Date(),
                updatedAt: new Date()
            }
        ]);
    },

    async down(queryInterface, Sequelize) {
        await queryInterface.dropTable({
            tableName: 'roles',
            schema: 'auth'
        });
    }
};
