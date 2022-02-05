import { AxiosResponse } from "axios";

export interface IApiClient {
  status: () => Promise<AxiosResponse<unknown>>;
}
