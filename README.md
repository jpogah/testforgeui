# testforgeui

## Web Service

A Go web service scaffold is available at `cmd/webservice` with core service code in `internal/webservice` and configuration loading in `internal/config`.

### Targets

- `make run` - run the web service locally
- `make build` - compile binary to `bin/webservice`
- `make test` - run all Go tests
- `make tidy` - sync module dependencies

### Configuration

Use environment variables directly or a local `.env` file:

- `APP_ENV` (default: `development`)
- `PORT` (default: `8080`)

Sample values are in `.env.example`.

## Hello World API

Minimal HTTP service that exposes a hello-world endpoint for local development and testing.

## Prerequisites

- Node.js 18+ (LTS recommended)
- npm 9+

## Setup

1. Clone the repository and enter it:
```bash
git clone <repo-url>
cd <repo-folder>
```
2. Install dependencies:
```bash
npm install
```

## Run

- Development mode:
```bash
npm run dev
```
- Production mode:
```bash
npm run build
npm start
```

Default local URL: `http://localhost:3000`

## Hello-World Endpoint Example

Request:
```bash
curl -i http://localhost:3000/hello-world
```

Example response:
```http
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8

{"message":"Hello, world!"}
```

## Test

```bash
npm test
```

## Verify (local/CI)

```bash
npm run verify
```
