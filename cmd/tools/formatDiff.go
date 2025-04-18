package tools

import (
	"fmt"
	"strings"
)

func FormatGitDiff(diffOutput string) string {
	var formattedDiff strings.Builder

	files := strings.Split(diffOutput, "diff --git")
	for _, fileDiff := range files[1:] {
		lines := strings.Split(fileDiff, "\n")
		if len(lines) == 0 {
			continue
		}

		splitLine := strings.Split(lines[0], " b/")
		fileName := "arquivo desconhecido"
		if len(splitLine) > 1 {
			fileName = splitLine[1]
		}

		formattedDiff.WriteString(fmt.Sprintf("# Arquivo: %s\n", fileName))
		formattedDiff.WriteString("## Mudanças\n")

		for _, line := range lines {
			if strings.HasPrefix(line, "+") && !strings.HasPrefix(line, "+++") {
				formattedDiff.WriteString(fmt.Sprintf("+ Adicionado: %s\n", strings.TrimSpace(line[1:])))
			} else if strings.HasPrefix(line, "-") && !strings.HasPrefix(line, "---") {
				formattedDiff.WriteString(fmt.Sprintf("- Removido: %s\n", strings.TrimSpace(line[1:])))
			}
		}

		formattedDiff.WriteString("\n---\n")
	}

	return formattedDiff.String()
}
