const { app } = require('../src/app')

describe("App", () => {
  test("should be thruty", () => {
    expect(app).toBeTruthy();
  })

  test("should be defined", () => {
    expect(app).toBeDefined();
  })

  test("should be a function", () => {
    expect(app).toBeInstanceOf(Function);
  })

  test("should have method a method listen", () => {
    expect(app.listen).toBeInstanceOf(Function);
  })
})
