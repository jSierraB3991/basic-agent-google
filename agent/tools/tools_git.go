package tools

import (
	"google.golang.org/genai"
)

func gitStatusTool() *genai.FunctionDeclaration {
	return &genai.FunctionDeclaration{
		Name:        "GitStatus",
		Description: "Muestra el estado del repositorio Git actual (archivos modificados, sin seguimiento, etc.).",
		Parameters: &genai.Schema{
			Type: genai.TypeObject,
			Properties: map[string]*genai.Schema{
				"workingDir": {
					Type:        genai.TypeString,
					Description: "Directorio de trabajo del repositorio Git (opcional, por defecto '.').",
				},
			},
		},
	}
}

func gitDiffTool() *genai.FunctionDeclaration {
	return &genai.FunctionDeclaration{
		Name:        "GitDiff",
		Description: "Muestra las diferencias (diff) en el repositorio Git o en un archivo específico.",
		Parameters: &genai.Schema{
			Type: genai.TypeObject,
			Properties: map[string]*genai.Schema{
				"workingDir": {
					Type:        genai.TypeString,
					Description: "Directorio de trabajo del repositorio Git (opcional, por defecto '.').",
				},
				"filePath": {
					Type:        genai.TypeString,
					Description: "Ruta del archivo específico para ver su diff (opcional).",
				},
			},
		},
	}
}

func gitLogTool() *genai.FunctionDeclaration {
	return &genai.FunctionDeclaration{
		Name:        "GitLog",
		Description: "Muestra el historial reciente de commits del repositorio Git.",
		Parameters: &genai.Schema{
			Type: genai.TypeObject,
			Properties: map[string]*genai.Schema{
				"workingDir": {
					Type:        genai.TypeString,
					Description: "Directorio de trabajo del repositorio Git (opcional, por defecto '.').",
				},
				"maxCount": {
					Type:        genai.TypeInteger,
					Description: "Número máximo de commits a mostrar (opcional, por defecto 10).",
				},
			},
		},
	}
}

func gitAddTool() *genai.FunctionDeclaration {
	return &genai.FunctionDeclaration{
		Name:        "GitAdd",
		Description: "Añade archivos al área de preparación (staging) de Git (ej. '.' para todos).",
		Parameters: &genai.Schema{
			Type:     genai.TypeObject,
			Required: []string{"pathspec"},
			Properties: map[string]*genai.Schema{
				"workingDir": {
					Type:        genai.TypeString,
					Description: "Directorio de trabajo del repositorio Git (opcional, por defecto '.').",
				},
				"pathspec": {
					Type:        genai.TypeString,
					Description: "Ruta del archivo o patrón a añadir (ej. '.' o 'main.go').",
				},
			},
		},
	}
}

func gitCommitTool() *genai.FunctionDeclaration {
	return &genai.FunctionDeclaration{
		Name:        "GitCommit",
		Description: "Realiza un commit en Git con los archivos preparados y el mensaje especificado.",
		Parameters: &genai.Schema{
			Type:     genai.TypeObject,
			Required: []string{"message"},
			Properties: map[string]*genai.Schema{
				"workingDir": {
					Type:        genai.TypeString,
					Description: "Directorio de trabajo del repositorio Git (opcional, por defecto '.').",
				},
				"message": {
					Type:        genai.TypeString,
					Description: "Mensaje descriptivo del commit.",
				},
			},
		},
	}
}
