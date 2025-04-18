from typing import Optional

from langchain.prompts import ChatPromptTemplate
from services.llm_base import LLMBase


class CommitAnalyzer:
    """
    Class to analyze code changes and generate semantic commit messages.
    """

    PROMPT_TEMPLATE = """
    Você é um especialista em análise de código e geração de mensagens de commit semânticas.

    Analise as seguintes modificações de código e gere uma mensagem de commit semântica:

    MODIFICAÇÕES:
    {changes}

    INFORMAÇÕES ADICIONAIS:
    - Idioma para o commit: {language}
    - Descrição curta fornecida pelo usuário: {description}
    - Tag sugerida pelo usuário (use se apropriada): {tag}

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
       - evite caracteres especiais como \\ ou outros símbolos não essenciais
    8. Sempre retorne a msg do commit nesse formato "tag(arquivo_ou_pasta_principal): título conciso da mudança"

    Retorne APENAS a mensagem de commit formatada, sem explicações adicionais, no seguinte formato:
    tag(pasta_ou_arquivo_principal): título conciso da mudança

    Exemplos:
    - feat(controller): implementa autenticacao de usuarios
    - fix(utils): corrige calculo de data incorreto
    - docs(readme): atualiza instrucoes de instalacao
    """

    def __init__(self, llm_model: Optional[LLMBase] = None):
        """
        Initializes the commit analyzer.

        Args:
            llm_model (LLMBase, optional): LLM model instance. If not provided, creates a new one.
        """
        self.llm_model = llm_model if llm_model else LLMBase()
        self.prompt_template = self.PROMPT_TEMPLATE

    def update_prompt(self, new_prompt: str) -> None:
        """
        Updates the prompt template used to generate commit messages.

        Args:
            new_prompt (str): New prompt template
        """
        self.prompt_template = new_prompt

    async def generate_commit_message(
        self,
        code_changes: str,
        language: str = "português",
        description: str = "",
        tag: str = ""
    ) -> str:
        """
        Analyzes code changes and generates a semantic commit message.

        Args:
            code_changes (str): The git diff or code changes to analyze
            language (str): Language to write the commit message in (default: "português")
            description (str): Short description provided by the user (optional)
            tag (str): Commit tag suggested by the user (optional)

        Returns:
            str: Semantic commit message
        """
        prompt = ChatPromptTemplate.from_template(self.prompt_template)

        message = prompt.format_messages(
            changes=code_changes,
            language=language,
            description=description,
            tag=tag
        )


        response = self.llm_model.get_model().invoke(message)


        return response.content.strip()
