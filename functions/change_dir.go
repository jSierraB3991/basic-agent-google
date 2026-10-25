package functions

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ChangeDir cambia el directorio de trabajo actual del proceso a la ruta especificada.
func ChangeDir(dirPath string) (string, error) {
	dirPath = strings.TrimSpace(dirPath)
	if dirPath == "" {
		return "", fmt.Errorf("changeDir: la ruta del directorio no puede estar vacía")
	}

	cleanPath := filepath.Clean(dirPath)

	if err := os.Chdir(cleanPath); err != nil {
		return "", fmt.Errorf("changeDir: error al cambiar al directorio '%s': %w", cleanPath, err)
	}

	currentDir, err := os.Getwd()
	if err != nil {
		return cleanPath, nil
	}

	return currentDir, nil
}
