.PHONY: all build fmt lint docker-build nethttp-build gin-build docker-compose-up clean

NETHTTP_DIR=./server/nethttp
GIN_DIR=./server/gin

all: fmt lint build docker-build docker-compose-up

fmt:
	@echo "Running go fmt..."
	@go fmt ./...

lint:
	@echo "Running golangci-lint..."
	@golangci-lint run ./...

build: nethttp-build gin-build

nethttp-build:
	@echo "Building net/http server binary..."
	@cd $(NETHTTP_DIR) && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o datetime-server .

gin-build:
	@echo "Building Gin server binary..."
	@cd $(GIN_DIR) && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o datetime-server .

docker-build:
	@echo "Building Docker images for both services..."
	@sudo docker build -t nethttp-datetime-server -f $(NETHTTP_DIR)/Dockerfile .
	@sudo docker build -t gin-datetime-server -f $(GIN_DIR)/Dockerfile .

docker-compose-up:
	@echo "Launching services using Docker Compose..."
	@sudo docker-compose up --build

clean:
	@echo "Cleaning up binaries..."
	@rm -f $(NETHTTP_DIR)/datetime-server
	@rm -f $(GIN_DIR)/datetime-server
