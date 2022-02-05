import { getClientV1 } from "../../client";
import { describe, it, expect } from "vitest";
import * as config from "../config";

const client = getClientV1(config.BASE_URL);

describe("GET /api", function () {
  it("server provides schemas", async function (done) {
    try {
      const res = await client.summary();
      expect(res.status).toBe(200);
      expect(res.statusText).toBe("OK");

      expect(res.data).toEqual({
        health: true,
        versions: ["v1"],
        title: "Go API Template",
        message: "Welcome to the Go API Template Application!",
      });

      done();
    } catch (err) {
      done(err);
    }
  });
});

describe("GET /api/status", function () {
  it("server is available", async function (done) {
    try {
      const res = await client.status();
      expect(res.status).toBe(200);
      expect(res.statusText).toBe("OK");

      expect(res.data).toEqual({
        status: "available",
        message: "The service is ready to use",
      });

      done();
    } catch (err) {
      done(err);
    }
  });
});
