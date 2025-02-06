// const {Sequelize} = require('sequelize');
// require('dotenv').config();
//
// const sequelize = new Sequelize(process.env.DB_NAME, process.env.DB_USER, process.env.DB_PASS, {
//     host: process.env.DB_HOST,
//     dialect: 'postgres',
//     logging: false,  // Disable SQL logs in production
// });
//
// module.exports = sequelize;


const {Sequelize} = require('sequelize');

// Define Sequelize connection
const sequelize = new Sequelize(process.env.DATABASE_URL, {
    dialect: 'postgres',
    logging: false, // Disable logging for cleaner output
});

// Function to connect to the database
const connectDB = async () => {
    try {
        await sequelize.authenticate();
        console.log('✅ Database connected successfully');
    } catch (error) {
        console.error('❌ Database connection failed:', error);
        process.exit(1);
    }
};

module.exports = {connectDB, sequelize}; // Ensure connectDB is exported
