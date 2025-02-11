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

        await queryInterface.createTable('refresh_tokens', {
            id: {
                type: Sequelize.UUID,
                defaultValue: Sequelize.UUIDV4,
                primaryKey: true
            },
            token: {
                type: Sequelize.STRING,
                allowNull: false,
                unique: true
            },
            userId: {
                type: Sequelize.UUID,
                allowNull: false,
                references: {
                    model: 'users',
                    key: 'id'
                },
                onUpdate: 'CASCADE',
                onDelete: 'CASCADE'
            },
            expiresAt: {
                type: Sequelize.DATE,
                allowNull: false
            },
            isRevoked: {
                type: Sequelize.BOOLEAN,
                defaultValue: false
            },
            createdAt: {
                allowNull: false,
                type: Sequelize.DATE
            },
            updatedAt: {
                allowNull: false,
                type: Sequelize.DATE
            }
        }, {
            schema: 'auth'
        });

        // Add indexes
        await queryInterface.addIndex('refresh_tokens', ['token'], {
            unique: true,
            schema: 'auth'
        });
        await queryInterface.addIndex('refresh_tokens', ['userId'], {
            schema: 'auth'
        });
    },

    async down(queryInterface, Sequelize) {
        /**
         * Add reverting commands here.
         *
         * Example:
         * await queryInterface.dropTable('users');
         */
        await queryInterface.dropTable({
            tableName: 'refresh_tokens',
            schema: 'auth'
        });
    }
};
