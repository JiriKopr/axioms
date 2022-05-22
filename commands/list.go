package commands

import (
	. "axioms/node"
	. "axioms/structure"
	. "axioms/utils"
	"fmt"
)

func ListAllAvailableTagsCommand() {
	CheckArgsLength(1)

	templateDirInfo, templatePath := HandleTemplateInput()

	root := DirNode{
		Node: Node{
			Info: templateDirInfo,
		},
		Files:   []*FileNode{},
		Subdirs: []*DirNode{},
	}

	WalkDir(DirInfo{Info: templateDirInfo, Path: templatePath, ParentNode: &root})

	tagsSet := NewSet[string]()
	GetAllDirectoryTags(&root, &tagsSet)

	fmt.Println("Available tags:")
	for _, tag := range tagsSet.Values() {
		fmt.Println(tag)
	}
}
