const bcrypt = require('bcryptjs');
const jwt = require('jsonwebtoken');
// const User = require('../models/user.model');
const {sequelize} = require('../../../configs/database');
const User = require('../models/user.model')(sequelize, require('sequelize').DataTypes);


const generateToken = (user) => {
    return jwt.sign(
        {
            id: user.id, role: user.role
        }, process.env.JWT_SECRET,
        {expiresIn: '1h', algorithm: 'HS256'});
};

const registerUser = async (name, email, password, roleId) => {
    const transaction = await sequelize.transaction();

    try {
        // Check for existing user within a transaction
        const existingUser = await User.findOne({
            where: {email},
            transaction
        });

        if (existingUser) {
            await transaction.rollback();
            throw new Error('User already exists');
        }

        const hashedPassword = await bcrypt.hash(password, 12);
        const newUser = await User.create(
            {name, email, password: hashedPassword, roleId},
            {transaction}
        );

        await transaction.commit();
        return generateToken(newUser);
    } catch (error) {
        await transaction.rollback();
        throw error;
    }
};

const loginUser = async (email, password) => {
    const user = await User.findOne({where: {email}});
    if (!user || !(await bcrypt.compare(password, user.password))) throw new Error('Invalid credentials');
    return generateToken(user);
};

module.exports = {registerUser, loginUser};
