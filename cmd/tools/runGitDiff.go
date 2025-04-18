package tools

import (
	"bytes"
	"fmt"
	"os/exec"
)

func RunGitDiff(projectPath string) (string, error) {
	cmd := exec.Command("git", "diff", "--cached")
	cmd.Dir = projectPath

	var out, stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("error executing git diff command: %s", stderr.String())
	}

	return out.String(), nil
}
