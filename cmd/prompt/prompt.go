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
	- A messagem tem que ser com palavras curtas diretas ao ponto com falando sobre todas alteraçoes feitas
	- A saída deve ser **apenas a mensagem de commit final**, sem comentários ou explicações adicionais.  
	

	Informações fornecidas:  
	- Linguagem usada no projeto: %s 
	- Descrição básica da mudança: %s 
	- Mudanças detalhadas no comando git diff:  
	
	%s



	`, language, stack, description, diffOutput)

	return prompt
}
