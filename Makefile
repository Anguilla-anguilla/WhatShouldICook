GOCMD = go
GOBUILD = $(GOCMD) build
GOTEST = $(GOCMD) test
BIN_DIR = bin 
APP_NAME = whatshouldicook

DB_URL = $(shell grep DB_URL .env | cut -d '=' -f2)
MIGRATE_PATH = internal/migrations


build:
	$(GOBUILD) ./cmd/api -o $(BIN_DIR)/$(APP_NAME)

run: build
	./$(BIN_DIR)/$(APP_NAME)

migrate-up:
	migrate -path $(MIGRATE_PATH) -database "$(DB_URL)" up

migrate-down:
	migrate -path $(MIGRATE_PATH) -database "$(DB_URL)" down

migrate-new:
	migrate create -ext sql -dir $(MIGRATE_PATH) -seq(name)

test:
	$(GOTEST) ./...

deps:
	$(GOCMD) mod download
	$(GOCMD) mod tidy

fmt:
	$(GOCMD) fmt ./...

clean:
	rm -rf $(BIN_DIR)/

.PHONY: run build test clean migrate-up migrate-down migrate-new deps fmt