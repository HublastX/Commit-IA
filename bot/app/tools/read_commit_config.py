
import json
import os
from typing import Dict


def read_commitia_config() -> Dict:
    """
    Read the commitia config file from config/config.json (inside /app in Docker).
    """

    config_path = os.path.join(os.path.dirname(__file__), "../config/config.json")
    config_path = os.path.abspath(config_path)
    if not os.path.exists(config_path):
        raise FileNotFoundError(
            f"Configuration not found at {config_path}. Please run 'commitia --update' to configure."
        )
    with open(config_path, 'r') as f:
        config = json.load(f)
    return config
