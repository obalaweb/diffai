# DiffAI Makefile

.PHONY: help build test clean install dev docker-build docker-run

# Default target
help:
	@echo "DiffAI - AI-powered Git assistant"
	@echo ""
	@echo "Available targets:"
	@echo "  build         - Build the Go CLI binary"
	@echo "  test          - Run tests"
	@echo "  clean         - Clean build artifacts"
	@echo "  install       - Install DiffAI CLI"
	@echo "  dev           - Start development environment"
	@echo "  docker-build  - Build Docker images"
	@echo "  docker-run    - Run Docker containers"
	@echo "  lint          - Run linters"
	@echo "  fmt           - Format code"

# Build the Go CLI binary
build:
	@echo "Building DiffAI CLI..."
	go build -o bin/diffai ./cmd/diffai
	@echo "✅ Build complete: bin/diffai"

# Run tests
test:
	@echo "Running tests..."
	go test -v ./...
	@echo "✅ Tests complete"

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -rf bin/
	go clean
	@echo "✅ Clean complete"

# Install DiffAI CLI
install: build
	@echo "Installing DiffAI CLI..."
	sudo cp bin/diffai /usr/local/bin/
	@echo "✅ Installation complete"

# Development environment
dev:
	@echo "Starting development environment..."
	@echo "Starting AI service..."
	cd ai-service && python -m uvicorn app.main:app --reload --host 0.0.0.0 --port 8080 &
	@echo "AI service started on http://localhost:8080"
	@echo "Use 'make dev-stop' to stop the service"

# Stop development environment
dev-stop:
	@echo "Stopping development environment..."
	pkill -f "uvicorn app.main:app" || true
	@echo "✅ Development environment stopped"

# Build Docker images
docker-build:
	@echo "Building Docker images..."
	docker build -t diffai-cli -f Dockerfile .
	docker build -t diffai-ai-service -f ai-service/Dockerfile ai-service/
	@echo "✅ Docker images built"

# Run Docker containers
docker-run:
	@echo "Starting Docker containers..."
	docker-compose up -d
	@echo "✅ Docker containers started"

# Stop Docker containers
docker-stop:
	@echo "Stopping Docker containers..."
	docker-compose down
	@echo "✅ Docker containers stopped"

# Run linters
lint:
	@echo "Running linters..."
	golangci-lint run
	@echo "✅ Linting complete"

# Format code
fmt:
	@echo "Formatting code..."
	go fmt ./...
	@echo "✅ Formatting complete"

# Initialize configuration
config-init:
	@echo "Initializing configuration..."
	./bin/diffai config init
	@echo "✅ Configuration initialized"

# Run with Docker Compose
compose-up:
	@echo "Starting with Docker Compose..."
	docker-compose up --build
	@echo "✅ Services started"

# Run with Docker Compose in background
compose-up-d:
	@echo "Starting with Docker Compose (detached)..."
	docker-compose up -d --build
	@echo "✅ Services started in background"

# Stop Docker Compose
compose-down:
	@echo "Stopping Docker Compose services..."
	docker-compose down
	@echo "✅ Services stopped"

# Show logs
logs:
	@echo "Showing logs..."
	docker-compose logs -f

# Setup development environment
setup:
	@echo "Setting up development environment..."
	go mod tidy
	cd ai-service && pip install -r requirements.txt
	@echo "✅ Development environment setup complete"
