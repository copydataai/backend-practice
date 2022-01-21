const express = require('express');
const cors = require('cors')
const { badImplementation, notFound, badRequest } = require('@hapi/boom')

const { ManageRoutes } = require('./controllers');


const app = express();

app.use(express.json());
app.use(express.urlencoded({ extended: false }));
app.use(cors({}))

app.use((req, res, next) => {
  console.log(req.method, req.url, req.statusCode)
  next()
})

app.get('/hello', (req, res, next) => {
  res.json({ hello: "hello World EC2" })
})
ManageRoutes(app);

// middleware to deal with 404 error
app.use((req, res, next) => {
  //let err = new Error('route does not exist')
  //err.status(404)
  next(badRequest())  // send error to next middleware
})

// receives error from last middleware
app.use((err, req, res, next) => {
  // if error 404, sends back message 'route does not exist'
  // otherwise it sends Murphy's message
  res.status(err.status || 500).send(err.message || `Don't force it get a larger hammer.`)
})

module.exports = { app };
