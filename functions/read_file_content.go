package functions

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ReadFileContent lee el contenido de un archivo.
// - workDirectory: Si es nil o "", se usa el directorio actual (".").
// - filename: Si es nil o "", busca la ruta en workDirectory.
func ReadFileContent(workDirectory *string, filename string) string {
	// 1. Determinar el directorio base
	baseDir := "."
	if workDirectory != nil && *workDirectory != "" {
		baseDir = *workDirectory
	}

	// 2. Construir la ruta final del archivo
	filePath := baseDir
	if filename == "" {
		return "Error el nombre del archivo es obligatorio"
	}
	filePatAbs, err := filepath.Abs(filepath.Join(baseDir, filename))
	if filename == "" {
		return "Error obteniendo la ruta abosluta filepath.Abs(baseDir + '/' + filename)"
	}

	baseDirAbs, err := filepath.Abs(baseDir)
	if filename == "" {
		return "Error obteniendo la ruta abosluta filepath.Abs(baseDir)"
	}

	// 3. Sanitización (CWE-22): Verificar que filePatAbs esté dentro de baseAbs
	rel, err := filepath.Rel(baseDirAbs, filePatAbs)
	if err != nil || strings.HasPrefix(rel, "..") || rel == ".." {
		return "Error: Intento de acceso denegado fuera del directorio permitido"
	}

	// 4. Verificar si la ruta es una carpeta o un archivo
	info, err := os.Stat(filePatAbs)
	if err != nil {
		return fmt.Sprintf("Error: No se encontró el archivo en %s", baseDir)
	}

	if info.IsDir() {
		return fmt.Sprintf("Error: En la carpeta %s solo encuentro un %s y es una carpeta", baseDir, filename)
	}

	// 5. Leer el contenido del archivo
	// #nosec G304 -- Sanitización validada mediante filepath.Rel arriba
	contenido, err := os.ReadFile(filePatAbs)
	if err != nil {
		return fmt.Sprintf("Error al leer el archivo %s", filePath)
	}

	return string(contenido)
}
