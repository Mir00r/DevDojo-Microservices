const {DataTypes} = require('sequelize');
const sequelize = require('../../../../authentication-service/src/configs/database');
const User = require('../../../../authentication-service/src/domains/auth/models/user.model');

const Review = sequelize.define('Review', {
    id: {type: DataTypes.UUID, defaultValue: DataTypes.UUIDV4, primaryKey: true},
    userId: {type: DataTypes.UUID, allowNull: false},
    content: {type: DataTypes.TEXT, allowNull: false},
    rating: {type: DataTypes.INTEGER, allowNull: false, validate: {min: 1, max: 5}}
}, {timestamps: true});

Review.belongsTo(User, {foreignKey: 'userId', onDelete: 'CASCADE'});

module.exports = Review;
