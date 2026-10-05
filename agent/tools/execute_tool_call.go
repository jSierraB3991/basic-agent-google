package tools

import (
	"fmt"

	"github.com/jsierrab3991/basic-agent-google/functions"
)

func ExecuteCall(name string, args map[string]any) (map[string]any, error) {
	switch name {

	case "CreateFile":
		filePath, ok := args["filePath"].(string)
		if !ok {
			return nil, fmt.Errorf("missing or invalid 'filePath' parameter")
		}

		content, ok := args["content"].(string)
		if !ok {
			return nil, fmt.Errorf("missing or invalid 'content' parameter")
		}

		result := functions.CreateFile(filePath, content)
		return map[string]any{"result": result}, nil

	case "GetFilesInfo":
		workingDir, ok := args["workingDirectory"].(string)
		if !ok {
			return nil, fmt.Errorf("missing or invalid 'workingDirectory' parameter")
		}

		var directory *string
		if val, exists := args["directory"]; exists && val != nil {
			if dirStr, ok := val.(string); ok {
				directory = &dirStr
			}
		}

		result := functions.GetFilesInfo(workingDir, directory)
		return map[string]any{"result": result}, nil

	case "DeleteFile":
		filePath, ok := args["filePath"].(string)
		if !ok {
			return nil, fmt.Errorf("missing or invalid 'filePath' parameter")
		}

		result := functions.DeleteFile(filePath)
		return map[string]any{"result": result}, nil

	case "CreateDir":
		dirPath, ok := args["dirPath"].(string)
		if !ok {
			return nil, fmt.Errorf("missing or invalid 'dirPath' parameter")
		}

		err := functions.CreateDir(dirPath)
		if err != nil {
			return nil, err
		}
		return map[string]any{"result": fmt.Sprintf("Directory '%s' created successfully", dirPath)}, nil

	case "DeleteDir":
		dirPath, ok := args["dirPath"].(string)
		if !ok {
			return nil, fmt.Errorf("missing or invalid 'dirPath' parameter")
		}

		err := functions.DeleteDir(dirPath)
		if err != nil {
			return nil, err
		}
		return map[string]any{"result": fmt.Sprintf("Directory '%s' deleted successfully", dirPath)}, nil

	case "ReadFileContent":
		filename, ok := args["filename"].(string)
		if !ok {
			return nil, fmt.Errorf("missing or invalid 'filename' parameter")
		}

		var workDirectory *string
		if val, exists := args["workDirectory"]; exists && val != nil {
			if dirStr, ok := val.(string); ok {
				workDirectory = &dirStr
			}
		}

		result := functions.ReadFileContent(workDirectory, filename)
		return map[string]any{"result": result}, nil

	case "RunTestInGo":
		targetFile, ok := args["targetFile"].(string)
		if !ok {
			return nil, fmt.Errorf("missing or invalid 'targetFile' parameter")
		}

		var searchRoot *string
		if val, exists := args["searchRoot"]; exists && val != nil {
			if rootStr, ok := val.(string); ok {
				searchRoot = &rootStr
			}
		}

		result := functions.RunTestInGo(targetFile, searchRoot)
		return map[string]any{"result": result}, nil

	case "RunPythonTests":
		dirPath, ok := args["dirPath"].(string)
		if !ok {
			return nil, fmt.Errorf("missing or invalid 'dirPath' parameter")
		}

		result := functions.RunPythonTests(dirPath)
		return map[string]any{"result": result}, nil

	default:
		return nil, fmt.Errorf("unknown function call: %s", name)
	}
}
