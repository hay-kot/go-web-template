interface Config {
  PORT: number;
  HOST: string;
  BASE_URL: string;
}

export const config: Config = Object.freeze({
  PORT: 3915,
  HOST: "http://localhost",
  BASE_URL: "http://localhost:" + 3915,
});
