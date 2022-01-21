const { Router } = require('express');
const { findAll, findOne } = require('../services/reviews')

const router = Router();

router.get("/", getReviews)
router.get("/:id", getReviewById)
async function getReviews(req, res, next) {
  try {
    const { pag, index } = req.query;
    const reviews = await findAll(pag, index);
    res.json({
      reviews
    })
  } catch (err) {
    next(err)
  }
}
async function getReviewById(req, res, next) {
  try {
    const { id } = req.params;
    const review = await findOne(id);
    res.json({
      review
    });
  } catch (err) {
    next(err);
  }
}

module.exports = router;
