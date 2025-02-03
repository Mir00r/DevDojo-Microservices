const bcrypt = require('bcryptjs');
const jwt = require('jsonwebtoken');
const User = require('../models/User');

const generateToken = (user) => {
    return jwt.sign({id: user.id, role: user.role}, process.env.JWT_SECRET, {expiresIn: '1h'});
};

const registerUser = async (name, email, password) => {
    const existingUser = await User.findOne({where: {email}});
    if (existingUser) throw new Error('User already exists');

    const hashedPassword = await bcrypt.hash(password, 10);
    const newUser = await User.create({name, email, password: hashedPassword});
    return generateToken(newUser);
};

const loginUser = async (email, password) => {
    const user = await User.findOne({where: {email}});
    if (!user || !(await bcrypt.compare(password, user.password))) throw new Error('Invalid credentials');
    return generateToken(user);
};

module.exports = {registerUser, loginUser};
