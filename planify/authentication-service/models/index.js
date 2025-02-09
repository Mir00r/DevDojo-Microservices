const fs = require('fs');
const path = require('path');
const {Sequelize} = require('sequelize');
const config = require('../config/config.json')['development'];

const sequelize = new Sequelize(config.database, config.username, config.password, {
    host: config.host,
    dialect: config.dialect,
    dialectOptions: config.dialectOptions,
    logging: false,
    define: {
        schema: 'auth'
    }
});

const db = {};

// Initialize models
fs.readdirSync(__dirname)
    .filter(file => file.endsWith('.js') && file !== 'index.js')
    .forEach(file => {
        const model = require(path.join(__dirname, file))(sequelize, Sequelize.DataTypes);
        db[model.name] = model;
    });

// Initialize associations
Object.keys(db).forEach(modelName => {
    if (db[modelName].associate) {
        db[modelName].associate(db);
    }
});

const initialize = async () => {
    try {
        // Create schema
        await sequelize.createSchema('auth').catch(() => console.log('Schema already exists'));

        // Sync models
        await sequelize.sync();

        console.log('✅ Database initialized successfully');
    } catch (error) {
        console.error('❌ Database initialization failed:', error);
        throw error;
    }
};

db.sequelize = sequelize;
db.Sequelize = Sequelize;
db.initialize = initialize;

module.exports = db;
