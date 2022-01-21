const db = require('../db/knex')
const knex = db.connection();

module.exports = {
  findAll: async (pag, limit) => {
    let reviews = await knex.select('*').from('reviews').limit(parseInt(limit)).offset(parseInt(pag));
    if (!reviews) { throw new Error("Not found reviews") }
    return { reviews };
  },
  findOne: async (id) => {
    let review = await knex('reviews').column('title', 'content', { author: 'users.name' }, { product: 'product_type.name' }).select().join('users', 'users.id', '=', 'reviews.user').join('products', 'products.id', '=', 'reviews.product').join('product_type', 'products.detail', '=', 'product_type.id');
    if (!review) { throw new Error("Not found review") }
    return { review };
  },
}
