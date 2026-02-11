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

## Run

```bash
npm start
```

## Test

```bash
npm test
```

## Verify (local/CI)

```bash
npm run verify
```
