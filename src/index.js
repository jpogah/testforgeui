const { getConfig } = require('./config');
const { startServer } = require('./server');

const config = getConfig();
const server = startServer(config);

function shutdown(signal) {
  console.log(`Received ${signal}. Shutting down...`);
  server.close((error) => {
    if (error) {
      console.error('Failed to shut down cleanly:', error);
      process.exit(1);
    }

    process.exit(0);
  });
}

process.on('SIGINT', () => shutdown('SIGINT'));
process.on('SIGTERM', () => shutdown('SIGTERM'));

server.on('error', (error) => {
  console.error('Server failed to start:', error);
  process.exit(1);
});
