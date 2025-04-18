package configpath

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	schemas "github.com/HublastX/Commit-IA/schema"
)

func LoadConfig() (*schemas.LLMConfig, error) {
	configPath, err := getConfigFilePath()
	if err != nil {
		return nil, err
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, nil
	}

	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %v", err)
	}

	var config schemas.LLMConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("error parsing config file: %v", err)
	}

	return &config, nil
}
