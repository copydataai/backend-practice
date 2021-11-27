const config = {
  node: {
    port: process.env.PORT,
    host: process.env.HOST,
    secret: process.env.SECRET
  },
  postgres: {
    host: process.env.POST_HOST,
    port: process.env.POST_PORT,
    user: process.env.POST_USER,
    password: process.env.POST_PASS,
    database: process.env.POST_DB
  }
}
