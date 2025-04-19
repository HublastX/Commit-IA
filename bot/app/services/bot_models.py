
from google import genai
from openai import OpenAI
from services.llm_base import LLMBase
from tools.read_commit_config import read_commitia_config

try:
    config = read_commitia_config()
except FileNotFoundError as e:
    print(e)
    config = None


client = None


if config:
    if config.get("provider") == "google" and config.get("api_key"):
        client = genai.Client(api_key=config.get("api_key"))
    elif config.get("provider") == "openai" and config.get("api_key"):
        client = OpenAI(api_key=config.get("api_key"))


def botModel(provider=None, model=None, api_key=None):
    if not config:
        raise ValueError("Configuration not found. Please run 'commitia --update' to configure.")

    provider = provider or config.get("provider")
    if not provider:
        raise ValueError("LLM provider not specified in configuration.")

    model = model or config.get("model")
    if not model:
        raise ValueError(f"LLM model not specified for provider '{provider}' in configuration.")

    api_key = api_key or config.get("api_key")
    if not api_key:
        raise ValueError(f"API key not found for provider '{provider}' in configuration.")

#    print(f"Using provider: {provider}, model: {model}")

    modelo_base = LLMBase(
        provider=provider,
        model=model,
        temperature=0.3,
        api_key=api_key
    )
    return modelo_base
