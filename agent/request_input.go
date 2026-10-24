package agent

import (
	"context"
	"fmt"

	"github.com/jsierrab3991/basic-agent-google/agent/tools"
	"google.golang.org/genai"
)

func (a *Agent) requestInput(messages []*genai.Content) (*genai.GenerateContentResponse, error) {
	if a.client == nil {
		return nil, fmt.Errorf("el cliente de Gemini no está inicializado")
	}

	ctx := a.ctx
	if ctx == nil {
		ctx = context.Background()
	}

	return a.client.Models.GenerateContent(
		ctx,
		a.model,
		messages,
		tools.GetConfigGenerator(a.systemPrompt),
	)
}
