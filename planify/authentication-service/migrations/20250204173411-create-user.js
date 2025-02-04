'use strict';
/** @type {import('sequelize-cli').Migration} */
module.exports = {
    up: async (queryInterface, Sequelize) => {
        await queryInterface.createTable(
            {schema: "auth", tableName: "users"},
            {
                id: {
                    type: Sequelize.UUID,
                    defaultValue: Sequelize.UUIDV4,
                    allowNull: false,
                    primaryKey: true,
                },
                name: {
                    type: Sequelize.STRING,
                    allowNull: false,
                },
                email: {
                    type: Sequelize.STRING,
                    unique: true,
                    allowNull: false,
                },
                password: {
                    type: Sequelize.STRING,
                    allowNull: false,
                },
                roleId: {
                    type: Sequelize.INTEGER,
                    references: {
                        model: {schema: "auth", tableName: "roles"},
                        key: "id",
                    },
                    onUpdate: "CASCADE",
                    onDelete: "SET NULL",
                },
                createdAt: {
                    type: Sequelize.DATE,
                    defaultValue: Sequelize.literal("CURRENT_TIMESTAMP"),
                },
                updatedAt: {
                    type: Sequelize.DATE,
                    defaultValue: Sequelize.literal("CURRENT_TIMESTAMP"),
                },
            }
        );
    },
    down: async (queryInterface) => {
        await queryInterface.dropTable({schema: "auth", tableName: "users"});
    },
};
