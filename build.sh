#!/bin/sh

GOFILE=main.go

echo "Building for macOS..."
GOOS=darwin GOARCH=amd64 go build -o simpleicons-macos $GOFILE

echo "Building for Linux..."
GOOS=linux GOARCH=amd64 go build -o simpleicons-linux $GOFILE

echo "Build complete. Binaries are located in the current directory."

