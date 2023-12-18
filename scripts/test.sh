#!/bin/bash

# Остановить скрипт при ошибках
set -e

# Запуск тестов
echo "Running tests..."
go test ./...

echo "Tests completed successfully!"
