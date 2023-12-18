#!/bin/bash

# Остановить скрипт при ошибках
set -e

# Сборка сервисов
echo "Building Driver service..."
go build -o bin/driver cmd/driver/main.go

echo "Building Location service..."
go build -o bin/location cmd/location/main.go

# Сборка Docker-образов (пример)
echo "Building Docker images..."
docker build -t driver-service -f deployments/Dockerfile.driver .
docker build -t location-service -f deployments/Dockerfile.location .

echo "Build completed successfully!"
