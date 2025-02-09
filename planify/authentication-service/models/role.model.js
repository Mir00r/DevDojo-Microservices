'use strict';
const {Model} = require('sequelize');

module.exports = (sequelize, DataTypes) => {
    class Role extends Model {
        /**
         * Helper method for defining associations.
         * This method is not a part of Sequelize lifecycle.
         * The `models/index` file will call this method automatically.
         */
        static associate(models) {
            // define association here
            // Define association with User model
            Role.hasMany(models.User, {
                foreignKey: 'roleId',
                sourceKey: 'id',
                as: 'users'
            });
        }
    }

    Role.init({
        id: {
            type: DataTypes.INTEGER,
            autoIncrement: true,
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
        }
    }, {
        sequelize,
        modelName: 'Role',
        tableName: 'roles',
        schema: 'auth',
        timestamps: true,
        indexes: [
            {
                unique: true,
                fields: ['name']
            }
        ]
    });

    return Role;
};
