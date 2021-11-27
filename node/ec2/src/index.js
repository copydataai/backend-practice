const express = require('express');

const { config } = require('./config')

const {host, port} = config.node

const app = express();



app.listen(port, host, () => {
  console.log("listening on ${host}:${port}");
});
