package functions

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// CreateFile crea un archivo en la ruta especificada con el contenido dado.
// Asegura que los directorios padre existan antes de escribir.
func CreateFile(filePath string, content string) string {
	filePath = strings.TrimSpace(filePath)
	if filePath == "" {
		return "createFile: la ruta del archivo no puede estar vacía"
	}

	if strings.TrimSpace(content) == "" {
		return "createFile: el contenido del archivo no puede estar vácio"
	}

	// Crear los directorios padre si no existen (permisos 0755 por defecto)
	dir := filepath.Dir(filePath)
	if dir != "." {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Sprintf("createFile: no se pudo crear el directorio '%s'", dir)
		}
	}

	filePathAbs, err := filepath.Abs(filePath)
	if err != nil {
		return fmt.Sprintf("createFile: no se pudo obtener la ruta absoluta de '%s'", filePath)
	}

	// Escribir el archivo con permisos de lectura y escritura solo para el propietario (0600)
	if err := os.WriteFile(filePathAbs, []byte(content), 0600); err != nil {
		return fmt.Sprintf("createFile: error al escribir en '%s'", filePathAbs)
	}

	return fmt.Sprintf("Success file created %s", filePath)
}

// DeleteFile elimina un archivo en la ruta especificada.
// Retorna nil si el archivo ya no existe.
func DeleteFile(filePath string) string {
	filePath = strings.TrimSpace(filePath)
	if filePath == "" {
		return "deleteFile: la ruta del archivo no puede estar vacía"
	}

	filePathAbs, err := filepath.Abs(filePath)
	if err != nil {
		return fmt.Sprintf("deleteFile: error al obtener la ruta absoluta del archivo fielPath: %s", filePath)
	}

	if err := os.Remove(filePathAbs); err != nil {
		// Ignorar el error si el archivo ya no existía
		if os.IsNotExist(err) {
			return fmt.Sprintf("success remove file %s", filePath)
		}
		return fmt.Sprintf("deleteFile: error al eliminar '%s'", filePath)
	}

	return fmt.Sprintf("success remove file %s", filePath)
}
