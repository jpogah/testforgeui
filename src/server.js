const http = require("node:http");

function requestHandler(req, res) {
  if (req.method === "GET" && req.url === "/hello-world") {
    res.writeHead(200, { "content-type": "text/plain; charset=utf-8" });
    res.end("hello world");
    return;
  }

  res.writeHead(404, { "content-type": "text/plain; charset=utf-8" });
  res.end("not found");
}

function createServer() {
  return http.createServer(requestHandler);
}

if (require.main === module) {
  const port = Number(process.env.PORT || 3000);
  createServer().listen(port, () => {
    console.log(`Server listening on http://localhost:${port}`);
  });
}

module.exports = {
  createServer,
  requestHandler
};
