package main

import (
	. "axioms/commands"
	"fmt"
	"os"
)

func main() {
	argsLength := len(os.Args[1:])

	if argsLength < 1 {
		fmt.Println("Command not provided")
		os.Exit(1)
	}

	commands := map[string]func(){
		"list":     ListAllAvailableTagsCommand,
		"generate": GenerateStructureFromTemplateCommand,
	}

	inputtedCommand := os.Args[1]
	selectedCommand, exists := commands[inputtedCommand]

	if !exists {
		fmt.Printf("'%s' command does not exist", inputtedCommand)
		os.Exit(1)
	}

	selectedCommand()
}
