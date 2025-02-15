'use strict';
const {Model} = require('sequelize');

module.exports = (sequelize, DataTypes) => {
    class User extends Model {
        /**
         * Helper method for defining associations.
         * This method is not a part of Sequelize lifecycle.
         * The `models/index` file will call this method automatically.
         */
        static associate(models) {
            // define association here
            // Define association with Role model
            User.belongsTo(models.Role, {
                foreignKey: 'roleId',
                targetKey: 'id',
                as: 'role'
            });

            // Add MFA association
            User.hasOne(models.MfaSecret, {
                foreignKey: 'userId',
                as: 'mfaSecret'
            });
        }
    }

    User.init({
        id: {
            type: DataTypes.UUID,
            defaultValue: DataTypes.UUIDV4,
            primaryKey: true,
            allowNull: false
        },
        name: {
            type: DataTypes.STRING,
            allowNull: false,
            validate: {
                notEmpty: true,
                len: [2, 100]
            }
        },
        email: {
            type: DataTypes.STRING,
            allowNull: false,
            unique: true,
            validate: {
                isEmail: true,
                notEmpty: true
            }
        },
        password: {
            type: DataTypes.STRING,
            allowNull: false,
            validate: {
                len: [6, 255]
            }
        },
        roleId: {
            type: DataTypes.INTEGER,
            allowNull: false,
            references: {
                model: 'roles',
                key: 'id'
            }
        },
        emailVerified: {  // Add this field
            type: DataTypes.BOOLEAN,
            allowNull: false,
            defaultValue: false
        },
        emailVerifiedAt: {
            type: DataTypes.DATE,
            allowNull: true
        },
        lastLogin: {
            type: DataTypes.DATE,
            allowNull: true
        },
        isActive: {
            type: DataTypes.BOOLEAN,
            defaultValue: true
        },
        // Add these fields to the User model
        isMfaEnabled: {
            type: DataTypes.BOOLEAN,
            defaultValue: false
        }
    }, {
        sequelize,
        modelName: 'User',
        tableName: 'users',
        schema: 'auth',
        timestamps: true,
        indexes: [
            {
                unique: true,
                fields: ['email']
            }
        ]
    });

    return User;
};
