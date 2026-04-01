.PHONY: run build test clean docker-up docker-down

APP_NAME=myapp-test

run:
	go run cmd/server/main.go

build:
	go build -o bin/$(APP_NAME) cmd/server/main.go

test:
	go test -v ./...

clean:
	rm -rf bin/

docker-up:
	docker-compose up -d --build

docker-down:
	docker-compose down
