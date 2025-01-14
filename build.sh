#!/bin/bash


BINARY_NAME="commitia"


DIST_DIR="dist"


if [ ! -d "$DIST_DIR" ]; then
  mkdir "$DIST_DIR"
fi


echo "Compilando o binário..."
go build -o "$BINARY_NAME"


if [ -f "$BINARY_NAME" ]; then
  mv "$BINARY_NAME" "$DIST_DIR/"
else
  echo "Erro: O binário '$BINARY_NAME' não foi gerado."
  exit 1
fi

