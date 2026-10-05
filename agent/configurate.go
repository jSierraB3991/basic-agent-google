package agent

import (
	"errors"
	"flag"
)

func (a *Agent) configurate() error {
	modelGemini := flag.String("model", "gemini-3.5-flash-lite", "Gemini model to use")
	verbose := flag.Bool("verbose", false, "Show more information")

	flag.Parse()
	argsPosicionales := flag.Args()

	if len(argsPosicionales) == 0 {
		return errors.New("I need a question.")
	}
	a.model = *modelGemini
	a.input = argsPosicionales[0]
	a.verbose = *verbose
	return nil
}
