const http = require('node:http');
const { route } = require('./router');

function requestHandler(req, res) {
  route(req, res);
}

function createServer() {
  return http.createServer(requestHandler);
}

function startServer(config) {
  const server = createServer();

  server.listen(config.port, config.host, () => {
    console.log(`Server listening on http://${config.host}:${config.port}`);
  });

  return server;
}

if (require.main === module) {
  const port = Number(process.env.PORT || 3000);
  const host = process.env.HOST || 'localhost';
  startServer({ port, host });
}

module.exports = {
  createServer,
  requestHandler,
  startServer,
};
