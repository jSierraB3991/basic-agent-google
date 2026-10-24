package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/jsierrab3991/basic-agent-google/agent"
)

func main() {
	fmt.Println("Simple Code agent With gemini")

	// Carga variables desde el archivo .env si existe
	_ = godotenv.Load()

	if os.Getenv("GEMINI_API_KEY") == "" {
		fmt.Fprintln(os.Stderr, "Advertencia: GEMINI_API_KEY no se encuentra configurada en el entorno ni en el archivo .env")
	}

	agent := agent.Agent{}
	agent.Start()
}
