const db = require('../../models');

const connectDB = async () => {
    try {
        await db.sequelize.authenticate();
        console.log('✅ Database connected successfully');

        // Initialize models and schema
        await db.initialize();
        console.log('✅ Models initialized successfully');

        return db;
    } catch (error) {
        console.error('❌ Database connection failed:', error);
        process.exit(1);
    }
};

module.exports = {connectDB, sequelize: db.sequelize};
