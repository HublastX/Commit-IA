import json
import os
from typing import Dict

from google import genai
from openai import OpenAI
from services.llm_base import LLMBase


# filepath: /home/wendellast/Documents/Github/hublast/Commit-IA/bot/app/services/bot_models.py
def read_commitia_config() -> Dict:
    """
    Read the commitia config file from config/config.json (dentro do /app no Docker).
    """
    # Caminho absoluto para config/config.json relativo ao diretório de trabalho
    config_path = os.path.join(os.path.dirname(__file__), "../config/config.json")
    config_path = os.path.abspath(config_path)
    if not os.path.exists(config_path):
        raise FileNotFoundError(
            f"Configuração não encontrada em {config_path}. Por favor, execute 'commitia --update' para configurar."
        )
    with open(config_path, 'r') as f:
        config = json.load(f)
    return config

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
        raise ValueError("Configuração não encontrada. Por favor, execute 'commitia --update' para configurar.")

    provider = provider or config.get("provider")
    if not provider:
        raise ValueError("LLM provider não especificado na configuração.")

    model = model or config.get("model")
    if not model:
        raise ValueError(f"LLM model não especificado para o provider '{provider}' na configuração.")

    api_key = api_key or config.get("api_key")



    if not api_key:
        raise ValueError(f"API key não encontrada para o provider '{provider}' na configuração.")

    print(f"Usando provider: {provider}, model: {model}")

    modelo_base = LLMBase(
        provider=provider,
        model=model,
        temperature=0.3,
        api_key=api_key
    )
    return modelo_base


