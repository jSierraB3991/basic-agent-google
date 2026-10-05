package agent

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jsierrab3991/basic-agent-google/agent/tools"
	"google.golang.org/genai"
)

type Agent struct {
	model        string
	verbose      bool
	systemPrompt string
}

func (a *Agent) Start() {
	err := a.configurate()
	if err != nil {
		log.Println(err)
		return
	}
	a.runAgent()
}

func (a *Agent) runAgent() {
	var messages []*genai.Content
	var input string
	for {
		var cleanInput string
		if input == "" {
			scanner := bufio.NewScanner(os.Stdin)
			fmt.Print(">: ")
			if scanner.Scan() {
				input = scanner.Text()
			}

			cleanInput = strings.ToLower(strings.TrimSpace(input))
			if cleanInput == "" {
				continue
			}
			if cleanInput == "quit" || cleanInput == "exit" || cleanInput == "salir" {
				fmt.Println("Adiós")
				break
			}
			messages = append(messages, &genai.Content{
				Role: "user",
				Parts: []*genai.Part{
					genai.NewPartFromText(cleanInput),
				},
			})
		}
		res, err := a.requestInput(messages)
		if err != nil {
			fmt.Printf("Error request to %s %s\n", a.model, err)
			input = ""
			continue
		}

		functionCalls := res.FunctionCalls()
		if len(functionCalls) == 0 {
			if a.verbose {
				fmt.Printf("model of gemini: %s\n", a.model)
				fmt.Printf("user input: %s\n", cleanInput)
				fmt.Printf("Prompt Token %d\n", res.UsageMetadata.PromptTokenCount)
				fmt.Printf("Response Token %d\n", res.UsageMetadata.CandidatesTokenCount)
			}
			fmt.Printf("Response: %s\n", res.Text())
			input = ""
		}

		messages = append(messages, &genai.Content{
			Role:  "model",
			Parts: res.Candidates[0].Content.Parts, // The model's original function call request
		})
		for _, call := range functionCalls {
			fmt.Printf("🤖 Gemini requested function call: %s with args: %v\n", call.Name, call.Args)
			resultExecuteTool, err := tools.ExecuteCall(call.Name, call.Args)
			var result map[string]any
			if err != nil {
				result = map[string]any{"error": err.Error()}
			} else {
				result = resultExecuteTool
			}

			messages = append(messages, &genai.Content{
				Role: "user",
				Parts: []*genai.Part{
					genai.NewPartFromFunctionResponse(call.Name, result),
				},
			})
		}

	}
}
