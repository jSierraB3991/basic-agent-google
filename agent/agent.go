package agent

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"google.golang.org/genai"
)

type Agent struct {
	model   string
	input   string
	verbose bool
}

func (a *Agent) Start() {
	err := a.configurate()
	if err != nil {
		log.Println(err)
		return
	}
	a.runAgent(a.input)
}

func (a *Agent) runAgent(input string) {
	for {
		cleanInput := strings.ToLower(strings.TrimSpace(input))
		if cleanInput == "" {
			scanner := bufio.NewScanner(os.Stdin)
			fmt.Print(">: ")
			if scanner.Scan() {
				cleanInput = scanner.Text()
			}
		} else {
			fmt.Printf(">: %s \n", cleanInput)
		}

		if cleanInput == "quit" || cleanInput == "exit" || cleanInput == "salir" {
			fmt.Println("Adiós")
			break
		}

		messages := []*genai.Content{
			{
				Role: "user",
				Parts: []*genai.Part{
					genai.NewPartFromText(cleanInput),
				},
			},
		}

		res, err := a.requestInput(messages)
		if err != nil {
			log.Fatalf("Error request to %s %s", a.model, err)
		}

		if a.verbose {
			fmt.Printf("model of gemini: %s\n", a.model)
			fmt.Printf("user input: %s\n", a.input)
			fmt.Printf("Prompt Token %d\n", res.UsageMetadata.PromptTokenCount)
			fmt.Printf("Response Token %d\n", res.UsageMetadata.CandidatesTokenCount)
		}
		fmt.Printf("Response: %s\n", res.Text())
		input = ""
	}
}
