package functions

import (
	"testing"
)

func TestReadFiles(t *testing.T) {
	folderTest := "tcpgochat"
	CreateDir(folderTest)

	folderBuild := "build"
	folderTest2 := folderTest + "/" + folderBuild
	CreateDir(folderTest2)
	test := []struct {
		workDirectory string
		directory     *string
		want          string
	}{
		{workDirectory: folderTest, directory: nil, want: GetFilesInfo(".", new(folderTest))},
		{workDirectory: folderTest, directory: new("build"), want: GetFilesInfo(".", new("tcpgochat/build"))},
		{workDirectory: "/home", directory: nil, want: GetFilesInfo("/", new("home"))},
		{workDirectory: "../.", directory: new("bin"), want: "Error la dirección no existe .././bin"},
		{workDirectory: folderTest, directory: new("bin"), want: "Error la dirección no existe tcpgochat/bin"},
		{workDirectory: "../.", directory: new("main.go"), want: ".././main.go no es una carpeta, es un archivo"},
		{workDirectory: "bin", directory: nil, want: "Error la dirección no existe bin"},
		{workDirectory: "../main.go", directory: nil, want: "../main.go no es una carpeta, es un archivo"},
	}
	for i, tt := range test {
		got := GetFilesInfo(tt.workDirectory, tt.directory)
		if tt.want != got {
			valueDirectory := func(dir *string) string {
				if dir == nil {
					return "nil"
				}
				return *dir
			}(tt.directory)
			t.Errorf("error in loop %d in %s . %s", i, tt.workDirectory, valueDirectory)
		}
	}
	DeleteDir(folderTest)
}
