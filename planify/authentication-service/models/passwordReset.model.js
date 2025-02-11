'use strict';
const {Model} = require('sequelize');

module.exports = (sequelize, DataTypes) => {
    class PasswordReset extends Model {
        static associate(models) {
            PasswordReset.belongsTo(models.User, {
                foreignKey: 'userId',
                as: 'user'
            });
        }
    }

    PasswordReset.init({
        id: {
            type: DataTypes.UUID,
            defaultValue: DataTypes.UUIDV4,
            primaryKey: true
        },
        userId: {
            type: DataTypes.UUID,
            allowNull: false,
            references: {
                model: 'users',
                key: 'id'
            }
        },
        token: {
            type: DataTypes.STRING,
            allowNull: false
        },
        expiresAt: {
            type: DataTypes.DATE,
            allowNull: false
        },
        isUsed: {
            type: DataTypes.BOOLEAN,
            defaultValue: false
        }
    }, {
        sequelize,
        modelName: 'PasswordReset',
        tableName: 'password_resets',
        schema: 'auth',
        indexes: [
            {
                unique: true,
                fields: ['token']
            }
        ]
    });

    return PasswordReset;
};
