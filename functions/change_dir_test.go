package functions

import (
	"os"
	"path/filepath"
	"testing"
)

func TestChangeDir(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "test-chdir-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	origDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get current wd: %v", err)
	}
	defer os.Chdir(origDir)

	res, err := ChangeDir(tempDir)
	if err != nil {
		t.Fatalf("expected no error changing dir, got %v", err)
	}

	evalTemp, _ := filepath.EvalSymlinks(tempDir)
	evalRes, _ := filepath.EvalSymlinks(res)

	if evalRes != evalTemp {
		t.Errorf("expected current working directory to be %s, got %s", evalTemp, evalRes)
	}
}

func TestChangeDir_Invalid(t *testing.T) {
	_, err := ChangeDir("")
	if err == nil {
		t.Error("expected error for empty path, got nil")
	}

	_, err = ChangeDir("non-existent-directory-xyz-12345")
	if err == nil {
		t.Error("expected error for non-existent directory, got nil")
	}
}
