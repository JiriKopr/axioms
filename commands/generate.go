package commands

import (
	. "axioms/node"
	. "axioms/structure"
	. "axioms/utils"
)

func GenerateStructureFromTemplateCommand() {
	CheckArgsLength(3)

	templateDirInfo, templatePath := HandleTemplateInput()
	targetPath := HandleTargetInput()

	mappings := TagMappings{}
	mappings.Init()

	root := DirNode{
		Node: Node{
			Info: templateDirInfo,
		},
		Files:   []*FileNode{},
		Subdirs: []*DirNode{},
	}

	WalkDir(DirInfo{Info: templateDirInfo, Path: templatePath, ParentNode: &root})

	CreateStructure(TargetStructure{DirNode: &root, TargetPath: targetPath, Mappings: mappings})
}
