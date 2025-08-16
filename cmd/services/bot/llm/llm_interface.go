package llm

import (
	"fmt"
	"strings"
)

type LLMClient interface {
	GenerateCommitMessage(prompt string) (string, error)
	GetProvider() string
	GetModel() string
	SetAPIKey(apiKey string)
}

type CommitAnalyzer struct {
	client LLMClient
}

const PromptTemplate = `
Você é um especialista em análise de código e geração de mensagens de commit semânticas.

Analise as seguintes modificações de código e gere uma mensagem de commit semântica:

MODIFICAÇÕES:
%s

INFORMAÇÕES ADICIONAIS:
- Idioma para o commit: %s
- Descrição curta fornecida pelo usuário: %s
- Tag sugerida pelo usuário (use se apropriada): %s

INSTRUÇÕES:
1. Analise cuidadosamente as modificações para entender o propósito da mudança
2. Determine a tag semântica mais apropriada entre:
   - feat: Uma nova funcionalidade
   - fix: Correção de um bug
   - docs: Alterações em documentação
   - style: Mudanças que não afetam o significado do código (formatação, etc)
   - refactor: Refatoração de código existente
   - perf: Melhorias de desempenho
   - test: Adição ou correção de testes
   - chore: Mudanças no processo de build, ferramentas, etc
   - revert: Reversão de um commit anterior
   - cleanup: Remoção de código obsoleto ou limpeza
   - build: Mudanças que afetam o sistema de build ou dependências externas
   - remover: Remoção de funcionalidades ou arquivos
3. Se o usuário forneceu uma tag, utilize-a apenas se for coerente com as modificações
4. Identifique a pasta principal ou arquivo principal que foi modificado
5. Crie um título conciso que explique a mudança (não mais que 50 caracteres)
6. Use o idioma especificado pelo usuário
7. Regras de formatação importantes:
   - use apenas letras minúsculas na mensagem de commit
   - não use aspas (simples ou duplas) na mensagem
   - evite caracteres especiais como \ ou outros símbolos não essenciais
8. Sempre retorne a msg do commit nesse formato "tag(arquivo_ou_pasta_principal): título conciso da mudança"

Retorne APENAS a mensagem de commit formatada, sem explicações adicionais, no seguinte formato:
tag(pasta_ou_arquivo_principal): título conciso da mudança

Exemplos:
- feat(controller): implementa autenticacao de usuarios
- fix(utils): corrige calculo de data incorreto
- docs(readme): atualiza instrucoes de instalacao
`

func NewCommitAnalyzer(provider, model, apiKey string) (*CommitAnalyzer, error) {
	var client LLMClient

	switch provider {
	case "openai":
		client = NewOpenAIClient(model)
	case "google":
		client = NewGoogleClient(model)
	default:
		return nil, fmt.Errorf("unsupported provider: %s", provider)
	}

	if apiKey != "" {
		client.SetAPIKey(apiKey)
	}

	return &CommitAnalyzer{
		client: client,
	}, nil
}

func (ca *CommitAnalyzer) AnalyzeCommit(codeChanges, description, tag, language string) (string, error) {
	prompt := fmt.Sprintf(PromptTemplate, codeChanges, language, description, tag)

	response, err := ca.client.GenerateCommitMessage(prompt)
	if err != nil {
		return "", fmt.Errorf("failed to generate commit message: %v", err)
	}

	return strings.TrimSpace(response), nil
}

func (ca *CommitAnalyzer) GetProvider() string {
	return ca.client.GetProvider()
}

func (ca *CommitAnalyzer) GetModel() string {
	return ca.client.GetModel()
}
