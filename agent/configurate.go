package agent

import (
	"flag"
)

func (a *Agent) configurate() error {
	modelGemini := flag.String("model", "gemini-3.5-flash-lite", "Gemini model to use")
	verbose := flag.Bool("verbose", false, "Show more information")

	flag.Parse()

	a.model = *modelGemini
	a.verbose = *verbose
	a.systemPrompt = "You are  coding agent assistant, to can create, read and delete files and folder, you are tools for this propose"
	return nil
}
