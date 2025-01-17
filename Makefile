GIT_VERSION := $(shell git describe --abbrev=8 --dirty --always --tags)
GIT_COMMIT := $(shell git rev-parse --short HEAD)
BUILD_TIME := $(shell date +%s)

.PHONY: all build test lint clean

all: lint test build

protobuf:
	protoc -I schema/protobuf/sync_pb/ schema/protobuf/sync_pb/*.proto --go_out=schema/protobuf/sync_pb/

build:
	go run main.go

test:
	go test -v ./...

lint:
	docker run -t --rm -v "$$(pwd):/app" -w /app golangci/golangci-lint golangci-lint run -v

clean:
	rm -f sync-server

docker:
	COMMIT=$(GIT_COMMIT) VERSION=$(GIT_VERSION) BUILD_TIME=$(BUILD_TIME) docker-compose build

docker-up:
	COMMIT=$(GIT_COMMIT) VERSION=$(GIT_VERSION) BUILD_TIME=$(BUILD_TIME) docker-compose up

docker-test:
	COMMIT=$(GIT_COMMIT) VERSION=$(GIT_VERSION) BUILD_TIME=$(BUILD_TIME) docker-compose -f docker-compose.yml run --rm dev make test

instrumented:
	gowrap gen -p github.com/brave/go-sync/datastore -i Datastore -t ./.prom-gowrap.tmpl -o ./datastore/instrumented_datastore.go
	gowrap gen -p github.com/brave/go-sync/cache -i RedisClient -t ./.prom-gowrap.tmpl -o ./cache/instrumented_redis.go
