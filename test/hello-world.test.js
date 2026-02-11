const test = require("node:test");
const assert = require("node:assert/strict");
const { requestHandler } = require("../src/server");

test("GET /hello-world returns 200 and hello world response", async () => {
  const req = {
    method: "GET",
    url: "/hello-world"
  };

  const response = await new Promise((resolve) => {
    const res = {
      statusCode: 0,
      headers: {},
      body: "",
      writeHead(statusCode, headers) {
        this.statusCode = statusCode;
        this.headers = headers;
      },
      end(body = "") {
        this.body = body;
        resolve(this);
      }
    };

    requestHandler(req, res);
  });

  assert.equal(response.statusCode, 200);
  assert.equal(response.body, "hello world");
});
