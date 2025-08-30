package tools

import (
	"fmt"
	"os"
)

func GetProjectPath() (string, error) {
	projectPath, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("erro ao obter o diretório atual: %v", err)
	}
	return projectPath, nil
}
