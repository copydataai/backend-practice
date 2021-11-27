const { Router } = require('express');


const router = Router();

router.post('/', (req, res) {
  res.json({cout: "Hello World"})
})


module.exports = router
