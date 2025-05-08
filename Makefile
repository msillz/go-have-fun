# Variables
BINARY_NAME=api-server
CMD_DIR=./cmd/api-server

# Default target: build the binary
build:
	go build -o ./bin/$(BINARY_NAME) $(CMD_DIR)

# Run the application
run:
	go run $(CMD_DIR)

# Clean up the binary
clean:
	rm -f $(BINARY_NAME)

# Phony targets
.PHONY: build run clean