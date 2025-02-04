'use strict';
/** @type {import('sequelize-cli').Migration} */
module.exports = {
    up: async (queryInterface, Sequelize) => {
        await queryInterface.createTable(
            {schema: "auth", tableName: "roles"},
            {
                id: {
                    type: Sequelize.INTEGER,
                    autoIncrement: true,
                    primaryKey: true,
                },
                name: {
                    type: Sequelize.STRING,
                    unique: true,
                    allowNull: false,
                }
            }
        );
    },
    down: async (queryInterface) => {
        await queryInterface.dropTable({schema: "auth", tableName: "roles"});
    },
};
