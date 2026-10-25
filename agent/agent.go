package agent

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/jsierrab3991/basic-agent-google/agent/tools"
	"google.golang.org/genai"
)

type Agent struct {
	model        string
	verbose      bool
	systemPrompt string
	maxToolSteps int
	client       *genai.Client
	ctx          context.Context
	cancel       context.CancelFunc
}

func (a *Agent) Start() {
	err := a.configure()
	if err != nil {
		fmt.Printf("Error de inicialización del agente: %v\n", err)
		return
	}
	if a.cancel != nil {
		defer a.cancel()
	}
	a.runAgent()
}

func clearTerminal() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func (a *Agent) runAgent() {
	var messages []*genai.Content
	scanner := bufio.NewScanner(os.Stdin)
	defer func() {
		err := scanner.Err()
		if err != nil {
			log.Printf("Error close scanner %v", err)
		}

	}()

	colorCwd := "\033[33m"
	resetColor := "\033[0m"
	for {
		if a.ctx != nil && a.ctx.Err() != nil {
			fmt.Println("\nOperación interrumpida.")
			break
		}

		cwd, _ := os.Getwd()
		fmt.Printf("%s %s %s > ", colorCwd, cwd, resetColor)
		if !scanner.Scan() {
			fmt.Println("\nHasta luego")
			break
		}

		rawInput := scanner.Text()
		trimmedInput := strings.TrimSpace(rawInput)
		if trimmedInput == "" {
			continue
		}

		lowerInput := strings.ToLower(trimmedInput)
		if lowerInput == "quit" || lowerInput == "exit" || lowerInput == "salir" {
			fmt.Println("Adiós")
			break
		}

		if lowerInput == "clear" || lowerInput == "cls" {
			clearTerminal()
			continue
		}

		// Preservar mayúsculas y minúsculas originales del usuario
		messages = append(messages, &genai.Content{
			Role: "user",
			Parts: []*genai.Part{
				genai.NewPartFromText(trimmedInput),
			},
		})

		stepCount := 0
		for {
			res, err := a.requestInput(messages)
			if err != nil {
				fmt.Printf("Error en la solicitud a %s: %v\n", a.model, err)
				break
			}

			if a.verbose {
				fmt.Printf("[Verbose] Modelo: %s\n", a.model)
				if res.UsageMetadata != nil {
					fmt.Printf("[Verbose] Tokens entrada: %d | Tokens salida: %d\n",
						res.UsageMetadata.PromptTokenCount, res.UsageMetadata.CandidatesTokenCount)
				}
			}

			// Mostrar información de uso y metadatos de tokens / uso general si está disponible en la respuesta
			if res.UsageMetadata != nil {
				fmt.Printf("📊 [Uso de Tokens] Prompt: %d | Candidatos (Respuesta): %d | Total: %d\n",
					res.UsageMetadata.PromptTokenCount,
					res.UsageMetadata.CandidatesTokenCount,
					res.UsageMetadata.TotalTokenCount,
				)
			}

			if len(res.Candidates) == 0 || res.Candidates[0].Content == nil {
				fmt.Println("No se recibió respuesta válida del modelo.")
				break
			}

			candidateContent := res.Candidates[0].Content
			functionCalls := res.FunctionCalls()

			respText := res.Text()
			if strings.TrimSpace(respText) != "" {
				fmt.Printf("🤖 Gemini: %s\n", respText)
			}

			messages = append(messages, &genai.Content{
				Role:  "model",
				Parts: candidateContent.Parts,
			})

			// Si no hubo llamadas a herramientas, fin del turno actual
			if len(functionCalls) == 0 {
				break
			}

			stepCount++
			if stepCount > a.maxToolSteps {
				fmt.Printf("⚠️ Límite de ejecución de herramientas alcanzado (%d pasos).\n", a.maxToolSteps)
				break
			}

			// Agrupar las respuestas de todas las herramientas llamadas en este turno en un único contenido
			var responseParts []*genai.Part
			for _, call := range functionCalls {
				fmt.Printf("🛠️ Herramienta: %s con argumentos: %v\n", call.Name, call.Args)
				toolResult, err := tools.ExecuteCall(call.Name, call.Args)
				var resultPayload map[string]any
				if err != nil {
					resultPayload = map[string]any{"error": err.Error()}
				} else {
					resultPayload = toolResult
				}
				responseParts = append(responseParts, genai.NewPartFromFunctionResponse(call.Name, resultPayload))
			}

			messages = append(messages, &genai.Content{
				Role:  "user",
				Parts: responseParts,
			})
		}
	}
}
