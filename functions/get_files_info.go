package functions

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func GetFilesInfo(workingDirectory string, directory *string) string {
	workingDirectoryAbs, err := filepath.Abs(workingDirectory)
	if err != nil {
		return fmt.Sprintf("%s no es un directorio valido", workingDirectory)
	}

	if directory == nil {
		directory = &workingDirectory
	}

	if *directory != workingDirectory {
		path := workingDirectory + "/" + *directory
		directory = &path
	}
	directoryAbs, err := filepath.Abs(*directory)
	if err != nil {
		return fmt.Sprintf("%s no es un directorio valido", *directory)
	}

	// Es válido si son exactamente iguales o si comienza con la ruta base + separador
	if !strings.HasPrefix(directoryAbs, workingDirectoryAbs) {
		return fmt.Sprintf("Error: %s debe iniciar con %s", directoryAbs, workingDirectoryAbs)
	}

	fileInfo, err := os.Stat(directoryAbs)
	if err != nil {
		return fmt.Sprintf("Error la dirección no existe %s", *directory)
	}

	if !fileInfo.IsDir() {
		return fmt.Sprintf("%s no es una carpeta, es un archivo", *directory)
	}

	entradas, err := os.ReadDir(directoryAbs)
	if err != nil {
		return fmt.Sprintf("Error leyendo la carpeta %s", *directory)
	}

	var finalResponse strings.Builder
	noShowDirectories := []string{".vscode", ".venv", "venv", "node_modules", ".git"}
	noShowFiles := []string{"__debug*", ".env", "build", ".gitignore"}

	for _, entrada := range entradas {
		name := entrada.Name()
		isDir := entrada.IsDir()
		if isDir {
			if slices.Contains(noShowDirectories, name) {
				continue
			}
		} else {
			if slices.Contains(noShowFiles, name) {
				continue
			}

		}
		size_info := func(info fs.FileInfo, err error) int64 {
			if err == nil {
				return info.Size()
			}
			return 0
		}(entrada.Info())
		finalResponse.WriteString(fmt.Sprintf("name: {%s}, file size: {%d}, isDir: {%t}\n", name, size_info, isDir))
	}
	return finalResponse.String()
}
