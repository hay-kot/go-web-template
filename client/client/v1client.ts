import { IApiClient } from "./client";
import axios, { Axios } from "axios";

interface Status {
  status: string;
  message: string;
}

export interface ApiSummary {
  health: boolean;
  versions: string[];
  title: string;
  message: string;
}

export class v1ApiClient implements IApiClient {
  version: string;
  baseUrl: string;
  requests: Axios;

  constructor(baseUrl: string, version = "v1") {
    this.version = version;
    this.baseUrl = baseUrl;
    this.requests = axios.create({
      baseURL: `${this.baseUrl}/${this.version}`,
    });
  }

  async status() {
    return this.requests.get<Status>(`${this.baseUrl}/api/status`);
  }

  async summary() {
    return this.requests.get<ApiSummary>(`${this.baseUrl}/api`);
  }
}
