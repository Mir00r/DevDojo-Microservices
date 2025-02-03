const reviewService = require('../services/reviewService');

exports.createReview = async (req, res) => {
    try {
        const {content, rating} = req.body;
        const review = await reviewService.createReview(req.user.id, content, rating);
        res.status(201).json(review);
    } catch (error) {
        res.status(400).json({error: error.message});
    }
};

exports.getReviews = async (req, res) => {
    try {
        const reviews = await reviewService.getReviews();
        res.json(reviews);
    } catch (error) {
        res.status(500).json({error: error.message});
    }
};
