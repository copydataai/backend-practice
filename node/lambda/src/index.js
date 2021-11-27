const express = require('express');

const {config} = require('./config')
const reviews = require('./controllers')

const {port, host} = config.node

const app = express();
app.use('/reviews', reviews);

app.listen(port, host, () => {
  console.log("listening on ${host}:{port}");
});
