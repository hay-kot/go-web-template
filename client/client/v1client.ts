import { fetch } from "cross-fetch";
import { IApiClient } from "./client";

export class v1ApiClient implements IApiClient {
  version: string;
  baseUrl: string;

  constructor(baseUrl: string, version = "v1") {
    this.version = version;
    this.baseUrl = baseUrl;
  }

  async status() {
    return fetch(`${this.baseUrl}/api/status`);
  }
}
