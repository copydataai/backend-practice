const { Pool } = require('pg')
const express = require('express');
const ServerlessHttp = require('serverless-http');

const pool = new Pool();
async function login(user) {
  const { rowCount, rows } = await pool.query(`SELECT password FROM users WHERE email = $1;`, [user.email])

  if (!rowCount) {
    return "first Signup"
  }
  if (user.password == rows[0].password) {
    return jwt
  } else {
    return "Enter your password correct"
  }
}

async function signUp(user) {
  const { rowCount } = await pool.query('SELECT email FROM public.users WHERE email = $1', [user.email])
  if (rowCount) {
    return "this account exist"
  }
  const newUser = await pool.query(`INSERT INTO public.users(name, country, speciality, role, created_at, email, password) VALUES($1,$2,$3, $4, $5, $6,$7);`, [user.name, user.country, user.speciality, user.role, JSON.stringify(new Date), user.email, user.password]);
  if (newUser.rowCount) {
    return "User create"
  }
}


// Routes
const app = express();

app.use(express.json())

app.use((req, res, next) => {
  console.log(req.url, req.body, req.statusCode)
  next()
})

app.post('/signup', async (req, res, next) => {
  try {
    const user = req.body
    const create = await signUp(user)
    res.status(201).json({
      create
    })
  } catch (err) {
    next(err)
  }
})

app.post('/login', async (req, res, next) => {
  try {
    const user = req.body
    const userLogin = await login(user)
    res.status(200).json({
      userLogin
    })
  } catch (err) {
    next(err);
  }
})

app.get('/hello', (req, res, next) => {
  res.json({ hello: "hello World" })
})

app.use((req, res, next) => {
  return res.status(404).json({
    error: "Not Found",
  });
});

// app.listen(3000, () => {
//   console.log('listening');
// })
//
module.exports.handler = ServerlessHttp(app)
