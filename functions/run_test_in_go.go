package functions

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func RunTestInGo(targetFile string, searchRoot *string) string {
	if targetFile != "main.go" {
		return "run_test_in_go: solo se permite main.go para correr pruebas en golang"
	}
	rootDirectory := "."
	if searchRoot != nil && strings.TrimSpace(*searchRoot) != "" {
		rootDirectory = *searchRoot
	}

	// 1. Obtener la ruta del directorio que contiene el archivo
	projectDir, err := getDirByFile(rootDirectory, targetFile)
	if err != nil {
		return fmt.Sprintf("run_test_in_go: error obteniendo el directorio donde se encuentra %s", targetFile)
	}

	fmt.Printf("Directorio encontrado para '%s': %s\n", targetFile, projectDir)
	fmt.Println("Ejecutando pruebas...")

	// 2. Ejecutar go test en ese directorio
	if err := RunGoTests(projectDir); err != nil {
		return fmt.Sprintf("Pruebas fallidas: %v\n", err)
	}

	return "¡Todas las pruebas pasaron con éxito!"
}

// GetDirByFile busca un archivo específico a partir de un directorio inicial
// y retorna la ruta absoluta de la carpeta que lo contiene.
func getDirByFile(startDir string, targetFileName string) (string, error) {
	var targetDir string

	err := filepath.Walk(startDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Evita buscar dentro de carpetas ocultas o de dependencias comunes (.git, vendor, node_modules)
		if info.IsDir() && (info.Name() == ".git" || info.Name() == "vendor") {
			return filepath.SkipDir
		}

		// Compara si el archivo coincide con el parámetro solicitado
		if !info.IsDir() && info.Name() == targetFileName {
			absPath, err := filepath.Abs(filepath.Dir(path))
			if err != nil {
				return err
			}
			targetDir = absPath
			return filepath.SkipAll // Detener la búsqueda al encontrar la primera coincidencia
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	if targetDir == "" {
		return "", fmt.Errorf("no se encontró la carpeta con el archivo '%s' en '%s'", targetFileName, startDir)
	}

	return targetDir, nil
}

// RunGoTests ejecuta 'go test -count=1 -v ./...' en el directorio especificado.
func RunGoTests(dir string) error {
	cmd := exec.Command("go", "test", "-count=1", "-v", "./...")
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
