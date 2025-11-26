# ====================================================================
# Makefile for the Art Decoder/Encoder CLI Tool
# ====================================================================

# Set the executable name (e.g., 'art-decoder')
TARGET = art-decoder

# Define the root of the Go module
MODULE_ROOT := .

# --- Basic Targets ---

.PHONY: all build run tidy clean help

all: build

## build: Compiles the Go source code into an executable.
build: tidy
	@echo "Building $(TARGET)..."
	go build -o $(TARGET) $(MODULE_ROOT)

## run: Runs the compiled program with command-line arguments.
# Use this target like: make run ARGS="--multi \"[5 #]A\nB\""
run: build
	@echo "Running $(TARGET)..."
	./$(TARGET) $(ARGS)

## tidy: Cleans up module dependencies and synchronizes go.mod/go.sum.
tidy:
	@echo "Tidying Go module..."
	go mod tidy

## test: Runs the project's Go tests (assuming tests exist in _test.go files).
test:
	@echo "Running tests..."
	go test ./...

## clean: Removes the compiled executable file.
clean:
	@echo "Cleaning compiled executable..."
	@rm -f $(TARGET)

## help: Displays available targets and their descriptions.
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-10s\033[0m %s\n", $$1, $$2}'