require('dotenv').config();
const app = require('./app');
const {connectDB} = require('./configs/database');

const PORT = process.env.PORT || 5000;

// Start the server after connecting to the database
connectDB().then(() => {
    app.listen(PORT, () => {
        console.log(`🚀 Auth Service running on port ${PORT}`);
    });
});
