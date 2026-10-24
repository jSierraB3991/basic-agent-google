package tools

import "google.golang.org/genai"

func runTestInGoTool() *genai.FunctionDeclaration {
	return &genai.FunctionDeclaration{
		Name:        "RunTestInGo",
		Description: "Executes Go test suites for a target test file, directory, or package path.",
		Parameters: &genai.Schema{
			Type: genai.TypeObject,
			Properties: map[string]*genai.Schema{
				"targetFile": {
					Type:        genai.TypeString,
					Description: "The test file, directory, or package target path to run Go tests on (e.g. './...', '.', 'portscanner', or a specific test file). Defaults to './...'.",
					Nullable:    new(true),
				},
				"searchRoot": {
					Type:        genai.TypeString,
					Description: "Optional root path from which to run search/testing.",
					Nullable:    new(true),
				},
			},
		},
	}
}

func runTestInPythonTool() *genai.FunctionDeclaration {
	return &genai.FunctionDeclaration{
		Name:        "RunPythonTests",
		Description: "Executes Python unit tests inside the target directory using pytest or unittest.",
		Parameters: &genai.Schema{
			Type: genai.TypeObject,
			Properties: map[string]*genai.Schema{
				"dirPath": {
					Type:        genai.TypeString,
					Description: "The directory containing Python tests.",
				},
			},
			Required: []string{"dirPath"},
		},
	}
}
