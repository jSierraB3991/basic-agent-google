package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/jsierrab3991/basic-agent-google/agent"
)

func main() {
	fmt.Println("Simple Code agent With gemini ")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("No load enviroment")
	}
	agent := agent.Agent{}
	agent.Start()
}
