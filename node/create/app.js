const { Pool } = require('pg');
const express = require('express');
const ServerlessHttp = require('serverless-http');


const pool = new Pool();

async function reviews(review, user) {
  await pool.query('INSERT INTO reviews(title, content, "user", product, created_at) VALUES($1, $2, $3, $4, $5);', [review.title, review.content, user, review.product, JSON.stringify(new Date)]);
  return "review created"
}

async function productType(name, detail) {
  await pool.query('INSERT INTO product_type(name, detail) VALUES($1, $2);', [name, detail])
  const { rows } = await pool.query('SELECT id from product_type where name = $1', [name])
  return rows[0].id
}

async function products(product) {
  const id = await productType(product.name, product.detail)
  await pool.query('INSERT INTO products(detail, trademark, manufacturing, "SKU", created_at) VALUES($1, $2, $3, $4, $5);', [id, product.trademark, product.manufacturing, product.sku, JSON.stringify(new Date)]);
  return "product created"
}

const app = express();
app.use(express.json());

app.use((req, res, next) => {
  console.log(req.url, req.body, req.headers);
  next()
})

app.post('/reviews', async (req, res, next) => {
  try {
    const review = req.body
    const newReview = await reviews(review);
    res.status(201).json({
      output: newReview
    })
  } catch (err) {
    next(err)
  }

})
//change routes for declare to nginx
app.post('/products', async (req, res, next) => {
  try {
    const product = req.body;
    const newProduct = await products(product);
    res.status(201).json({
      output: newProduct,
    })

  } catch (err) {
    next(err)
  }
})
app.get('/hello', (req, res, next) => {
  res.json({ hello: "hello World" })
})

app.use((req, res, next) => res.status(404).json({
  error: "Not Found"
}));


module.exports.handler = ServerlessHttp(app)
