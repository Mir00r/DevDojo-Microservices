require('dotenv').config();
const app = require('./app');
const {connectDB} = require('./configs/database');

const PORT = process.env.PORT || 5000;

const startServer = async () => {
    try {
        // Connect to database and initialize models
        await connectDB();

        // Start the server
        app.listen(PORT, () => {
            console.log(`🚀 Auth Service running on port ${PORT}`);
        });
    } catch (error) {
        console.error('Failed to start server:', error);
        process.exit(1);
    }
};

startServer();
