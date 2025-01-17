package prompt

import (
	"fmt"
	"strings"
)

func FormatGitDiff(diffOutput string) string {
	var formattedDiff strings.Builder

	files := strings.Split(diffOutput, "diff --git")
	for _, fileDiff := range files[1:] {
		lines := strings.Split(fileDiff, "\n")
		if len(lines) == 0 {
			continue
		}

		splitLine := strings.Split(lines[0], " b/")
		fileName := "arquivo desconhecido"
		if len(splitLine) > 1 {
			fileName = splitLine[1]
		}

		formattedDiff.WriteString(fmt.Sprintf("# Arquivo: %s\n", fileName))
		formattedDiff.WriteString("## Mudanças\n")

		for _, line := range lines {
			if strings.HasPrefix(line, "+") && !strings.HasPrefix(line, "+++") {
				formattedDiff.WriteString(fmt.Sprintf("+ Adicionado: %s\n", strings.TrimSpace(line[1:])))
			} else if strings.HasPrefix(line, "-") && !strings.HasPrefix(line, "---") {
				formattedDiff.WriteString(fmt.Sprintf("- Removido: %s\n", strings.TrimSpace(line[1:])))
			}
		}

		formattedDiff.WriteString("\n---\n")
	}

	return formattedDiff.String()
}

func CreateCommitMessage(diffOutput, language, description string) string {
	formattedDiff := FormatGitDiff(diffOutput)

	prompt := fmt.Sprintf(`
        Com base nas informações abaixo, crie uma mensagem de commit no padrão Conventional Commits, que utiliza prefixos específicos para categorizar o tipo de mudança seguido de uma descrição breve.

        Prefixos mais usados:
        - feat: Adição de nova funcionalidade
        - fix: Correção de bugs
        - chore: Alterações menores ou de manutenção sem impacto na funcionalidade

        Regras obrigatórias:
        - A mensagem deve ser baseada no que foi alterado no git diff e descrever as ações feitas.
        - Use frases curtas e diretas, separando as ações realizadas no commit por vírgulas.
        - **Fale das alterações do commit, separando com vírgulas e de forma simples e curta.**
        - Não fale o caminho dos arquivos alterados ou criados.
        - Não invente mensagens ou coisas que não têm no commit.
        - **Não use caracteres especiais como crase ou aspas.**
        - A primeira linha deve começar com o prefixo correto (feat, fix, chore).
        - Não inclua caminhos completos dos arquivos, apenas nomes principais se necessário.
        - Utilize obrigatoriamente o idioma **%s** na resposta.
        - Escreva a mensagem como se fosse em primeira pessoa.
        - A mensagem tem que ser com palavras curtas, diretas ao ponto, e mencionar todas as alterações feitas.
        - A saída deve ser **apenas a mensagem de commit final**, sem comentários ou explicações adicionais.

        Informações fornecidas:
        - Idioma do commit: %s
        - Descrição básica da mudança: %s
        - Mudanças detalhadas no comando git diff:

        %s

        Exemplo de mensagem esperada:
        - "feat: Criando sistema de login, realizado função para ordenar arquivos, chore: melhoria nomes das variáveis"

        Escreva a mensagem do commit usando poucas palavras e no idioma: **%s**.
`, language, language, description, formattedDiff, language)

	return prompt
}
