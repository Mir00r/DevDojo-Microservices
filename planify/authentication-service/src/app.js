const express = require('express');
const cors = require('cors');
const helmet = require('helmet');
const morgan = require('morgan');
const authRoutes = require('./routes/authRoutes');
const swagger = require('../config/swagger');

const app = express();

// Middleware
app.use(express.json());  // Parse JSON requests
app.use(cors());          // Enable CORS
app.use(helmet());        // Security headers
app.use(morgan('dev'));   // Logging

// Swagger documentation route
app.use('/api-docs', swagger.serve, swagger.setup);


// Routes
app.use('/api/auth', authRoutes);

// Error handling middleware
app.use((err, req, res, next) => {
    console.error(err.stack);

    // Send more specific error responses
    const statusCode = err.statusCode || 500;
    const message = err.message || 'Internal Server Error';

    res.status(statusCode).json({
        status: 'error',
        message,
        ...(process.env.NODE_ENV === 'development' && {stack: err.stack})
    });
});

// 404 handler
app.use((req, res) => {
    res.status(404).json({
        status: 'error',
        message: 'Route not found'
    });
});

module.exports = app;
