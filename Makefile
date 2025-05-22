all: build

APP_NAME ?= ./temp

# Default target: build the binary
all: build

prod:
	docker compose -f prod.yml up

build:
	go build -o ${APP_NAME} ./cmd/main.go

run:
	docker compose up

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
