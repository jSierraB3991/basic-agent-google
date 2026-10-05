package functions

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

// RunPythonTests ejecuta las pruebas de Python en el directorio especificado y retorna el resultado como string.
// Primero intenta ejecutar con 'pytest'; si no está instalado, intenta con 'python3 -m unittest'.
func RunPythonTests(dirPath string) string {
	dirPath = strings.TrimSpace(dirPath)
	if dirPath == "" {
		return "RunPythonTests error: la ruta del directorio no puede estar vacía"
	}

	cleanPath := filepath.Clean(dirPath)

	// Intentar ejecutar con pytest
	output, err := runCommand(cleanPath, "pytest", "-v")
	if err == nil {
		return output
	}

	// Si pytest falla por no estar instalado/disponible, intentar con el módulo estándar unittest
	fallbackOutput, fallbackErr := runCommand(cleanPath, "python3", "-m", "unittest", "discover", "-v")
	if fallbackErr == nil {
		return fallbackOutput
	}

	// Si ambos fallan o las pruebas no pasaron, devolver el detalle del resultado combinando salidas
	return fmt.Sprintf("Resultado de pruebas en '%s':\n\n--- pytest ---\n%s\n--- unittest ---\n%s", cleanPath, output, fallbackOutput)
}

// runCommand es un helper privado que ejecuta comandos dentro del directorio objetivo
func runCommand(dir string, name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	cmd.Dir = dir

	// CombinedOutput captura tanto stdout como stderr
	outputBytes, err := cmd.CombinedOutput()
	outputStr := string(outputBytes)

	if err != nil {
		return fmt.Sprintf("Error ejecutando '%s': %v\nSalida:\n%s", name, err, outputStr), err
	}

	return outputStr, nil
}
