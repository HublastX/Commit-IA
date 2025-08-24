#!/bin/bash
set -e

ROOT_DIR=$(pwd)
CMD_DIR="${ROOT_DIR}/cmd"
DIST_DIR="${ROOT_DIR}/dist"
BINARY_NAME="commitia"

echo "==> Creating dist directory..."
mkdir -p "${DIST_DIR}"

echo "==> Changing to cmd directory..."
cd "${CMD_DIR}"

echo "==> Compiling Linux/macOS binary..."
go build -o "${DIST_DIR}/${BINARY_NAME}" ./
chmod +x "${DIST_DIR}/${BINARY_NAME}"

cd "${ROOT_DIR}"

if [ ! -f "${DIST_DIR}/${BINARY_NAME}" ]; then
    echo "Error: Linux/macOS binary was not generated."
    exit 1
fi

echo "==> Linux/macOS build completed: ${DIST_DIR}/${BINARY_NAME}"
