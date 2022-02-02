import { exec } from "child_process";
import * as config from "./config";

before(() => {
  console.log("Starting Client Tests");
  console.log({
    PORT: config.PORT,
    HOST: config.HOST,
    BASE_URL: config.BASE_URL,
  });
});

after(() => {
  const pc = exec("pkill -SIGTERM api"); // Kill background API process

  pc.stdout.on("data", (data) => {
    console.log(`stdout: ${data}`);
  });
});
