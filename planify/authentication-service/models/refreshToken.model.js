'use strict';
const {Model} = require('sequelize');

module.exports = (sequelize, DataTypes) => {
    class RefreshToken extends Model {
        static associate(models) {
            RefreshToken.belongsTo(models.User, {
                foreignKey: 'userId',
                as: 'user'
            });
        }
    }

    RefreshToken.init({
        id: {
            type: DataTypes.UUID,
            defaultValue: DataTypes.UUIDV4,
            primaryKey: true
        },
        token: {
            type: DataTypes.STRING,
            allowNull: false,
            unique: true
        },
        userId: {
            type: DataTypes.UUID,
            allowNull: false,
            references: {
                model: 'users',
                key: 'id'
            }
        },
        expiresAt: {
            type: DataTypes.DATE,
            allowNull: false
        },
        isRevoked: {
            type: DataTypes.BOOLEAN,
            defaultValue: false
        }
    }, {
        sequelize,
        modelName: 'RefreshToken',
        tableName: 'refresh_tokens',
        schema: 'auth'
    });

    return RefreshToken;
};
