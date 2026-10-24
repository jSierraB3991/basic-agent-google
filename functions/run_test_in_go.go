package functions

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// RunTestInGo ejecuta pruebas de Go en el directorio o paquete especificado.
// target: puede ser un archivo (ej: main.go, scanner_test.go), una carpeta, o "./..."
// searchRoot: directorio raíz opcional desde donde buscar o ejecutar.
func RunTestInGo(target string, searchRoot *string) string {
	rootDirectory := "."
	if searchRoot != nil && strings.TrimSpace(*searchRoot) != "" {
		rootDirectory = strings.TrimSpace(*searchRoot)
	}

	target = strings.TrimSpace(target)
	if target == "" {
		target = "./..."
	}

	runDir := rootDirectory
	testPkg := "./..."

	// Si target termina en /... o es ., es un paquete o ruta directa
	if target == "./..." || target == "." {
		runDir = rootDirectory
		testPkg = target
	} else if strings.HasSuffix(target, "/...") {
		runDir = rootDirectory
		testPkg = target
	} else {
		// Verificar si target es un directorio o archivo existente relativo a rootDirectory
		fullPath := filepath.Join(rootDirectory, target)
		fi, err := os.Stat(fullPath)
		if err == nil {
			if fi.IsDir() {
				runDir = fullPath
				testPkg = "./..."
			} else {
				runDir = filepath.Dir(fullPath)
				testPkg = "./..."
			}
		} else {
			// Intentar buscar el archivo recursivamente si fue dado sólo por nombre
			foundDir, err := getDirByFile(rootDirectory, target)
			if err == nil {
				runDir = foundDir
				testPkg = "./..."
			} else {
				// Tratarlo como paquete o patrón de go test
				runDir = rootDirectory
				testPkg = target
			}
		}
	}

	output, err := RunGoTests(runDir, testPkg)
	if err != nil {
		if strings.TrimSpace(output) != "" {
			return fmt.Sprintf("Pruebas fallidas:\n%s", output)
		}
		return fmt.Sprintf("Pruebas fallidas: %v", err)
	}

	if strings.TrimSpace(output) == "" {
		return "¡Todas las pruebas pasaron con éxito! (sin salida detallada)"
	}
	return fmt.Sprintf("¡Todas las pruebas pasaron con éxito!\n%s", output)
}

// getDirByFile busca un archivo específico a partir de un directorio inicial
// y retorna la ruta absoluta de la carpeta que lo contiene.
func getDirByFile(startDir string, targetFileName string) (string, error) {
	var targetDir string

	err := filepath.Walk(startDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && (info.Name() == ".git" || info.Name() == "vendor" || info.Name() == "node_modules") {
			return filepath.SkipDir
		}

		if !info.IsDir() && info.Name() == targetFileName {
			absPath, err := filepath.Abs(filepath.Dir(path))
			if err != nil {
				return err
			}
			targetDir = absPath
			return filepath.SkipAll
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

// RunGoTests ejecuta 'go test -count=1 -v <pkg>' en el directorio especificado y captura la salida.
func RunGoTests(dir string, args ...string) (string, error) {
	cmdArgs := []string{"test", "-count=1", "-v"}
	if len(args) > 0 {
		cmdArgs = append(cmdArgs, args...)
	} else {
		cmdArgs = append(cmdArgs, "./...")
	}

	cmd := exec.Command("go", cmdArgs...)
	cmd.Dir = dir

	outputBytes, err := cmd.CombinedOutput()
	return string(outputBytes), err
}
