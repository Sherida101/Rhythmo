# Makefile for Rhythmo (Go CLI + Spring Boot Web + Docker)

# Paths
WEB_DIR=web
CLI_DIR=cli

# Env
ENV_FILE=$(CLI_DIR)/.env

# Java
JAVA_BUILD_CMD=mvn -f $(WEB_DIR)/pom.xml clean package -DskipTests
JAVA_RUN_JAR=$(WEB_DIR)/target/*.jar

# Go
GO_MAIN=$(CLI_DIR)/main.go
GO_RUN_CMD=go run $(GO_MAIN)
GO_BUILD_CMD=go build -o $(CLI_DIR)/rhythmo-cli $(GO_MAIN)

# Docker
DOCKER_COMPOSE=docker-compose

# Default target
.PHONY: all
all: help

.PHONY: help
help:
	@echo "🛠  Rhythmo Dev Toolkit"
	@echo "Available targets:"
	@echo "  make web         → Build Spring Boot web app"
	@echo "  make run-web     → Run Spring Boot web app"
	@echo "  make cli         → Run Go CLI"
	@echo "  make build-cli   → Build Go CLI binary"
	@echo "  make docker-up   → Start all services (Docker)"
	@echo "  make docker-down → Stop services"
	@echo "  make tidy        → Run go mod tidy"
	@echo "  make lint        → Run Go linters"
	@echo "  make clean       → Clean builds"

.PHONY: web
web:
	@echo "📦 Building Spring Boot web app..."
	@$(JAVA_BUILD_CMD)

.PHONY: run-web
run-web: web
	@echo "🚀 Running Spring Boot..."
	@java -jar $(JAVA_RUN_JAR)

.PHONY: cli
cli:
	@echo "🧠 Running Go CLI..."
	@cd $(CLI_DIR) && go run main.go

.PHONY: build-cli
build-cli:
	@echo "🛠 Building Go CLI..."
	@cd $(CLI_DIR) && go build -o rhythmo-cli main.go

.PHONY: tidy
tidy:
	@echo "🔧 Running go mod tidy..."
	@cd $(CLI_DIR) && go mod tidy

.PHONY: lint
lint:
	@echo "🔍 Running golint..."
	@cd $(CLI_DIR) && golint ./...

.PHONY: docker-up
docker-up:
	@echo "🐳 Starting Docker services..."
	@$(DOCKER_COMPOSE) up --build

.PHONY: docker-down
docker-down:
	@echo "🛑 Stopping Docker services..."
	@$(DOCKER_COMPOSE) down

.PHONY: clean
clean:
	@echo "🧹 Cleaning up..."
	@rm -f $(CLI_DIR)/rhythmo-cli
	@rm -rf $(WEB_DIR)/target
