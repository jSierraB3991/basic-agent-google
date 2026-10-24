package functions

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunTestInGo(t *testing.T) {
	// Find repo root
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get working dir: %v", err)
	}

	repoRoot := filepath.Dir(wd)
	searchRoot := repoRoot

	// Run tests pointing specifically to portscanner package
	resPortScanner := RunTestInGo("portscanner", &searchRoot)
	if !strings.Contains(resPortScanner, "¡Todas las pruebas pasaron con éxito!") {
		t.Errorf("expected portscanner tests to pass, got: %s", resPortScanner)
	}
}
