const Review = require('../models/Review');

const createReview = async (userId, content, rating) => {
    return await Review.create({userId, content, rating});
};

const getReviews = async () => {
    return await Review.findAll({include: 'User'});
};

module.exports = {createReview, getReviews};
