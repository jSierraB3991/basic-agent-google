package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"google.golang.org/genai"
)

func run_agent() {
	modelGemini := flag.String("model", "gemini-3.5-flash-lite", "Gemini model to use")
	verbose := flag.Bool("verbose", false, "Show more information")

	flag.Parse()
	argsPosicionales := flag.Args()

	if len(argsPosicionales) == 0 {
		log.Fatal("I nedd a prompt.")
	}

	ctx := context.Background()
	userInput := argsPosicionales[0]
	messages := []*genai.Content{
		{
			Role: "user",
			Parts: []*genai.Part{
				genai.NewPartFromText(userInput),
			},
		},
	}

	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatalf("Error to create client: %v", err)
	}

	res, err := client.Models.GenerateContent(
		ctx,
		*modelGemini,
		messages,
		nil,
	)
	if err != nil {
		log.Fatalf("Error generate the content: %v", err)
	}

	if *verbose {
		fmt.Printf("model of gemini: %s\n", *modelGemini)
		fmt.Printf("user input: %s\n", userInput)
		fmt.Printf("Prompt Token %d\n", res.UsageMetadata.PromptTokenCount)
		fmt.Printf("Response Token %d\n", res.UsageMetadata.CandidatesTokenCount)
	}
	fmt.Printf("Response: %s\n", res.Text())
}

func main() {
	fmt.Println("Herramienta para generar archivos para una API en net/http")
	godotenv.Load()
	run_agent()
}
