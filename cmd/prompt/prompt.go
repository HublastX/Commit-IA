package prompt

import "fmt"

func CreateCommitMessage(diffOutput, language, description, stack string) string {
	prompt := fmt.Sprintf(`
	Com base nas informações fornecidas abaixo, crie uma mensagem de commit seguindo o padrão de Conventional Commits, 
	que é amplamente adotado para tornar as mensagens de commit mais descritivas e úteis. Este padrão utiliza prefixos 
	específicos para categorizar o tipo de mudança realizada, seguido de uma breve descrição. 

	Os prefixos do padrão de Conventional Commits mais comuns são:
		- feat: Uma nova funcionalidade
		- fix: Uma correção de bug
		- chore: Mudanças de manutenção ou pequenas correções que não alteram a funcionalidade

	A título de informação e para seu melhor entendimento, o projeto em questão utiliza a linguagem de programação %s. 

	Descrição básica da mudança fornecida pelo programador a qual você deve usar como base para o início da sua mensagem: '%s'

	Abaixo seguem as mudanças detalhadas (incluindo arquivos modificados e o que foi incluído, alterado ou excluído) geradas pelo comando 'git diff': 

	%s

	Com base nas informações acima, melhore a descrição básica para criar uma mensagem de commit objetiva.

	Regras obrigatórias:
	- Você deve seguir o padrão de Conventional Commits.
	- A primeira linha da mensagem deve obrigatoriamente começar com um dos prefixos do padrão de Conventional Commits (feat, fix, chore, etc.) seguido de uma descrição concisa que explique o que foi feito
	- Após a primeira linha, sempre adicione uma explicação objetiva das alterações realizadas, o motivo da mudança e, se aplicável, o impacto da mudança.
	- Sempre que possível, cite na mensagem de commit somente os arquivos principais modificados sem incluir o path.
	- NÃO adicione comentários ou explicações adicionais além da mensagem de commit gerada.
	- NÃO utilize símbolos como  ou qualquer outra formatação para identificar a mensagem de commit.
	- NÃO adicione quebras de linha ou espaços em branco antes da mensagem de commit.
	- A saída deve ser APENAS a mensagem de commit final conforme as instruções.
	- Utilize obrigatoriamente o idioma '%s' na geração de sua resposta.

	-- Mande messagens curtas e diretas sobre o assunto
	
	- Responda de forma curta use poucas palavras e bem especifico somente no que no que o usuario perguntou. 
	`, stack, description, diffOutput, language)

	return CallProviderAPI(prompt)
}

func CallProviderAPI(prompt string) string {
	messages := []map[string]string{
		{
			"role": "system",
			"content": `
				Você é um assistente que ajuda a gerar mensagens de commit para um repositório Git. 
				As mensagens de commit devem seguir o padrão de Conventional Commits, que usa prefixos específicos para categorizar o tipo de mudança realizada (feat, fix, chore, etc.). 
				A descrição deve ser concisa e clara, explicando o que foi feito, o motivo da mudança e, se aplicável, o impacto da mudança.
				As mensagens devem ser geradas com base nas alterações fornecidas pelo comando 'git diff' e em uma descrição básica fornecida pelo usuário opcionalmente. 

				Regras obrigatórias:
				- NÃO adicione comentários ou explicações adicionais além da mensagem de commit gerada.
				- NÃO utilize símbolos como   ou qualquer outra formatação para identificar a mensagem de commit.
				- NÃO adicione quebras de linha ou espaços em branco antes da mensagem de commit.
				- A saída deve ser APENAS a mensagem de commit final conforme as instruções.
				- A primeira linha da mensagem deve obrigatoriamente começar com um dos prefixos do padrão de Conventional Commits (feat, fix, chore, etc.).
				- Após a primeira linha, sempre adicione uma explicação objetiva das alterações realizadas, o motivo da mudança e, se aplicável, o impacto da mudança.

				Se as instruções não forem seguidas corretamente, o resultado não será aceito.

				Mande messagens curtas e diretas sobre o assunto

				- Responda de forma curta use poucas palavras e bem especifico somente no que no que o usuario perguntou. 
			`,
		},
		{"role": "user", "content": prompt},
	}

	return messages[1]["content"]
}
