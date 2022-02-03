export interface IApiClient {
  status: () => Promise<Response>;
}
