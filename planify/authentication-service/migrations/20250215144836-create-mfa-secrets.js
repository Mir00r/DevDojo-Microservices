'use strict';

/** @type {import('sequelize-cli').Migration} */
module.exports = {
    async up(queryInterface, Sequelize) {
        /**
         * Add altering commands here.
         *
         * Example:
         * await queryInterface.createTable('users', { id: Sequelize.INTEGER });
         */

        // First ensure the auth schema exists
        await queryInterface.sequelize.query('CREATE SCHEMA IF NOT EXISTS auth;');

        // Create mfa_secrets table
        await queryInterface.createTable({
            tableName: 'mfa_secrets',
            schema: 'auth'
        }, {
            id: {
                type: Sequelize.UUID,
                defaultValue: Sequelize.UUIDV4,
                primaryKey: true,
                allowNull: false
            },
            userId: {
                type: Sequelize.UUID,
                allowNull: false,
                references: {
                    model: {
                        tableName: 'users',
                        schema: 'auth'
                    },
                    key: 'id'
                },
                onUpdate: 'CASCADE',
                onDelete: 'CASCADE'
            },
            secret: {
                type: Sequelize.STRING,
                allowNull: false
            },
            backupCodes: {
                type: Sequelize.ARRAY(Sequelize.STRING),
                defaultValue: [],
                allowNull: false
            },
            isEnabled: {
                type: Sequelize.BOOLEAN,
                defaultValue: false
            },
            lastUsedAt: {
                type: Sequelize.DATE,
                allowNull: true
            },
            verifiedAt: {
                type: Sequelize.DATE,
                allowNull: true
            },
            createdAt: {
                type: Sequelize.DATE,
                allowNull: false,
                defaultValue: Sequelize.literal('CURRENT_TIMESTAMP')
            },
            updatedAt: {
                type: Sequelize.DATE,
                allowNull: false,
                defaultValue: Sequelize.literal('CURRENT_TIMESTAMP')
            },
            deletedAt: {
                type: Sequelize.DATE,
                allowNull: true
            }
        });

        // Add indexes
        await queryInterface.addIndex({
            schema: 'auth',
            tableName: 'mfa_secrets'
        }, ['userId'], {
            unique: true,
            name: 'mfa_secrets_user_id_idx'
        });
    },

    async down(queryInterface, Sequelize) {
        /**
         * Add reverting commands here.
         *
         * Example:
         * await queryInterface.dropTable('users');
         */
        // Remove indexes first
        await queryInterface.removeIndex({
            schema: 'auth',
            tableName: 'mfa_secrets'
        }, 'mfa_secrets_user_id_idx');

        // Drop the table
        await queryInterface.dropTable({
            tableName: 'mfa_secrets',
            schema: 'auth'
        });
    }
};
