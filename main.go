package main

import (
	"fmt"
	"log"

	"github.com/jSierraB3991/http-cli/agent"
	"github.com/joho/godotenv"
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
