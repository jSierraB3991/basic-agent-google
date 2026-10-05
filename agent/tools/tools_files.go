package tools

import "google.golang.org/genai"

func createFileTool() *genai.FunctionDeclaration {
	return &genai.FunctionDeclaration{
		Name:        "CreateFile",
		Description: "Creates a new file with the specified content at the given path.",
		Parameters: &genai.Schema{
			Type: genai.TypeObject,
			Properties: map[string]*genai.Schema{
				"filePath": {
					Type:        genai.TypeString,
					Description: "Relative or absolute path of the file to create.",
				},
				"content": {
					Type:        genai.TypeString,
					Description: "The exact textual content to write into the file.",
				},
			},
			Required: []string{"filePath", "content"},
		},
	}
}

func deleteFileTool() *genai.FunctionDeclaration {
	return &genai.FunctionDeclaration{
		Name:        "DeleteFile",
		Description: "Deletes a specified file from the filesystem.",
		Parameters: &genai.Schema{
			Type: genai.TypeObject,
			Properties: map[string]*genai.Schema{
				"filePath": {
					Type:        genai.TypeString,
					Description: "Path of the file to delete.",
				},
			},
			Required: []string{"filePath"},
		},
	}
}

func readFileContentTool() *genai.FunctionDeclaration {
	return &genai.FunctionDeclaration{
		Name:        "ReadFileContent",
		Description: "Reads and returns the complete content of a target file.",
		Parameters: &genai.Schema{
			Type: genai.TypeObject,
			Properties: map[string]*genai.Schema{
				"workDirectory": {
					Type:        genai.TypeString,
					Description: "Optional working directory root for relative paths.",
					Nullable:    new(true),
				},
				"filename": {
					Type:        genai.TypeString,
					Description: "The name or path of the file to read.",
				},
			},
			Required: []string{"filename"},
		},
	}
}
