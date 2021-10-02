#! /bin/ash

go mod tidy >> /dev/null

if [ $? -eq 0 ]; then
    echo "URL: http://localhost:8080/"
    go run main.go
else
    echo "Failed to start"
fi

