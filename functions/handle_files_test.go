package functions

import (
	"testing"
)

// CreateFile crea un archivo en la ruta especificada con el contenido dado.
// Asegura que los directorios padre existan antes de escribir.
func TestCreateFile(t *testing.T) {
	testCases := []struct {
		name          string
		fileDirectory string
		fileContent   string
		deleteFile    bool
		want          string
	}{
		{name: "ALL OK 1", fileDirectory: "test.txt", fileContent: "Hello world", want: "Success file created test.txt", deleteFile: true},
		{name: "Error path vacío", fileDirectory: "", fileContent: "Hello world", want: "createFile: la ruta del archivo no puede estar vacía", deleteFile: false},
		{name: "Error path vacío con espacios", fileDirectory: "    ", fileContent: "Hello world", want: "createFile: la ruta del archivo no puede estar vacía", deleteFile: false},
		{name: "Error content vacío", fileDirectory: "test.txt", fileContent: "", want: "createFile: el contenido del archivo no puede estar vácio", deleteFile: false},
		{name: "Error content vacío con espacios", fileDirectory: "test.txt", fileContent: "     ", want: "createFile: el contenido del archivo no puede estar vácio", deleteFile: false},
		{name: "Error content vacío", fileDirectory: "/root/route/fall/test.txt", fileContent: "Hello world", want: "createFile: no se pudo crear el directorio '/root/route/fall'", deleteFile: false},
	}

	for _, tc := range testCases {
		got := CreateFile(tc.fileDirectory, tc.fileContent)
		if got != tc.want {
			t.Errorf("error in loop %s in %s . %s", tc.name, tc.fileDirectory, tc.fileContent)
		}
		if tc.deleteFile {
			DeleteFile("./test.txt")
		}
	}
}

// DeleteFile elimina un archivo en la ruta especificada.
// Retorna nil si el archivo ya no existe.
func TestDeleteFile(t *testing.T) {
	testCases := []struct {
		name          string
		fileDirectory string
		createFile    bool
		want          string
	}{
		{name: "ALL OK", fileDirectory: "test.txt", createFile: true, want: "success remove file test.txt"},
		{name: "Delete by path void", fileDirectory: "", createFile: false, want: "deleteFile: la ruta del archivo no puede estar vacía"},
		{name: "Delete by path with spaces", fileDirectory: "   ", createFile: false, want: "deleteFile: la ruta del archivo no puede estar vacía"},
		{name: "Delete by path wrong", fileDirectory: "/root/lolcat/test/main.py", createFile: false, want: "deleteFile: error al eliminar '/root/lolcat/test/main.py'"},
	}

	for _, tc := range testCases {
		if tc.createFile {
			CreateFile(tc.fileDirectory, "Lorem ipsum")
		}

		got := DeleteFile(tc.fileDirectory)
		if got != tc.want {
			t.Errorf("error in loop %s in %s", tc.name, tc.fileDirectory)
		}
	}
}
