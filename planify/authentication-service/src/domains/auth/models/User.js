const {DataTypes} = require('sequelize');
const sequelize = require('../../../configs/database');

const User = sequelize.define('User', {
    id: {type: DataTypes.UUID, defaultValue: DataTypes.UUIDV4, primaryKey: true},
    name: {type: DataTypes.STRING, allowNull: false},
    email: {type: DataTypes.STRING, allowNull: false, unique: true},
    password: {type: DataTypes.STRING, allowNull: false},
    role: {type: DataTypes.ENUM('user', 'admin'), defaultValue: 'user'}
}, {timestamps: true});

module.exports = User;
