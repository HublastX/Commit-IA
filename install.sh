#!/bin/bash


COMMITIA_PATH="dist/commitia"


if [ ! -f "$COMMITIA_PATH" ]; then
    echo "Erro: O arquivo '$COMMITIA_PATH' não foi encontrado."
    exit 1
fi


chmod +x "$COMMITIA_PATH"


sudo mv "$COMMITIA_PATH" /usr/local/bin/commitia


if [ $? -eq 0 ]; then
    echo "O binário 'commitia' foi movido para /usr/local/bin e está pronto para uso."
else
    echo "Erro ao mover o arquivo para /usr/local/bin."
    exit 1
fi
