import request from "supertest";
import assert from "assert";
import { config } from "../config";

const api = request(config.BASE_URL);

describe("GET /api", function () {
  it("responds with json", function (done) {
    api
      .get("/api")
      .set("Accept", "application/json")
      .expect("Content-Type", /json/)
      .expect(200, done);
  });
});
