const knex = require('knex')

const { host, port, database, user, password, client } = require('../config/database')

const connection = () => {
  try {
    const connection = knex({
      client,
      connection: `${client}://${user}:${password}@${host}:${port}/${database}`,
      // migrations: {
      //   tableName
      // }
    });
    return connection;
  } catch (err) {
    console.log(err);
  }
}

module.exports = { connection };
