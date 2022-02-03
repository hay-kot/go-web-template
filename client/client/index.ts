import { v1ApiClient } from "./v1client";
import { IApiClient } from "./client";

export function getClient(baseUrl: string, version = "v1"): IApiClient {
  switch (version) {
    case "v1":
      return new v1ApiClient(baseUrl, version);
    default:
      throw new Error(`Unsupported version: ${version}`);
  }
}
