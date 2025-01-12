package stackproject

import (
	"fmt"
	"os"
	"strings"
)

func IdentifyProjectLanguages() string {
	fileToLanguage := map[string]string{
		"package.json":  "Node.js",
		"go.mod":        "Go",
		"composer.json": "PHP",
		"pom.xml":       "Java",
	}

	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Erro ao obter o diretório atual: %v\n", err)
		return ""
	}

	files, err := os.ReadDir(currentDir)
	if err != nil {
		fmt.Printf("Erro ao ler o diretório: %v\n", err)
		return ""
	}

	languages := map[string]struct{}{}

	for _, file := range files {
		if file.IsDir() && file.Name() == "venv" {
			languages["Python"] = struct{}{}
		} else if language, exists := fileToLanguage[file.Name()]; exists {
			languages[language] = struct{}{}
		}
	}

	var detected []string
	for lang := range languages {
		detected = append(detected, lang)
	}

	return strings.Join(detected, ", ")
}
