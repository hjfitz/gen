# Variables
APP_NAME := gen
BUILD_DIR := bin
MAIN_FILE := main.go
INSTALL_DIR := $(HOME)/.bin

# Default target
.PHONY: all
all: clean build

# Build the application
.PHONY: build
build:
	@echo "Building the application..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_FILE)

# Run the application
.PHONY: run
run:
	@echo "Running the application..."
	@go run $(MAIN_FILE)

# Clean up build artifacts
.PHONY: clean
clean:
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)

# Install dependencies
.PHONY: deps
deps:
	@echo "Installing dependencies..."
	@go mod tidy

# Format the code
.PHONY: fmt
fmt:
	@echo "Formatting code..."
	@go fmt ./...

# Install the application
.PHONY: install
install: clean build
	@echo "Installing the application to $(INSTALL_DIR)..."
	@mkdir -p $(INSTALL_DIR)
	@cp $(BUILD_DIR)/$(APP_NAME) $(INSTALL_DIR)/
	@echo "Installation complete!"

# Show help
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  all     - Build the application (default)"
	@echo "  build   - Build the application"
	@echo "  run     - Run the application"
	@echo "  clean   - Clean up build artifacts"
	@echo "  deps    - Install dependencies"
	@echo "  fmt     - Format the code"
	@echo "  help    - Show this help message"
