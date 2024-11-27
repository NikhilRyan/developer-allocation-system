.PHONY: build run test docker-build docker-run migrate

build:
    go build -o bin/main ./cmd/app/main.go

run:
    go run ./cmd/app/main.go

test:
    go test ./...

docker-build:
    docker build -t developer-allocation-system .

docker-run:
    docker-compose up

migrate:
    ./scripts/migrate.sh
