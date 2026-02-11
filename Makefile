GO ?= go
BINARY ?= webservice
CMD_PATH ?= ./cmd/webservice
BIN_DIR ?= ./bin

.PHONY: run build test tidy clean

run:
	$(GO) run $(CMD_PATH)

build:
	mkdir -p $(BIN_DIR)
	$(GO) build -o $(BIN_DIR)/$(BINARY) $(CMD_PATH)

test:
	$(GO) test ./...

tidy:
	$(GO) mod tidy

clean:
	rm -rf $(BIN_DIR)
