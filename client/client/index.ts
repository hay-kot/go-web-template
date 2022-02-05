import { v1ApiClient } from "./v1client";
import { IApiClient } from "./client";

export function getClientV1(baseUrl: string): v1ApiClient {
  return new v1ApiClient(baseUrl, "v1");
}
