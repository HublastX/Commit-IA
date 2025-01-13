package prompt

import "fmt"

func CreateCommitMessage(diffOutput, language, description, stack string) string {
	prompt := fmt.Sprintf(`
	Com base nas informações abaixo, crie uma mensagem de commit no padrão Conventional Commits, que utiliza prefixos específicos para categorizar o tipo de mudança seguido de uma descrição breve.  

	Prefixos mais usados:  
	- feat: Adição de nova funcionalidade  
	- fix: Correção de bugs  
	- chore: Alterações menores ou de manutenção sem impacto na funcionalidade  

	Regras obrigatórias:  
	- A mensagem deve ser baseada no que foi alterado no git diff e descrever as ações feitas.  
	- Use frases curtas e diretas, separando as ações realizadas no commit por vírgulas.
	-  ** Fale das alterações do commit, seprando com virgulas e simples e curta **
	- Não fale o caminho dos arquivos alterados ou criados
	- Não Invente messagens ou coisas que nao tem no commit
	-  **Não Use caracter especial como Craze ou aspas**
	- A primeira linha deve começar com o prefixo correto (feat, fix, chore,). 
	- Não inclua caminhos completos dos arquivos, apenas nomes principais se necessário.  
	- Utilize obrigatoriamente o idioma ** %s ** na resposta.  
	- Esceve a messagem como fosse em primeira pessoa
	- A saída deve ser **apenas a mensagem de commit final**, sem comentários ou explicações adicionais.  
	

	Informações fornecidas:  
	- Linguagem usada no projeto: %s 
	- Descrição básica da mudança: %s 
	- Mudanças detalhadas no comando git diff:  
	
	%s



	`, language, stack, description, diffOutput)

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

				- Responda com poucas palavras, uma messagem curta de direto no maximo 40 caracter
				
			`,
		},
		{"role": "user", "content": prompt},
	}

	return messages[1]["content"]
}
