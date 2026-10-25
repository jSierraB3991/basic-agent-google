package tools

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestExecuteCall(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "tools_test_*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	testFilePath := filepath.Join(tempDir, "sample.txt")

	// 1. Test CreateFile
	res, err := ExecuteCall("CreateFile", map[string]any{
		"filePath": testFilePath,
		"content":  "Hello from tools test",
	})
	if err != nil {
		t.Fatalf("CreateFile tool error: %v", err)
	}
	if !strings.Contains(res["result"].(string), "Success") {
		t.Errorf("unexpected CreateFile result: %v", res)
	}

	// 2. Test ReadFileContent
	res, err = ExecuteCall("ReadFileContent", map[string]any{
		"filename": testFilePath,
	})
	if err != nil {
		t.Fatalf("ReadFileContent tool error: %v", err)
	}
	if res["result"].(string) != "Hello from tools test" {
		t.Errorf("unexpected ReadFileContent result: %v", res)
	}

	// 3. Test GetFilesInfo
	res, err = ExecuteCall("GetFilesInfo", map[string]any{
		"workingDirectory": tempDir,
	})
	if err != nil {
		t.Fatalf("GetFilesInfo tool error: %v", err)
	}
	if !strings.Contains(res["result"].(string), "sample.txt") {
		t.Errorf("expected sample.txt in GetFilesInfo, got: %v", res)
	}

	// 4. Test DeleteFile
	res, err = ExecuteCall("TestDeleteFile", map[string]any{ // wait, DeleteFile
		"filePath": testFilePath,
	})
	// Let's call DeleteFile correctly
	res, err = ExecuteCall("DeleteFile", map[string]any{
		"filePath": testFilePath,
	})
	if err != nil {
		t.Fatalf("DeleteFile tool error: %v", err)
	}
	if !strings.Contains(res["result"].(string), "success remove") {
		t.Errorf("unexpected DeleteFile result: %v", res)
	}

	// 5. Test GitStatus Tool Dispatch
	res, err = ExecuteCall("GitStatus", map[string]any{
		"workingDir": ".",
	})
	if err != nil {
		t.Fatalf("GitStatus tool error: %v", err)
	}
	if res["result"] == nil {
		t.Errorf("expected result from GitStatus")
	}

	// 6. Test Unknown Tool
	_, err = ExecuteCall("NonExistentTool", map[string]any{})
	if err == nil {
		t.Errorf("expected error for non-existent tool call")
	}
}
