package agent

import (
	"context"
	"log"

	"google.golang.org/genai"
)

func (a *Agent) requestInput(messages []*genai.Content) (*genai.GenerateContentResponse, error) {

	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatalf("Error to create client: %v", err)
	}
	return client.Models.GenerateContent(
		ctx,
		a.model,
		messages,
		nil,
	)
}
