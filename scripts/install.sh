#!/bin/bash

set -e


ROOT_DIR=$(pwd)
CMD_DIR="${ROOT_DIR}/cmd"
DIST_DIR="${ROOT_DIR}/dist"

BINARY_NAME="commitai"
BINARY_PATH="${DIST_DIR}/${BINARY_NAME}"
INSTALL_PATH="/usr/local/bin/${BINARY_NAME}"


echo "==> Creating dist directory..."
mkdir -p "${DIST_DIR}"


echo "==> Changing to cmd directory for compilation..."
cd "${CMD_DIR}"


echo "==> Compiling binary..."
go build -o "${BINARY_PATH}" ./


cd "${ROOT_DIR}"


if [ ! -f "${BINARY_PATH}" ]; then
  echo "Error: Binary '${BINARY_PATH}' was not generated."
  exit 1
fi


chmod +x "${BINARY_PATH}"


echo "==> Installing binary to ${INSTALL_PATH} (may require admin password)..."
sudo cp "${BINARY_PATH}" "${INSTALL_PATH}"

echo "==> Binary '${BINARY_NAME}' successfully installed to ${INSTALL_PATH}."
echo "Use '${BINARY_NAME}' to run the program."
