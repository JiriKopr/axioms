package main

import (
	. "axioms/commands"
	. "axioms/utils"
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

	argsIndexes := GetArgsIndexesInstance()
	inputtedCommand := os.Args[argsIndexes.Command]
	selectedCommand, exists := commands[inputtedCommand]

	if !exists {
		fmt.Printf("'%s' command does not exist", inputtedCommand)
		os.Exit(1)
	}

	selectedCommand()
}
