#!/bin/bash

# Запуск сервисов через Docker Compose
echo "Starting services with Docker Compose..."
docker-compose -f deployments/docker-compose.yml up

# Или запуск локально (пример)
# echo "Starting Driver service..."
# ./bin/driver

# echo "Starting Location service..."
# ./bin/location
