package functions

import (
	"fmt"
	"testing"
)

func TestGetContentFile(t *testing.T) {
	fileTest := "test.txt"
	fileContentTest := "Hola mundo"
	dirTest := "folder_test"

	CreateFile(fileTest, fileContentTest)
	CreateDir(dirTest)
	test := []struct {
		workDirectory *string
		filename      string
		want          string
	}{
		{workDirectory: new("."), filename: fileTest, want: fileContentTest},
		{workDirectory: nil, filename: fileTest, want: fileContentTest},
		{workDirectory: new("."), filename: "", want: "Error el nombre del archivo es obligatorio"},
		{workDirectory: nil, filename: "", want: "Error el nombre del archivo es obligatorio"},
		{workDirectory: new(dirTest), filename: "test.txt", want: fmt.Sprintf("Error: No se encontró el archivo en %s", dirTest)},
		{workDirectory: new("."), filename: dirTest, want: fmt.Sprintf("Error: En la carpeta . solo encuentro un %s y es una carpeta", dirTest)},
		{workDirectory: nil, filename: dirTest, want: fmt.Sprintf("Error: En la carpeta . solo encuentro un %s y es una carpeta", dirTest)},
	}
	for i, tt := range test {
		got := ReadFileContent(tt.workDirectory, tt.filename)
		if tt.want != got {
			workDirectory := func(dir *string) string {
				if dir == nil {
					return "nil"
				}
				return *dir
			}(tt.workDirectory)
			t.Errorf("error in loop %d in %s . %s", i, workDirectory, tt.filename)
		}
	}
	DeleteFile(fileTest)
	DeleteDir(dirTest)
}
