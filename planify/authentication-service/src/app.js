const express = require('express');
const cors = require('cors');
const helmet = require('helmet');
const morgan = require('morgan');
const authRoutes = require('./routes/authRoutes');
// const userRoutes = require('./routes/user.routes');

const app = express();

// Middleware
app.use(express.json());  // Parse JSON requests
app.use(cors());          // Enable CORS
app.use(helmet());        // Security headers
app.use(morgan('dev'));   // Logging

// Routes
app.use('/api/auth', authRoutes);
// app.use('/api/users', userRoutes);

// Error handling middleware
app.use((err, req, res, next) => {
    console.error(err.stack);
    res.status(500).json({error: 'Internal Server Error'});
});

module.exports = app;
