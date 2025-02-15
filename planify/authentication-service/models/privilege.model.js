'use strict';
const {Model} = require('sequelize');

module.exports = (sequelize, DataTypes) => {
    class Privilege extends Model {
        static associate(models) {
            // Define associations here if needed
            // For example, many-to-many with roles
            Privilege.belongsToMany(models.Role, {
                through: 'RolePrivileges',
                foreignKey: 'privilegeId',
                as: 'roles'
            });
        }
    }

    Privilege.init({
        id: {
            type: DataTypes.UUID,
            defaultValue: DataTypes.UUIDV4,
            primaryKey: true,
            allowNull: false
        },
        name: {
            type: DataTypes.STRING,
            allowNull: false,
            unique: true,
            validate: {
                notEmpty: true
            }
        },
        description: {
            type: DataTypes.STRING,
            allowNull: true
        },
        module: {
            type: DataTypes.STRING,
            allowNull: false,
            validate: {
                notEmpty: true
            }
        },
        isActive: {
            type: DataTypes.BOOLEAN,
            defaultValue: true
        }
    }, {
        sequelize,
        modelName: 'Privilege',
        tableName: 'privileges',
        schema: 'auth',
        timestamps: true
    });

    return Privilege;
};
