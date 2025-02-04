const {DataTypes} = require('sequelize');
const sequelize = require('../../../configs/database');

// const User = sequelize.define('User', {
//     id: {type: DataTypes.UUID, defaultValue: DataTypes.UUIDV4, primaryKey: true},
//     name: {type: DataTypes.STRING, allowNull: false},
//     email: {type: DataTypes.STRING, allowNull: false, unique: true},
//     password: {type: DataTypes.STRING, allowNull: false},
//     role: {type: DataTypes.ENUM('user', 'admin'), defaultValue: 'user'},
// }, {timestamps: true}, { schema: 'auth', tableName: 'Users' } );
//
// module.exports = User;


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
        {schema: "auth"}  // <-- Assign schema
    );

    User.associate = (models) => {
        User.belongsTo(models.Role, {foreignKey: "roleId"});
    };

    return User;
};

