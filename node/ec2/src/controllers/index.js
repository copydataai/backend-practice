const reviews = require('./reviews.controller.js');

function ManageRoutes(app) {
  app.use("/posts", reviews);
}

module.exports = { ManageRoutes };
