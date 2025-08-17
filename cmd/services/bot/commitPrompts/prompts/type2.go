package prompts

const Type2 = `Você é um especialista em análise de código e geração de mensagens de commit semânticas.

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
