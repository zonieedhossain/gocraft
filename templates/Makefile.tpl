APP_NAME := {{ .AppName | default "gocraft-app" }}

run:
	go run ./cmd/main.go

build:
	go build -o bin/$(APP_NAME) ./cmd/main.go

migrate-up:
	migrate -path ./migrations -database $$DB_URL up

migrate-down:
	migrate -path ./migrations -database $$DB_URL down 1

