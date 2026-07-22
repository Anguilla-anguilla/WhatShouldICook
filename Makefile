GOCMD = go
GOBUILD = $(GOCMD) build
GOTEST = $(GOCMD) test
BIN_DIR = bin 
APP_NAME = whatshouldicook

build:
	$(GOBUILD) ./cmd/api -o $(BIN_DIR)/$(APP_NAME)

run: build
	./$(BIN_DIR)/$(APP_NAME)


# test:

deps:
	$(GOCMD) mod download
	$(GOCMD) mod tidy

fmt:
	$(GOCMD) fmt ./...

clean:
	rm -rf $(BIN_DIR)/

.PHONY: run build test clean