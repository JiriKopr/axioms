package structure

import (
	. "axioms/mappings"
	. "axioms/node"
	. "axioms/utils"
	"fmt"
	"io/ioutil"
	"os"
)



func WalkDir(dirInfo DirInfo) {
	currentPath := dirInfo.Path
	parentNode := dirInfo.ParentNode

	items, _ := ioutil.ReadDir(currentPath)

	for _, item := range items {
		newNode := Node{
			Info: item,
		}

		itemPath := fmt.Sprintf("%s/%s", currentPath, item.Name())

		if item.IsDir() {
			dirNode := DirNode{
				Node: newNode,
			}

			parentNode.Subdirs = append(parentNode.Subdirs, &dirNode)

			WalkDir(DirInfo{Path: itemPath, Info: item, ParentNode: &dirNode})
			continue
		}

		fileNode := FileNode{
			Node:  newNode,
			Lines: []*Line{},
		}

		parentNode.Files = append(parentNode.Files, &fileNode)

		lines := loadFileContent(itemPath)

		fileNode.Lines = append(fileNode.Lines, lines[:]...)
	}
}

type TargetStructure struct {
	DirNode    *DirNode
	TargetPath string
	Mappings   TagMappings
}

func CreateStructure(targetStructure TargetStructure) {
	dirNode := targetStructure.DirNode
	targetPath := targetStructure.TargetPath
	mappings := targetStructure.Mappings

	for _, fileNode := range dirNode.Files {
		transformedFileName := TransformNodeName(&fileNode.Node, mappings)
		targetFilePath := fmt.Sprintf("%s/%s", targetPath, transformedFileName)

		file, err := os.Create(targetFilePath)

		if err != nil {
			fmt.Printf("Error while creating %s\n", targetFilePath)
			os.Exit(1)
		}

		writeLines(file, fileNode, mappings)
	}

	for _, subDir := range dirNode.Subdirs {
		transformedDirName := TransformNodeName(&subDir.Node, mappings)
		dirPath := fmt.Sprintf("%s/%s", targetPath, transformedDirName)

		os.Mkdir(dirPath, os.FileMode(0731))

		targetStructure.DirNode = subDir
		targetStructure.TargetPath = dirPath
		CreateStructure(targetStructure)
	}
}

func GetAllDirectoryTags(dirNode *DirNode, tagsSet *Set[string]) []string {
	for _, fileNode := range dirNode.Files {
		fileTags := getAllTagsFromLines(fileNode.Lines)
		fileNameTags := getTagsFromString(fileNode.Node.Info.Name())

		tagsSet.AddAll(fileTags)
		tagsSet.AddAll(fileNameTags)
	}

	for _, subDir := range dirNode.Subdirs {
		dirNameTags := getTagsFromString(dirNode.Node.Info.Name())

		tagsSet.AddAll(dirNameTags)

		GetAllDirectoryTags(subDir, tagsSet)
	}

	return tagsSet.Values()
}
