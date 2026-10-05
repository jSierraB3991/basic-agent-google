package functions

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// CreateDir crea un directorio y todos los directorios padres necesarios.
// Los permisos 0755 permiten lectura/ejecución para todos y escritura para el propietario.
func CreateDir(dirPath string) error {
	dirPath = strings.TrimSpace(dirPath)
	if dirPath == "" {
		return fmt.Errorf("createDir: la ruta del directorio no puede estar vacía")
	}

	// Normaliza la ruta para evitar inconsistencias según el sistema operativo
	cleanPath := filepath.Clean(dirPath)

	if err := os.MkdirAll(cleanPath, 0755); err != nil {
		return fmt.Errorf("createDir: error al crear el directorio '%s'", cleanPath)
	}

	return nil
}

// DeleteDir elimina un directorio y todo su contenido de forma recursiva.
// Si la carpeta no existe, retorna nil (idempotente).
func DeleteDir(dirPath string) error {
	dirPath = strings.TrimSpace(dirPath)
	if dirPath == "" {
		return fmt.Errorf("deleteDir: la ruta del directorio no puede estar vacía")
	}

	cleanPath := filepath.Clean(dirPath)

	// Previene eliminar accidentalmente la raíz o el directorio actual
	if cleanPath == "." || cleanPath == "/" || cleanPath == filepath.VolumeName(cleanPath)+"\\" {
		return fmt.Errorf("deleteDir: no se permite eliminar la ruta raíz o el directorio actual ('%s')", cleanPath)
	}

	if err := os.RemoveAll(cleanPath); err != nil {
		return fmt.Errorf("deleteDir: error al eliminar el directorio '%s'", cleanPath)
	}

	return nil
}
