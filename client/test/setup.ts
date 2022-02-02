import { exec } from "child_process";

before(() => {
  console.log("Hello World");
  const cp = exec("cd ../../ && make api");

  cp.stdout.on("data", (data) => {
    console.log(data.toString());
  });
});

after(() => {
  console.log("Goodbye World");
});
