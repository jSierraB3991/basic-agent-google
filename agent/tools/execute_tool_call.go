package tools

import (
	"fmt"
	"strings"

	"github.com/jsierrab3991/basic-agent-google/functions"
)

func ExecuteCall(name string, args map[string]any) (map[string]any, error) {
	switch name {

	case "CreateFile":
		filePath, ok := args["filePath"].(string)
		if !ok || strings.TrimSpace(filePath) == "" {
			return nil, fmt.Errorf("missing or invalid 'filePath' parameter")
		}

		content, ok := args["content"].(string)
		if !ok {
			return nil, fmt.Errorf("missing or invalid 'content' parameter")
		}

		result := functions.CreateFile(filePath, content)
		return map[string]any{"result": result}, nil

	case "GetFilesInfo":
		workingDir := "."
		if val, exists := args["workingDirectory"]; exists && val != nil {
			if wdStr, ok := val.(string); ok && strings.TrimSpace(wdStr) != "" {
				workingDir = strings.TrimSpace(wdStr)
			}
		}

		var directory *string
		if val, exists := args["directory"]; exists && val != nil {
			if dirStr, ok := val.(string); ok && strings.TrimSpace(dirStr) != "" {
				trimmed := strings.TrimSpace(dirStr)
				directory = &trimmed
			}
		}

		result := functions.GetFilesInfo(workingDir, directory)
		return map[string]any{"result": result}, nil

	case "DeleteFile":
		filePath, ok := args["filePath"].(string)
		if !ok || strings.TrimSpace(filePath) == "" {
			return nil, fmt.Errorf("missing or invalid 'filePath' parameter")
		}

		result := functions.DeleteFile(filePath)
		return map[string]any{"result": result}, nil

	case "CreateDir":
		dirPath, ok := args["dirPath"].(string)
		if !ok || strings.TrimSpace(dirPath) == "" {
			return nil, fmt.Errorf("missing or invalid 'dirPath' parameter")
		}

		err := functions.CreateDir(dirPath)
		if err != nil {
			return nil, err
		}
		return map[string]any{"result": fmt.Sprintf("Directory '%s' created successfully", dirPath)}, nil

	case "DeleteDir":
		dirPath, ok := args["dirPath"].(string)
		if !ok || strings.TrimSpace(dirPath) == "" {
			return nil, fmt.Errorf("missing or invalid 'dirPath' parameter")
		}

		err := functions.DeleteDir(dirPath)
		if err != nil {
			return nil, err
		}
		return map[string]any{"result": fmt.Sprintf("Directory '%s' deleted successfully", dirPath)}, nil

	case "ChangeDir":
		dirPath, ok := args["dirPath"].(string)
		if !ok || strings.TrimSpace(dirPath) == "" {
			return nil, fmt.Errorf("missing or invalid 'dirPath' parameter")
		}

		newDir, err := functions.ChangeDir(dirPath)
		if err != nil {
			return nil, err
		}
		return map[string]any{"result": fmt.Sprintf("Changed working directory to '%s'", newDir)}, nil

	case "ReadFileContent":
		filename, ok := args["filename"].(string)
		if !ok || strings.TrimSpace(filename) == "" {
			return nil, fmt.Errorf("missing or invalid 'filename' parameter")
		}

		var workDirectory *string
		if val, exists := args["workDirectory"]; exists && val != nil {
			if dirStr, ok := val.(string); ok && strings.TrimSpace(dirStr) != "" {
				trimmed := strings.TrimSpace(dirStr)
				workDirectory = &trimmed
			}
		}

		result := functions.ReadFileContent(workDirectory, filename)
		return map[string]any{"result": result}, nil

	case "RunTestInGo":
		targetFile := "./..."
		if val, exists := args["targetFile"]; exists && val != nil {
			if tf, ok := val.(string); ok && strings.TrimSpace(tf) != "" {
				targetFile = strings.TrimSpace(tf)
			}
		}

		var searchRoot *string
		if val, exists := args["searchRoot"]; exists && val != nil {
			if rootStr, ok := val.(string); ok && strings.TrimSpace(rootStr) != "" {
				trimmed := strings.TrimSpace(rootStr)
				searchRoot = &trimmed
			}
		}

		result := functions.RunTestInGo(targetFile, searchRoot)
		return map[string]any{"result": result}, nil

	case "RunPythonTests":
		dirPath, ok := args["dirPath"].(string)
		if !ok || strings.TrimSpace(dirPath) == "" {
			return nil, fmt.Errorf("missing or invalid 'dirPath' parameter")
		}

		result := functions.RunPythonTests(dirPath)
		return map[string]any{"result": result}, nil

	default:
		return nil, fmt.Errorf("unknown function call: %s", name)
	}
}
