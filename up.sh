#!/bin/sh

# Exits immediately if a command exits with a non-zero status
set -e

# Run 'docker-compose up' for source database deployment
docker-compose -f docker/db/docker-compose.yml up -d

# Build the Go application
go build -o OAuth2-authorization-code-flow-using-Google src/main.go

# Run the application
./OAuth2-authorization-code-flow-using-Google