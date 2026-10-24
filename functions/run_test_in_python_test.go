package functions

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunPythonTests(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "python_test_*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	testCode := `
import unittest

class SampleTest(unittest.TestCase):
    def test_example(self):
        self.assertEqual(1 + 1, 2)

if __name__ == '__main__':
    unittest.main()
`
	testFilePath := filepath.Join(tempDir, "test_sample.py")
	if err := os.WriteFile(testFilePath, []byte(testCode), 0644); err != nil {
		t.Fatalf("failed to write test file: %v", err)
	}

	result := RunPythonTests(tempDir)
	if !strings.Contains(result, "exitosas") && !strings.Contains(result, "OK") {
		t.Errorf("expected python tests to succeed, got: %s", result)
	}
}
