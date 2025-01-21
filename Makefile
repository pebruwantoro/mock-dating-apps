.PHONY: 

all: build/main

build/main: cmd/main.go generated
	@echo "Building..."
	go build -o $@ $<

init:
	go mod tidy
	go mod vendor

test:
	go clean -testcache
	go test -short -coverprofile coverage.out -short -v ./...

docker-up:
	docker compose up --build -d

docker-down:
	docker compose down --volumes