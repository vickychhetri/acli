#!/bin/sh

mkdir -p dist

echo "Building for Linux..."
GOOS=linux GOARCH=amd64 go build -o dist/acli-linux .

echo "Building for macOS..."
GOOS=darwin GOARCH=amd64 go build -o dist/acli-mac .

echo "Building for Windows..."
GOOS=windows GOARCH=amd64 go build -o dist/acli.exe .

echo "Done! Files in dist/"
