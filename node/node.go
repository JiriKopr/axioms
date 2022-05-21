package node

import (
	"fmt"
	"os"
	"strings"
)

type Line struct {
	LineNumber int
	Content    string
}

type Node struct {
	Info os.FileInfo
}

type DirNode struct {
	Node    Node
	Subdirs []*DirNode
	Files   []*FileNode
}

type FileNode struct {
	Node  Node
	Lines []*Line
}

type DirInfo struct {
	Path       string
	Info       os.FileInfo
	ParentNode *DirNode
}

func (node FileNode) String() string {
	content := ""
	for _, line := range node.Lines {
		content += line.Content + " "
	}
	return node.Node.Info.Name() + "(" + content + ")"
}

func (dirNode DirNode) String() string {
	subnodesStrings := []string{}

	for _, subnode := range dirNode.Files {
		subnodesStrings = append(subnodesStrings, subnode.String())
	}

	for _, subnode := range dirNode.Subdirs {
		subnodesStrings = append(subnodesStrings, subnode.String())
	}

	return fmt.Sprintf("%s -> {%s}", dirNode.Node.Info.Name(), strings.Join(subnodesStrings, ", "))
}
