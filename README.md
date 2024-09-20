
# DateTime Server

This project showcases two implementations of an HTTP server that provides the current date and time:

1. A basic implementation using Go's standard `net/http` package
2. An enhanced implementation using the Gin web framework

Both servers return the current date and time in either plain text or JSON format, depending on the `Accept` header in the request.

## Table of Contents

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Running the Servers](#running-the-servers)
- [API Endpoints](#api-endpoints)
- [Docker Support](#docker-support)
- [Makefile Commands](#makefile-commands)

## Features

- Two server implementations: `net/http` and Gin
- Support for plain text and JSON responses
- Docker containerization for easy deployment
- Comprehensive test suite for the client package
- Makefile for automating common tasks




## Prerequisites

- Go 1.16 or higher
- Docker (optional, for containerization)
- Make (optional, for using Makefile commands)

## Installation

1. Clone the repository:
   ```
   git clone https://github.com/MohamedFadel01/datetime-server.git
   cd datetime-server
   ```

2. Install dependencies:
   ```
   go mod download
   ```

## Running the Servers

### Locally

To run the servers locally:

1. Net/HTTP server:
   ```
   go run ./server/nethttp/main.go
   ```

2. Gin server:
   ```
   go run ./server/gin/main.go
   ```

The `net/http` server will be available at `http://localhost:8000`, and the Gin server at `http://localhost:9000`.

### Using Docker

1. Build the Docker images:
   ```
   make docker-nethttp
   make docker-gin
   ```

2. Run the containers:
   ```
   docker run -p 8000:8000 nethttp-datetime-server
   docker run -p 9000:9000 gin-datetime-server
   ```

Alternatively, use Docker Compose to run both servers simultaneously:
```
docker-compose up
```

## API Endpoints

Both servers provide the following endpoints:

- `GET /`: Returns the current datetime in plain text
- `GET /json`: Returns the current datetime in JSON format

Example usage:

```bash
# Plain text response
curl http://localhost:8000/

# JSON response
curl -H "Accept: application/json" http://localhost:8000/json
```

## Docker Support

The project includes Dockerfiles for both server implementations and a `docker-compose.yml` file for easy deployment. See the [Running the Servers](#running-the-servers) section for Docker-related commands.


## Makefile Commands

The project includes a Makefile with the following commands:

- `make all`: Run formatting, linting, and build the project
- `make fmt`: Format the Go code
- `make lint`: Run golangci-lint
- `make build-nethttp`: Build the net/http server binary
- `make build-gin`: Build the Gin server binary
- `make docker-nethttp`: Build the Docker image for the net/http server
- `make docker-gin`: Build the Docker image for the Gin server
- `make docker-compose-up`: Start both servers using Docker Compose
