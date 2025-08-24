#!/bin/bash

# Build script para CommitIA
# Gera binários para diferentes plataformas

set -e

VERSION=${1:-"v2.0.0"}
OUTPUT_DIR="releases"

echo "🚀 Construindo CommitIA $VERSION"

# Limpar diretório de releases
rm -rf $OUTPUT_DIR
mkdir -p $OUTPUT_DIR

# Função para build
build_binary() {
    local goos=$1
    local goarch=$2
    local output_name=$3
    
    echo "📦 Construindo para $goos/$goarch..."
    
    GOOS=$goos GOARCH=$goarch go build \
        -ldflags "-s -w -X main.Version=$VERSION" \
        -o "$OUTPUT_DIR/$output_name" \
        .
    
    # Comprimir se não for Windows
    if [ "$goos" != "windows" ]; then
        gzip "$OUTPUT_DIR/$output_name"
        mv "$OUTPUT_DIR/$output_name.gz" "$OUTPUT_DIR/$output_name"
    fi
    
    echo "✅ $output_name criado"
}

# Builds para diferentes plataformas
echo "🐧 Linux builds..."
build_binary "linux" "amd64" "commitia-linux-amd64"
build_binary "linux" "arm64" "commitia-linux-arm64"

echo "🍎 macOS builds..."
build_binary "darwin" "amd64" "commitia-darwin-amd64"
build_binary "darwin" "arm64" "commitia-darwin-arm64"

echo "🪟 Windows builds..."
build_binary "windows" "amd64" "commitia-windows-amd64.exe"

echo ""
echo "✅ Todos os binários foram criados em $OUTPUT_DIR/"
ls -la $OUTPUT_DIR/

echo ""
echo "📋 Para testar localmente:"
echo "   cp $OUTPUT_DIR/commitia-linux-amd64 ./commitia"
echo "   chmod +x ./commitia"
echo "   ./commitia --help"

echo ""
echo "📤 Para fazer release no GitHub:"
echo "   gh release create $VERSION $OUTPUT_DIR/* --title \"Release $VERSION\""