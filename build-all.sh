#!/bin/bash

set -e

ROOT_DIR=$(pwd)
CMD_DIR="${ROOT_DIR}/cmd"
DIST_DIR="${ROOT_DIR}/dist"
BINARY_NAME="commitia"

echo "==> Creating dist directory..."
mkdir -p "${DIST_DIR}"

echo "==> Changing to cmd directory for compilation..."
cd "${CMD_DIR}"

echo "==> Building for multiple platforms..."

# Linux x64
echo "Building for Linux x64..."
GOOS=linux GOARCH=amd64 go build -o "${DIST_DIR}/${BINARY_NAME}-linux-amd64" ./

# Windows x64
echo "Building for Windows x64..."
GOOS=windows  go build -o "${DIST_DIR}/${BINARY_NAME}-windows-amd64.exe" ./

# macOS x64
echo "Building for macOS x64..."
GOOS=darwin GOARCH=amd64 go build -o "${DIST_DIR}/${BINARY_NAME}-darwin-amd64" ./

# macOS ARM64 (M1/M2)
echo "Building for macOS ARM64..."
GOOS=darwin GOARCH=arm64 go build -o "${DIST_DIR}/${BINARY_NAME}-darwin-arm64" ./

cd "${ROOT_DIR}"

echo "==> All binaries built successfully:"
ls -la "${DIST_DIR}/"

echo ""
echo "To test on your current platform:"
case "$(uname -s)" in
    "Linux")
        echo "  ./dist/${BINARY_NAME}-linux-amd64"
        ;;
    "Darwin")
        if [[ "$(uname -m)" == "arm64" ]]; then
            echo "  ./dist/${BINARY_NAME}-darwin-arm64"
        else
            echo "  ./dist/${BINARY_NAME}-darwin-amd64"
        fi
        ;;
    *)
        echo "  Use the appropriate binary for your platform"
        ;;
esac
