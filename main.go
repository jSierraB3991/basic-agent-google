package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"google.golang.org/genai"
)

func run_agent(ctx context.Context) {
	modelGemini := flag.String("model", "gemini-3.5-flash-lite", "Gemini model to use")

	flag.Parse()
	argsPosicionales := flag.Args()

	if len(argsPosicionales) == 0 {
		log.Fatal("I nedd a prompt.")
	}

	userInput := argsPosicionales[0]
	messages := []*genai.Content{
		{
			Role: "user",
			Parts: []*genai.Part{
				genai.NewPartFromText(userInput),
			},
		},
	}

	fmt.Printf("model of gemini: %s\n", *modelGemini)
	fmt.Printf("prompt: %s\n", userInput)
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

	if res == nil || res.UsageMetadata == nil {
		log.Println("response or metadata is null")
		return
	}
	fmt.Printf("Prompt Token %d\n", res.UsageMetadata.PromptTokenCount)
	fmt.Printf("Response Token %d\n", res.UsageMetadata.CandidatesTokenCount)
	fmt.Println(res.Text())
}

func main() {
	fmt.Println("Herramienta para generar archivos para una API en net/http")
	ctx := context.Background()
	godotenv.Load()
	run_agent(ctx)
}
