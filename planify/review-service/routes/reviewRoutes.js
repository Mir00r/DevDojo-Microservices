const express = require('express');
const {createReview, getReviews} = require('../domains/reviews/controllers/reviewController');
const authenticate = require('../../authentication-service/src/middleware/authMiddleware');
const router = express.Router();

router.post('/', authenticate, createReview);
router.get('/', getReviews);

module.exports = router;
