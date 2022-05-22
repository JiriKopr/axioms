package utils

import (
	"fmt"
	"os"
	"strings"
	"sync"
)

var lock = &sync.Mutex{}

type ArgsIndexes struct {
	Command  int
	Template int
	Target   int
	Tags     int
}

var instance *ArgsIndexes

func GetArgsIndexesInstance() *ArgsIndexes {
	if instance == nil {

		// Lock rutines
		lock.Lock()
		defer lock.Unlock()

		// Check again if some rutine didn't create instance in the mean time
		if instance == nil {
			// axioms list     template
			// axioms generate template target tags
			command := 1
			template := command + 1
			target := template + 1
			tags := target + 1

			instance = &ArgsIndexes{
				Command:  command,
				Template: template,
				Target:   target,
				Tags:     tags,
			}
		}
	}

	return instance
}

func CheckArgsLength(min int) {
	argsIndexes := GetArgsIndexesInstance()
	argsLength := len(os.Args[argsIndexes.Command+1:])

	if argsLength < min {
		fmt.Println("Not enough arguments")
		os.Exit(1)
	}
}

func HandleTemplateInput() (os.FileInfo, string) {
	argsIndexes := GetArgsIndexesInstance()
	templatePath := os.Args[argsIndexes.Template]

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
	argsIndexes := GetArgsIndexesInstance()
	targetPath := os.Args[argsIndexes.Target]

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
