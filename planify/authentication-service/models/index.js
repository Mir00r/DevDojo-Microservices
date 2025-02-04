const fs = require('fs');
const path = require('path');
const {Sequelize} = require('sequelize');
const config = require('../config/config.json')['development'];

const sequelize = new Sequelize(config.database, config.username, config.password, {
    host: config.host, dialect: config.dialect, dialectOptions: config.dialectOptions, logging: console.log
});

const db = {};

// 🚀 Ensure schemas exist before model initialization
const schemas = ['auth'];

async function initializeSchemas() {
    await Promise.all(schemas.map(schema => sequelize.createSchema(schema, {logging: false}).catch(() => {
    })));
}

initializeSchemas().then(() => {
    // 📌 Read model files dynamically
    fs.readdirSync(__dirname)
        .filter(file => file.endsWith('.js') && file !== 'index.js')
        .forEach(file => {
            const model = require(path.join(__dirname, file))(sequelize, Sequelize.DataTypes);
            db[model.name] = model;
        });

    // 📌 Define associations if any
    Object.keys(db).forEach(modelName => {
        if (db[modelName].associate) {
            db[modelName].associate(db);
        }
    });

    db.sequelize = sequelize;
    db.Sequelize = Sequelize;
});

module.exports = db;
