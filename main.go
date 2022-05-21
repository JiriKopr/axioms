package main

import (
	. "axioms/mappings"
	. "axioms/node"
	. "axioms/structure"
	. "axioms/utils"
	"fmt"
	"os"
)

// TODO: Get mappings from Args
// TODO: Add option to list all variable inputs
func main() {
	argsLength := len(os.Args[1:])

	if argsLength < 2 {
		fmt.Println("Not enough arguments")
		os.Exit(1)
	}

	selectedTemplate := os.Args[1]
	targetPath := os.Args[2]

	directoryInfo, error := os.Stat(fmt.Sprintf("./templates/%s", selectedTemplate))

	if os.IsNotExist(error) {
		fmt.Printf("Template %s not found\n", selectedTemplate)
		os.Exit(1)
	} else if error != nil {
		fmt.Printf("Error while reading %s template\n", selectedTemplate)
		os.Exit(1)
	}

	mappings := TagMappings{}
	mappings.Init()

	root := DirNode{
		Node: Node{
			Info: directoryInfo,
		},
		Files:   []*FileNode{},
		Subdirs: []*DirNode{},
	}

	WalkDir(DirInfo{Info: directoryInfo, Path: "./templates", ParentNode: &root})

	CreateStructure(TargetStructure{DirNode: &root, TargetPath: targetPath, Mappings: mappings})

	tagsSet := NewSet[string]()
	GetAllDirectoryTagsAndMods(&root, &tagsSet)
	fmt.Println(tagsSet.Values())
}
