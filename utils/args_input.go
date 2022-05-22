package utils

import (
	"fmt"
	"os"
	"strings"
)

func CheckArgsLength(min int) {
	argsLength := len(os.Args[2:])

	if argsLength < min {
		fmt.Println("Not enough arguments")
		os.Exit(1)
	}
}

func HandleTemplateInput() (os.FileInfo, string) {
	templatePath := os.Args[2]

	directoryInfo, templateError := os.Stat(templatePath)

	dividedPath := strings.Split(templatePath, "/")
	selectedTemplate := dividedPath[:len(dividedPath)-1]

	if os.IsNotExist(templateError) {
		fmt.Printf("Template %s not found\n", selectedTemplate)
		os.Exit(1)
	} else if templateError != nil {
		fmt.Printf("Error while reading %s template\n", selectedTemplate)
		os.Exit(1)
	}

	return directoryInfo, templatePath
}

func HandleTargetInput() string {
	targetPath := os.Args[3]

	_, targetError := os.Stat(targetPath)
	if os.IsNotExist(targetError) {
		fmt.Printf("Target path %s does not exist\n", targetPath)
		os.Exit(1)
	} else if targetError != nil {
		fmt.Printf("Error while locating target path %s\n", targetPath)
		os.Exit(1)
	}

	return targetPath
}
