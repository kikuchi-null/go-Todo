#!/bin/ash

echo "==>Checking..."
go mod tidy >/dev/null 2>&1

if [ $? -eq 0 ]; then
    echo "==>Go live (http://localhost:8080/)"
    go run main.go
else
    echo "==>Failed to start"
fi

