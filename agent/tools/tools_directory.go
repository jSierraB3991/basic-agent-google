package tools

import "google.golang.org/genai"

func createDirTool() *genai.FunctionDeclaration {
	return &genai.FunctionDeclaration{
		Name:        "CreateDir",
		Description: "Creates a directory (and any necessary parent directories) at the specified path.",
		Parameters: &genai.Schema{
			Type: genai.TypeObject,
			Properties: map[string]*genai.Schema{
				"dirPath": {
					Type:        genai.TypeString,
					Description: "Path of the directory to create, need the of new folder.",
				},
			},
			Required: []string{"dirPath"},
		},
	}
}

func deleteDirTool() *genai.FunctionDeclaration {
	return &genai.FunctionDeclaration{
		Name:        "DeleteDir",
		Description: "Deletes a directory and its contents.",
		Parameters: &genai.Schema{
			Type: genai.TypeObject,
			Properties: map[string]*genai.Schema{
				"dirPath": {
					Type:        genai.TypeString,
					Description: "Path of the directory to remove.",
				},
			},
			Required: []string{"dirPath"},
		},
	}
}

func getFilesInfoTool() *genai.FunctionDeclaration {
	return &genai.FunctionDeclaration{
		Name:        "GetFilesInfo",
		Description: "Lists files and subdirectories inside a working directory or optional subdirectory.",
		Parameters: &genai.Schema{
			Type: genai.TypeObject,
			Properties: map[string]*genai.Schema{
				"workingDirectory": {
					Type:        genai.TypeString,
					Description: "The absolute path of the main working directory.",
				},
				"directory": {
					Type:        genai.TypeString,
					Description: "Optional relative path of a subdirectory to list.",
					Nullable:    new(true),
				},
			},
			Required: []string{"workingDirectory"},
		},
	}
}
