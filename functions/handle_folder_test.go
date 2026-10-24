package functions

import (
	"os"
	"path/filepath"
	"testing"
)

func TestHandleFolder(t *testing.T) {
	tempBase, err := os.MkdirTemp("", "folder_test_*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempBase)

	targetDir := filepath.Join(tempBase, "nested", "deep", "dir")

	// 1. CreateDir
	if err := CreateDir(targetDir); err != nil {
		t.Errorf("CreateDir failed: %v", err)
	}

	info, err := os.Stat(targetDir)
	if err != nil || !info.IsDir() {
		t.Errorf("expected directory to exist")
	}

	// 2. Empty dirPath
	if err := CreateDir("   "); err == nil {
		t.Errorf("expected error for empty dirPath")
	}

	// 3. DeleteDir
	if err := DeleteDir(targetDir); err != nil {
		t.Errorf("DeleteDir failed: %v", err)
	}

	if _, err := os.Stat(targetDir); !os.IsNotExist(err) {
		t.Errorf("expected directory to be deleted")
	}

	// 4. Delete root prevention
	if err := DeleteDir("."); err == nil {
		t.Errorf("expected error when trying to delete '.'")
	}
}
