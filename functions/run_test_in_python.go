package functions

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

// RunPythonTests ejecuta las pruebas de Python en el directorio especificado y retorna el resultado como string.
// Si 'pytest' está disponible en el sistema, lo utiliza. De lo contrario, recurre a 'python3 -m unittest'.
func RunPythonTests(dirPath string) string {
	dirPath = strings.TrimSpace(dirPath)
	if dirPath == "" {
		return "RunPythonTests error: la ruta del directorio no puede estar vacía"
	}

	cleanPath := filepath.Clean(dirPath)

	// Si pytest está instalado en el PATH, ejecutarlo directamente
	if _, err := exec.LookPath("pytest"); err == nil {
		output, err := runCommand(cleanPath, "pytest", "-v")
		if err != nil {
			return fmt.Sprintf("Pruebas con pytest fallidas o con errores:\n%s", output)
		}
		return fmt.Sprintf("Pruebas con pytest exitosas:\n%s", output)
	}

	// De lo contrario, usar el módulo estándar unittest de python3
	if _, err := exec.LookPath("python3"); err == nil {
		output, err := runCommand(cleanPath, "python3", "-m", "unittest", "discover", "-v")
		if err != nil {
			return fmt.Sprintf("Pruebas con unittest fallidas o con errores:\n%s", output)
		}
		return fmt.Sprintf("Pruebas con unittest exitosas:\n%s", output)
	}

	return "RunPythonTests error: no se encontró ni 'pytest' ni 'python3' en el sistema"
}

// runCommand es un helper privado que ejecuta comandos dentro del directorio objetivo
func runCommand(dir string, name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	cmd.Dir = dir

	// CombinedOutput captura tanto stdout como stderr
	outputBytes, err := cmd.CombinedOutput()
	outputStr := string(outputBytes)

	if err != nil {
		if strings.TrimSpace(outputStr) != "" {
			return outputStr, err
		}
		return fmt.Sprintf("Error ejecutando '%s': %v", name, err), err
	}

	return outputStr, nil
}
