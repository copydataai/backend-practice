const { app } = require('../../src/app.js');
const request = require('supertest');

describe("GET /reviews simple", () => {
  const endpoint = "/reviews"
  test("should return a 200 status code", async () => {
    const res = await request(app).get(endpoint).send();
    expect(res.statusCode).toBe(200);
  })

  test("should return an object as a respond", async () => {
    const res = await request(app).get(endpoint).send();
    expect(res.body).toBeInstanceOf(Object);

  })
  test("should return a header content-type with value application/json", async () => {
    const res = await request(app).get(endpoint).send();
    expect(res.type).toEqual(expect.stringContaining('json'));
  })
})

describe("GET /reviews:id simple", () => {
  const endpoint = "/reviews/25"
  test("should return a 200 status code", async () => {
    const res = await request(app).get(endpoint).send();
    expect(res.statusCode).toBe(200);
  })

  test("should return an object as a respond", async () => {
    const res = await request(app).get(endpoint).send();
    expect(res.body).toBeInstanceOf(Object);

  })
  test("should return a header content-type with value application/json", async () => {
    const res = await request(app).get(endpoint).send();
    expect(res.type).toEqual(expect.stringContaining('json'));
  })
})
