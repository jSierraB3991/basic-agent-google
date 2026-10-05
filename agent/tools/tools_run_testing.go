package tools

import "google.golang.org/genai"

func runTestInGoTool() *genai.FunctionDeclaration {
	return &genai.FunctionDeclaration{
		Name:        "RunTestInGo",
		Description: "Executes Go test suites for a target test file or directory.",
		Parameters: &genai.Schema{
			Type: genai.TypeObject,
			Properties: map[string]*genai.Schema{
				"targetFile": {
					Type:        genai.TypeString,
					Description: "The test file or package target path to run Go tests on.",
				},
				"searchRoot": {
					Type:        genai.TypeString,
					Description: "Optional root path from which to run search/testing.",
					Nullable:    new(true),
				},
			},
			Required: []string{"targetFile"},
		},
	}
}

func runTestInPythonTool() *genai.FunctionDeclaration {
	return &genai.FunctionDeclaration{
		Name:        "RunPythonTests",
		Description: "Executes Python unit tests inside the target directory.",
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
