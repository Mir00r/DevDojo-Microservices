const jwt = require('jsonwebtoken');
require('dotenv').config();

/**
 * Middleware to authenticate users via JWT.
 */
const authenticate = (req, res, next) => {
    const token = req.header('Authorization');

    if (!token) {
        return res.status(401).json({error: 'Access Denied. No token provided.'});
    }

    try {
        const decoded = jwt.verify(token.replace("Bearer ", ""), process.env.JWT_SECRET);
        req.user = decoded; // Attach user info to request object
        next();
    } catch (err) {
        return res.status(403).json({error: 'Invalid token.'});
    }
};

module.exports = authenticate;
