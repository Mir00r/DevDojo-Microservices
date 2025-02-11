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

        await queryInterface.createTable('password_resets', {
            id: {
                type: Sequelize.UUID,
                defaultValue: Sequelize.UUIDV4,
                primaryKey: true
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
            token: {
                type: Sequelize.STRING,
                allowNull: false
            },
            expiresAt: {
                type: Sequelize.DATE,
                allowNull: false
            },
            isUsed: {
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
        await queryInterface.addIndex('password_resets', ['token'], {
            unique: true,
            schema: 'auth'
        });
        await queryInterface.addIndex('password_resets', ['userId'], {
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
            tableName: 'password_resets',
            schema: 'auth'
        });
    }
};
