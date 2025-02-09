'use strict';

/** @type {import('sequelize-cli').Migration} */
module.exports = {
    async up(queryInterface, Sequelize) {
        await queryInterface.createTable({
            tableName: 'users',
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
                allowNull: false
            },
            email: {
                type: Sequelize.STRING,
                allowNull: false,
                unique: true
            },
            password: {
                type: Sequelize.STRING,
                allowNull: false
            },
            emailVerified: {
                type: Sequelize.BOOLEAN,
                defaultValue: false,
                allowNull: false
            },
            roleId: {
                type: Sequelize.INTEGER,
                allowNull: false,
                references: {
                    model: {
                        tableName: 'roles',
                        schema: 'auth'
                    },
                    key: 'id'
                },
                onUpdate: 'CASCADE',
                onDelete: 'RESTRICT'
            },
            lastLogin: {
                type: Sequelize.DATE,
                allowNull: true
            },
            isActive: {
                type: Sequelize.BOOLEAN,
                defaultValue: true
            },
            emailVerifiedAt: {
                allowNull: true,
                type: Sequelize.DATE
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

        // Add index for email
        await queryInterface.addIndex({
            tableName: 'users',
            schema: 'auth'
        }, ['email'], {
            unique: true,
            name: 'users_email_idx'
        });
    },

    async down(queryInterface, Sequelize) {
        await queryInterface.dropTable({
            tableName: 'users',
            schema: 'auth'
        });
    }
};
