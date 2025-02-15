'use strict';
const {Model} = require('sequelize');

module.exports = (sequelize, DataTypes) => {
    class MfaSecret extends Model {
        static associate(models) {
            // Define association with User model
            MfaSecret.belongsTo(models.User, {
                foreignKey: 'userId',
                as: 'user'
            });
        }
    }

    MfaSecret.init({
        id: {
            type: DataTypes.UUID,
            defaultValue: DataTypes.UUIDV4,
            primaryKey: true,
            allowNull: false
        },
        userId: {
            type: DataTypes.UUID,
            allowNull: false,
            references: {
                model: 'users',
                key: 'id'
            },
            onUpdate: 'CASCADE',
            onDelete: 'CASCADE'
        },
        secret: {
            type: DataTypes.STRING,
            allowNull: false
        },
        backupCodes: {
            type: DataTypes.ARRAY(DataTypes.STRING),
            defaultValue: [],
            allowNull: false
        },
        isEnabled: {
            type: DataTypes.BOOLEAN,
            defaultValue: false
        },
        lastUsedAt: {
            type: DataTypes.DATE,
            allowNull: true
        },
        verifiedAt: {
            type: DataTypes.DATE,
            allowNull: true
        }
    }, {
        sequelize,
        modelName: 'MfaSecret',
        tableName: 'mfa_secrets',
        schema: 'auth',
        timestamps: true,
        paranoid: true // Soft deletes
    });

    return MfaSecret;
};
