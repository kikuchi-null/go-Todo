#! /bin/ash

go mod tidy >> /dev/null

if [ $? -eq 0 ]; then
    echo "Go Live."
    go run main.go
else
    echo "\"go mod tidy\" is failed"
fi

