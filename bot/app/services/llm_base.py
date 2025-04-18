import os
from typing import Optional

from langchain.schema.language_model import BaseLanguageModel
from langchain_google_genai import ChatGoogleGenerativeAI
from langchain_openai import ChatOpenAI


class LLMBase:
    """
    Classe base para criar e gerenciar modelos de linguagem.
    """

    PROVIDERS = {
        "google": {
            "models": ["gemini-pro", "gemini-1.5-pro", "gemini-2.0-flash"],
            "class": ChatGoogleGenerativeAI,
            "env_var": "GOOGLE_API_KEY",
        },
        "openai": {
            "models": ["gpt-3.5-turbo", "gpt-4", "gpt-4o"],
            "class": ChatOpenAI,
            "env_var": "OPENAI_API_KEY",
        }
    }

    def __init__(
        self,
        provider: str = "google",
        model: str = "gemini-2.0-flash",
        temperature: float = 0.3,
        api_key: Optional[str] = None,
    ):
        self._validate_provider_model(provider, model)

        if api_key:
            os.environ[self.PROVIDERS[provider]["env_var"]] = api_key
        elif not os.getenv(self.PROVIDERS[provider]["env_var"]):
            raise ValueError(
                f"API key não fornecida e variável de ambiente {self.PROVIDERS[provider]['env_var']} não encontrada."
            )

        self.llm = self._create_model(provider, model, temperature)
        self.provider = provider
        self.model = model

    def _validate_provider_model(self, provider: str, model: str) -> None:
        if provider not in self.PROVIDERS:
            available_providers = ", ".join(self.PROVIDERS.keys())
            raise ValueError(
                f"Provedor '{provider}' não suportado. Use um dos seguintes: {available_providers}"
            )

        if model not in self.PROVIDERS[provider]["models"]:
            available_models = ", ".join(self.PROVIDERS[provider]["models"])
            raise ValueError(
                f"Modelo '{model}' não disponível para o provedor '{provider}'. Use um dos seguintes: {available_models}"
            )

    def _create_model(
        self, provider: str, model: str, temperature: float
    ) -> BaseLanguageModel:
        model_class = self.PROVIDERS[provider]["class"]

        if provider == "google":
            return model_class(model=model, temperature=temperature)
        elif provider == "openai":
            return model_class(model=model, temperature=temperature)
        else:
            raise ValueError(
                f"Configuração para provedor '{provider}' não implementada."
            )

    def change_model(
        self,
        provider: str,
        model: str,
        temperature: float = 0.3,
        api_key: Optional[str] = None,
    ) -> None:
        self._validate_provider_model(provider, model)

        if api_key:
            os.environ[self.PROVIDERS[provider]["env_var"]] = api_key

        self.llm = self._create_model(provider, model, temperature)
        self.provider = provider
        self.model = model

    def get_model(self) -> BaseLanguageModel:
        return self.llm
