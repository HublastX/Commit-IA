package services

import (
	"fmt"

	"github.com/HublastX/Commit-IA/tools"
)

func CreateCommitMessage(diffOutput, language, description, tagCommit string) string {
	formattedDiff := tools.FormatGitDiff(diffOutput)

	prompt := fmt.Sprintf(`

    Informações fornecidas:
        - Utilize obrigatoriamente o idioma **%s** na resposta.
        - Tag do commit que deve ser utilizada: %s
        - Descrição básica da mudança: %s
        - Mudanças detalhadas no comando git diff:

        %s

        Escreva a mensagem do commit usando poucas palavras e no idioma: **%s**.
`, language, tagCommit, description, formattedDiff, language)

	return prompt
}
