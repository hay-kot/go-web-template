interface Config {
  PORT: string;
  HOST: string;
  BASE_URL: string;
}

class Configuration implements Config {
  PORT: string;
  HOST: string;
  BASE_URL: string;
  constructor() {
    (this.PORT = process.env.API_WEB_PORT || "3000"),
      (this.HOST = "http" + process.env.API_WEB_HOST || "http://127.0.0.1"),
      (this.BASE_URL = this.HOST + ":" + this.PORT);
  }
}

export const config = new Configuration();
