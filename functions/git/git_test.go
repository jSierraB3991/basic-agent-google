package git

import (
	"strings"
	"testing"
)

func TestGitStatus(t *testing.T) {
	// Solo probamos que la función ejecute git status en el directorio actual (puede o no ser repo git, pero no debe crashear)
	res := GitStatus(".")
	if res == "" {
		t.Errorf("Esperaba alguna respuesta de git status")
	}
}

func TestGitDiff(t *testing.T) {
	res := GitDiff(".", nil)
	if strings.Contains(res, "Error ejecutando 'git diff'") {
		// Es posible que no sea un repo git en algunos entornos de prueba o sí
		t.Logf("Resultado de git diff: %s", res)
	}
}

func TestGitLog(t *testing.T) {
	count := 5
	res := GitLog(".", &count)
	if res == "" {
		t.Errorf("Esperaba alguna respuesta de git log")
	}
}

func TestGitCommitEmpty(t *testing.T) {
	res := GitCommit(".", "")
	if !strings.Contains(res, "mensaje del commit no puede estar vacío") {
		t.Errorf("Esperaba error por commit vacío, obtuvo: %s", res)
	}
}
