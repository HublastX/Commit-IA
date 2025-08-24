.PHONY: help build release npm-pack npm-publish test clean all

VERSION ?= v2.0.0
NPM_VERSION = $(shell echo $(VERSION) | sed 's/^v//')

help: ## Mostrar esta ajuda
	@echo "CommitIA - Comandos disponíveis:"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'

build: ## Construir binários para todas as plataformas
	@echo "🚀 Construindo binários..."
	chmod +x scripts/build-releases.sh
	./scripts/build-releases.sh $(VERSION)

test-local: ## Testar binário local (Linux)
	@echo "🧪 Testando binário local..."
	cp releases/commitia-linux-amd64 ./commitia-test
	chmod +x ./commitia-test
	./commitia-test --help
	rm -f ./commitia-test

release: build ## Fazer release no GitHub (requer gh CLI)
	@echo "📤 Criando release $(VERSION) no GitHub..."
	gh release create $(VERSION) releases/* \
		--title "Release $(VERSION) - Multi-platform" \
		--notes "✨ Suporte completo para Linux, macOS e Windows\n\n📦 **Instalação via NPM:**\n\`\`\`bash\nnpm install -g commit-ia\n\`\`\`\n\n🔧 **Binários disponíveis:**\n- Linux (amd64, arm64)\n- macOS (amd64, arm64) \n- Windows (amd64, arm64)"

npm-pack: ## Gerar pacote NPM
	@echo "📦 Gerando pacote NPM..."
	npm pack
	@echo "✅ Pacote gerado: commit-ia-$(NPM_VERSION).tgz"

npm-test: npm-pack ## Testar instalação NPM local
	@echo "🧪 Testando instalação NPM local..."
	npm uninstall -g commit-ia 2>/dev/null || true
	npm install -g ./commit-ia-$(NPM_VERSION).tgz
	@echo "✅ Testando comando..."
	commitia --help || echo "❌ Teste falhou"

npm-publish: npm-pack ## Publicar no NPM (CUIDADO!)
	@echo "⚠️  ATENÇÃO: Isto irá publicar no NPM público!"
	@echo "Press Ctrl+C para cancelar, Enter para continuar..."
	@read
	npm publish

update-version: ## Atualizar versão no package.json
	@echo "📝 Atualizando versão para $(NPM_VERSION)..."
	npm version $(NPM_VERSION) --no-git-tag-version
	@echo "✅ package.json atualizado"

clean: ## Limpar arquivos temporários
	@echo "🧹 Limpando arquivos temporários..."
	rm -rf releases/
	rm -f *.tgz
	rm -f commitia-test

all: clean update-version build release npm-pack ## Processo completo: build + release + npm

# Comandos de desenvolvimento
dev-install: ## Instalar dependências de desenvolvimento
	npm install

dev-test: ## Testar com binário de desenvolvimento
	@echo "🔧 Testando com binário local..."
	mkdir -p dist
	go build -o dist/commitia .
	chmod +x dist/commitia
	./bin/cli.js --help

# Informações
info: ## Mostrar informações do projeto
	@echo "📋 Informações do projeto:"
	@echo "  Nome: $(shell jq -r '.name' package.json)"
	@echo "  Versão atual: $(shell jq -r '.version' package.json)"
	@echo "  Versão alvo: $(VERSION) -> $(NPM_VERSION)"
	@echo "  Repositório: $(shell jq -r '.repository.url' package.json)"