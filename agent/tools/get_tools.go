package tools

import "google.golang.org/genai"

func getTools() *genai.Tool {
	return &genai.Tool{
		FunctionDeclarations: []*genai.FunctionDeclaration{
			createFileTool(),
			getFilesInfoTool(),
			deleteFileTool(),
			createDirTool(),
			deleteDirTool(),
			readFileContentTool(),
			runTestInGoTool(),
			runTestInPythonTool(),
		},
	}
}

func GetConfigGenerator(systemPrompt string) *genai.GenerateContentConfig {
	tools := getTools()
	return &genai.GenerateContentConfig{
		Tools: []*genai.Tool{tools},
		SystemInstruction: &genai.Content{
			Parts: []*genai.Part{
				genai.NewPartFromText(systemPrompt),
			},
		},
	}
}
