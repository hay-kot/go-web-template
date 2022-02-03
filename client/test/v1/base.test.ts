import { getClient } from "../../client";
import { describe, it, expect } from "vitest";
import * as config from "../config";

const client = getClient(config.BASE_URL, "v1");

describe("GET /api/status", function () {
  it("server is available", async function (done) {
    try {
      const res = await client.status();

      expect(res.status).toBe(200);
      expect(res.statusText).toBe("OK");
      expect(res.headers.get("content-type")).toBe("application/json");

      const body = await res.json();
      expect(body).toEqual({
        status: "available",
        message: "The service is ready to use",
      });

      done();
    } catch (err) {
      done(err);
    }
  });
});
