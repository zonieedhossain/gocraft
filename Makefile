PROJECT_NAME ?= clean-arch-demo

.PHONY: init
init:
	@echo "Creating full project structure: $(PROJECT_NAME)..."

	# Root files
	@mkdir -p $(PROJECT_NAME)
	@touch $(PROJECT_NAME)/README.md
	@touch $(PROJECT_NAME)/go.mod
	@touch $(PROJECT_NAME)/go.sum

	# .gitignore
	@echo "# Go Build" > $(PROJECT_NAME)/.gitignore
	@echo "bin/" >> $(PROJECT_NAME)/.gitignore
	@echo "build/" >> $(PROJECT_NAME)/.gitignore
	@echo "*.out" >> $(PROJECT_NAME)/.gitignore
	@echo ".vscode/" >> $(PROJECT_NAME)/.gitignore
	@echo ".idea/" >> $(PROJECT_NAME)/.gitignore
	@echo ".env" >> $(PROJECT_NAME)/.gitignore
	@echo ".DS_Store" >> $(PROJECT_NAME)/.gitignore

	# CMD
	@mkdir -p $(PROJECT_NAME)/cmd
	@echo "package main\n\n// main.go is the application entrypoint" > $(PROJECT_NAME)/cmd/main.go

	# Deployments
	@mkdir -p $(PROJECT_NAME)/deployments/docker
	@mkdir -p $(PROJECT_NAME)/deployments/k8s
	@mkdir -p $(PROJECT_NAME)/deployments/terraform
	@touch $(PROJECT_NAME)/deployments/docker/.gitkeep
	@touch $(PROJECT_NAME)/deployments/k8s/.gitkeep
	@touch $(PROJECT_NAME)/deployments/terraform/.gitkeep

	# Docs
	@mkdir -p $(PROJECT_NAME)/docs
	@echo "# docs: contains design, architecture, and documentation files" > $(PROJECT_NAME)/docs/README.md

	# Internal
	@mkdir -p $(PROJECT_NAME)/internal/config
	@echo "package config\n\n// config: loads and manages application configuration" > $(PROJECT_NAME)/internal/config/config.go

	@mkdir -p $(PROJECT_NAME)/internal/domain/entity
	@echo "package entity\n\n// entity: contains core business entities (models)" > $(PROJECT_NAME)/internal/domain/entity/entity.go

	@mkdir -p $(PROJECT_NAME)/internal/domain/service
	@echo "package service\n\n// service: defines domain service interfaces" > $(PROJECT_NAME)/internal/domain/service/service.go

	@mkdir -p $(PROJECT_NAME)/internal/usecase
	@echo "package usecase\n\n// usecase: application-specific business logic" > $(PROJECT_NAME)/internal/usecase/usecase.go

	@mkdir -p $(PROJECT_NAME)/internal/handler/http
	@echo "package http\n\n// http: handles HTTP routes and controllers" > $(PROJECT_NAME)/internal/handler/http/http_handler.go

	@mkdir -p $(PROJECT_NAME)/internal/handler/grpc
	@echo "package grpc\n\n// grpc: gRPC handlers and server implementation" > $(PROJECT_NAME)/internal/handler/grpc/grpc_handler.go

	@mkdir -p $(PROJECT_NAME)/internal/repository/postgres
	@echo "package postgres\n\n// postgres: database operations using PostgreSQL" > $(PROJECT_NAME)/internal/repository/postgres/repo.go

	@mkdir -p $(PROJECT_NAME)/internal/repository/redis
	@echo "package redis\n\n// redis: caching or session storage using Redis" > $(PROJECT_NAME)/internal/repository/redis/cache.go

	@mkdir -p $(PROJECT_NAME)/internal/middleware
	@echo "package middleware\n\n// middleware: handles auth, logging, rate-limiting, etc." > $(PROJECT_NAME)/internal/middleware/middleware.go

	@mkdir -p $(PROJECT_NAME)/internal/service
	@echo "package service\n\n// service: integrates with 3rd-party APIs (email, payment, etc.)" > $(PROJECT_NAME)/internal/service/service.go

	@mkdir -p $(PROJECT_NAME)/internal/job
	@echo "package job\n\n// job: background jobs and scheduled workers" > $(PROJECT_NAME)/internal/job/job.go

	@mkdir -p $(PROJECT_NAME)/internal/logger
	@echo "package logger\n\n// logger: initializes structured logging (Zap, Logrus)" > $(PROJECT_NAME)/internal/logger/logger.go

	@mkdir -p $(PROJECT_NAME)/internal/errors
	@echo "package errors\n\n// errors: app-wide error definitions and wrapping" > $(PROJECT_NAME)/internal/errors/errors.go

	@mkdir -p $(PROJECT_NAME)/internal/util
	@echo "package util\n\n// util: helper functions (validators, string utils, etc.)" > $(PROJECT_NAME)/internal/util/util.go

	# Migrations
	@mkdir -p $(PROJECT_NAME)/migrations
	@echo "-- migrations: contains SQL schema changes" > $(PROJECT_NAME)/migrations/README.md

	# Tests
	@mkdir -p $(PROJECT_NAME)/test/fixtures
	@mkdir -p $(PROJECT_NAME)/test/mocks
	@echo "package test\n\n// test_sample: sample integration test" > $(PROJECT_NAME)/test/test_sample.go
	@touch $(PROJECT_NAME)/test/fixtures/.gitkeep
	@touch $(PROJECT_NAME)/test/mocks/.gitkeep

	@echo "Structure created for: $(PROJECT_NAME)"

.PHONY: clean
clean:
	@echo "🧹 Cleaning up $(PROJECT_NAME)..."
	@rm -rf $(PROJECT_NAME)
	@echo "Removed project $(PROJECT_NAME)"
