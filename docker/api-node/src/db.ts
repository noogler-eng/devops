import pg from 'pg'
const { Pool } = pg
import dotenv from "dotenv";
dotenv.config();

const pool = new Pool({
  connectionString: process.env.DATABASE_URL || "",
});

// the pool will emit an error on behalf of any idle clients
// it contains if a backend error or network partition happens
pool.on('error', (err, client) => {
  console.error('Unexpected error on idle client', err);
  process.exit(-1);
});

// async/await - check out a client
// another way to connect with the postgres database
const getDateTime = async () => {
  const client = await pool.connect();
  try {
    const res = await client.query('SELECT NOW() as now;');
    console.log(res);
    return res.rows[0];
  } catch (err) {
    console.log(err);
  } finally {
    client.release();
  }
};

export default getDateTime;