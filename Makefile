all: build

APP_NAME ?= ./temp

# Default target: build the binary
all: build

build:
	go build -o ${APP_NAME} ./cmd/main.go

run:
	air

dev: run

lint:
	staticcheck ./...

vet:
	go vet ./...

check: vet lint test

fmt:
	go fmt ./...

test:
	go test ./...

clean:
	rm ${APP_NAME}
