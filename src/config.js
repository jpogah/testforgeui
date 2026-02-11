const DEFAULT_HOST = '127.0.0.1';
const DEFAULT_PORT = 3000;

function parsePort(rawPort) {
  const parsed = Number.parseInt(rawPort, 10);

  if (!Number.isInteger(parsed) || parsed <= 0 || parsed > 65535) {
    return DEFAULT_PORT;
  }

  return parsed;
}

function getConfig(env = process.env) {
  const host = env.HOST || DEFAULT_HOST;
  const port = parsePort(env.PORT);

  return { host, port };
}

module.exports = {
  DEFAULT_HOST,
  DEFAULT_PORT,
  getConfig,
};
