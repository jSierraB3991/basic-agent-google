package git

import (
	"fmt"
	"os/exec"
	"strings"
)

// RunGitCommand ejecuta un comando git con los argumentos dados en el directorio de trabajo especificado.
func RunGitCommand(workingDir string, args ...string) string {
	if workingDir == "" {
		workingDir = "."
	}

	cmdArgs := append([]string{}, args...)
	cmd := exec.Command("git", cmdArgs...)
	cmd.Dir = workingDir

	outputBytes, err := cmd.CombinedOutput()
	outputStr := string(outputBytes)

	if err != nil {
		if strings.TrimSpace(outputStr) != "" {
			return fmt.Sprintf("Error ejecutando 'git %s':\n%s", strings.Join(args, " "), outputStr)
		}
		return fmt.Sprintf("Error ejecutando 'git %s': %v", strings.Join(args, " "), err)
	}

	if strings.TrimSpace(outputStr) == "" {
		return fmt.Sprintf("Comando 'git %s' ejecutado exitosamente (sin salida).", strings.Join(args, " "))
	}

	return outputStr
}

// GitStatus ejecuta 'git status'
func GitStatus(workingDir string) string {
	return RunGitCommand(workingDir, "status")
}

// GitDiff ejecuta 'git diff' o diff de un archivo específico
func GitDiff(workingDir string, filePath *string) string {
	args := []string{"diff"}
	if filePath != nil && strings.TrimSpace(*filePath) != "" {
		args = append(args, strings.TrimSpace(*filePath))
	}
	return RunGitCommand(workingDir, args...)
}

// GitLog ejecuta 'git log' con un límite opcional de commits
func GitLog(workingDir string, maxCount *int) string {
	args := []string{"log", "--oneline"}
	if maxCount != nil && *maxCount > 0 {
		args = append(args, fmt.Sprintf("-n%d", *maxCount))
	} else {
		args = append(args, "-n10")
	}
	return RunGitCommand(workingDir, args...)
}

// GitAdd ejecuta 'git add <path>'
func GitAdd(workingDir string, pathspec string) string {
	if strings.TrimSpace(pathspec) == "" {
		pathspec = "."
	}
	return RunGitCommand(workingDir, "add", pathspec)
}

// GitCommit ejecuta 'git commit -m "<message>"'
func GitCommit(workingDir string, message string) string {
	if strings.TrimSpace(message) == "" {
		return "Error: el mensaje del commit no puede estar vacío."
	}
	return RunGitCommand(workingDir, "commit", "-m", message)
}
