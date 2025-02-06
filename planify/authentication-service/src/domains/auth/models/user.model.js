const {DataTypes} = require('sequelize');
const sequelize = require('../../../configs/database').sequelize;

module.exports = (sequelize, DataTypes) => {
    const User = sequelize.define(
        "User",
        {
            id: {
                type: DataTypes.UUID,
                defaultValue: DataTypes.UUIDV4,
                primaryKey: true,
            },
            name: {
                type: DataTypes.STRING,
                allowNull: false,
            },
            email: {
                type: DataTypes.STRING,
                unique: true,
                allowNull: false,
            },
            password: {
                type: DataTypes.STRING,
                allowNull: false,
            },
            roleId: {
                type: DataTypes.INTEGER,
                allowNull: false,
            }
        },
        {schema: "auth", tableName: "users"}  // <-- Assign schema
    );

    User.associate = (models) => {
        User.belongsTo(models.Role, {
            foreignKey: "roleId",
            as: "role"
        });
    };

    return User;
};

