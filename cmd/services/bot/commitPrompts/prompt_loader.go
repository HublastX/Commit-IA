package commitprompts

import (
	"fmt"
)

const PromptType1 = `Você é um especialista em análise de código e geração de mensagens de commit semânticas.

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
8. Se emoji estiver habilitado: escolha o emoji apropriado e coloque DEPOIS da tag, formato: "tag(arquivo_ou_pasta) emoji: título"
9. Se emoji NÃO estiver habilitado: formato padrão "tag(arquivo_ou_pasta_principal): título conciso da mudança"

Retorne APENAS a mensagem de commit formatada, sem explicações adicionais.

Exemplos SEM emoji:
- feat(controller): implementa autenticacao de usuarios
- fix(utils): corrige calculo de data incorreto
- docs(readme): atualiza instrucoes de instalacao

Exemplos COM emoji:
- feat(controller) ✨: implementa autenticacao de usuarios
- fix(utils) 🐛: corrige calculo de data incorreto
- docs(readme) 📝: atualiza instrucoes de instalacao`

const PromptType2 = `Você é um especialista em análise de código e geração de mensagens de commit semânticas.

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
4. Crie um título conciso que explique a mudança (não mais que 50 caracteres)
5. Use o idioma especificado pelo usuário
6. Regras de formatação importantes:
   - use apenas letras minúsculas na mensagem de commit
   - não use aspas (simples ou duplas) na mensagem
   - evite caracteres especiais como \ ou outros símbolos não essenciais
7. Se emoji estiver habilitado: escolha o emoji apropriado e coloque DEPOIS da tag, formato: "tag emoji: título"
8. Se emoji NÃO estiver habilitado: formato padrão "tag: título conciso da mudança"

Retorne APENAS a mensagem de commit formatada, sem explicações adicionais.

Exemplos SEM emoji:
- feat: implementa autenticacao de usuarios
- fix: corrige calculo de data incorreto
- docs: atualiza instrucoes de instalacao

Exemplos COM emoji:
- feat ✨: implementa autenticacao de usuarios
- fix 🐛: corrige calculo de data incorreto
- docs 📝: atualiza instrucoes de instalacao`

const PromptType3 = `Você é um especialista em análise de código e geração de mensagens de commit.

Analise as seguintes modificações de código e gere uma mensagem de commit natural:

MODIFICAÇÕES:
%s

INFORMAÇÕES ADICIONAIS:
- Idioma para o commit: %s
- Descrição curta fornecida pelo usuário: %s

INSTRUÇÕES:
1. Analise cuidadosamente as modificações para entender o propósito da mudança
2. Crie uma mensagem natural e descritiva que explique o que foi feito
3. Use o idioma especificado pelo usuário
4. Seja conciso mas claro (não mais que 72 caracteres)
5. Regras de formatação importantes:
   - use apenas letras minúsculas na mensagem de commit
   - não use aspas (simples ou duplas) na mensagem
   - evite caracteres especiais como \ ou outros símbolos não essenciais
6. NÃO use tags semânticas (feat:, fix:, etc)
7. Se emoji estiver habilitado: coloque o emoji apropriado NO INÍCIO da mensagem
8. Se emoji NÃO estiver habilitado: mensagem natural sem emoji
9. Escreva como se fosse uma descrição natural do que foi implementado

Retorne APENAS a mensagem de commit formatada, sem explicações adicionais.

Exemplos SEM emoji:
- implementa sistema de autenticacao de usuarios
- corrige calculo incorreto de datas
- atualiza documentacao de instalacao

Exemplos COM emoji:
- ✨ implementa sistema de autenticacao de usuarios
- 🐛 corrige calculo incorreto de datas
- 📝 atualiza documentacao de instalacao`

func GetPrompt(promptType int, useEmoji bool) (string, error) {
	var basePrompt string

	switch promptType {
	case 1:
		basePrompt = PromptType1
	case 2:
		basePrompt = PromptType2
	case 3:
		basePrompt = PromptType3
	default:
		basePrompt = PromptType1
	}

	if useEmoji {
		emojiAddition, err := GetEmojiPromptAddition()
		if err != nil {
			return "", fmt.Errorf("error loading emoji data: %v", err)
		}
		basePrompt += emojiAddition
	}

	return basePrompt, nil
}

const BaseCustomPrompt = `Você é um especialista em análise de código e geração de mensagens de commit.

Analise as seguintes modificações de código e gere uma mensagem de commit seguindo o formato específico solicitado pelo usuário:

MODIFICAÇÕES:
%s

INFORMAÇÕES ADICIONAIS:
- Idioma para o commit: %s
- Descrição curta fornecida pelo usuário: %s
- Tag sugerida pelo usuário: %s

FORMATO PERSONALIZADO SOLICITADO PELO USUÁRIO:
%s

INSTRUÇÕES:
1. Analise cuidadosamente as modificações para entender o propósito da mudança
2. Use o idioma especificado pelo usuário
3. Considere a descrição e tag fornecidas pelo usuário, mas adapte conforme necessário
4. Siga EXATAMENTE o formato personalizado solicitado pelo usuário
5. Se emoji estiver habilitado: escolha o emoji apropriado e integre no formato personalizado (não substitua partes do formato, apenas adicione o emoji)
6. Regras de formatação importantes:
   - use apenas letras minúsculas na mensagem de commit
   - não use aspas (simples ou duplas) na mensagem
   - evite caracteres especiais como \ ou outros símbolos não essenciais

Retorne APENAS a mensagem de commit formatada conforme o formato personalizado solicitado, sem explicações adicionais.`

func GetCustomPrompt(customFormatText string, useEmoji bool) (string, error) {
	if customFormatText == "" {
		return "", fmt.Errorf("custom format text is empty")
	}

	basePrompt := BaseCustomPrompt

	if useEmoji {
		emojiAddition, err := GetEmojiPromptAddition()
		if err != nil {
			return "", fmt.Errorf("error loading emoji data: %v", err)
		}
		basePrompt += emojiAddition
	}

	return basePrompt, nil
}
