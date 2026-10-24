package functions

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ReadFileContent lee el contenido de un archivo.
// - workDirectory: Directorio base opcional. Si se especifica, se valida el confinamiento del archivo.
// - filename: Nombre o ruta del archivo a leer.
func ReadFileContent(workDirectory *string, filename string) string {
	if filename == "" {
		return "Error el nombre del archivo es obligatorio"
	}

	baseDir := "."
	explicitWorkDir := false
	if workDirectory != nil && *workDirectory != "" {
		baseDir = *workDirectory
		explicitWorkDir = true
	}

	baseDirAbs, err := filepath.Abs(baseDir)
	if err != nil {
		return fmt.Sprintf("Error obteniendo la ruta absoluta de '%s'", baseDir)
	}

	var filePathAbs string
	if filepath.IsAbs(filename) {
		filePathAbs = filepath.Clean(filename)
		// Si se especificó un directorio base explícito, verificar que la ruta absoluta no se salga de él
		if explicitWorkDir {
			rel, err := filepath.Rel(baseDirAbs, filePathAbs)
			if err != nil || strings.HasPrefix(rel, "..") || rel == ".." {
				return "Error: Intento de acceso denegado fuera del directorio permitido"
			}
		}
	} else {
		filePathAbs, err = filepath.Abs(filepath.Join(baseDir, filename))
		if err != nil {
			return fmt.Sprintf("Error obteniendo la ruta absoluta de '%s'", filename)
		}
		// Sanitización (CWE-22): Prevenir path traversal relativo fuera de baseDir
		rel, err := filepath.Rel(baseDirAbs, filePathAbs)
		if err != nil || strings.HasPrefix(rel, "..") || rel == ".." {
			return "Error: Intento de acceso denegado fuera del directorio permitido"
		}
	}

	// Verificar si la ruta es una carpeta o un archivo
	info, err := os.Stat(filePathAbs)
	if err != nil {
		return fmt.Sprintf("Error: No se encontró el archivo en %s", baseDir)
	}

	if info.IsDir() {
		return fmt.Sprintf("Error: En la carpeta %s solo encuentro un %s y es una carpeta", baseDir, filename)
	}

	// Leer el contenido del archivo
	contenido, err := os.ReadFile(filePathAbs)
	if err != nil {
		return fmt.Sprintf("Error al leer el archivo %s", filename)
	}

	return string(contenido)
}
