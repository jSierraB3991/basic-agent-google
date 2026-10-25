package agent

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"strings"

	"google.golang.org/genai"
)

func (a *Agent) configurate() error {
	return a.configure()
}

func (a *Agent) configure() error {
	defaultModel := os.Getenv("GEMINI_MODEL")
	if strings.TrimSpace(defaultModel) == "" {
		defaultModel = "gemini-3.5-flash-lite"
	}

	defaultVerbose := false
	if v := os.Getenv("VERBOSE"); v == "true" || v == "1" {
		defaultVerbose = true
	}

	var modelFlag string
	var verboseFlag bool
	var maxStepsFlag int

	if flag.Lookup("model") == nil {
		flag.StringVar(&modelFlag, "model", defaultModel, "Gemini model to use")
		flag.BoolVar(&verboseFlag, "verbose", defaultVerbose, "Show more information")
		flag.IntVar(&maxStepsFlag, "max-steps", 15, "Max tool calls per turn")
		flag.Parse()
	} else {
		if f := flag.Lookup("model"); f != nil {
			modelFlag = f.Value.String()
		}
		if f := flag.Lookup("verbose"); f != nil {
			verboseFlag = f.Value.String() == "true"
		}
		maxStepsFlag = 15
	}

	a.model = modelFlag
	a.verbose = verboseFlag
	if maxStepsFlag <= 0 {
		maxStepsFlag = 15
	}
	a.maxToolSteps = maxStepsFlag

	a.systemPrompt = `You are an expert AI software engineering assistant. You help users inspect, write, debug, and test code.
You have access to tools to interact with the local filesystem and execute unit tests:
- GetFilesInfo: inspect directory structure and files.
- ReadFileContent: read the contents of files.
- CreateFile: write new files or overwrite existing files.
- DeleteFile: delete unnecessary files.
- CreateDir: create directories and required parent directories.
- DeleteDir: delete directories and their contents.
- RunTestInGo: execute Go unit tests (pass target package, test file, or './...').
- RunPythonTests: execute Python unit tests in a specified directory.

Guidelines:
1. Always explore and read relevant files before making modifications.
2. When creating or updating code, follow idiomatic conventions and best practices.
3. Run tests after making changes to verify correctness. If tests fail, analyze the error output and fix the issue.
4. Be concise and explain your actions clearly.`

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	a.ctx = ctx
	a.cancel = cancel

	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		return err
	}
	a.client = client

	return nil
}
