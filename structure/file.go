package structure

import (
	. "axioms/node"
	. "axioms/utils"
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func loadFileContent(path string) []*Line {
	content, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Printf("ERROR reading file %s -> %s\n", path, err.Error())
		os.Exit(1)
	}

	stringContent := string(content)

	lineStrings := strings.Split(stringContent, "\n")

	lines := []*Line{}
	for lineNumber, lineContent := range lineStrings {
		line := Line{Content: lineContent, LineNumber: lineNumber}

		lines = append(lines, &line)
	}

	return lines
}

func writeLines(file *os.File, fileNode *FileNode, mappings TagMappings) {
	writer := bufio.NewWriter(file)

	for _, line := range fileNode.Lines {
		transformedContent := TransformString(line.Content, mappings)

		writer.WriteString(transformedContent + "\n")
	}

	writer.Flush()
}

func getTagsFromString(value string) []string {

	lineReg := GetTagsAndModsRegexp()

	foundTagsAndMods := lineReg.FindAllString(value, -1)

	return Map(foundTagsAndMods, func(tagAndMod string) string {
		tag, _ := GetTagAndMod(tagAndMod)

		return tag
	})
}

func getAllTagsFromLines(lines []*Line) []string {
	tagsAndModsSet := NewSet[string]()

	for _, line := range lines {
		lineTags := getTagsFromString(line.Content)

		for _, tag := range lineTags {
			tagsAndModsSet.Add(tag)
		}
	}

	return tagsAndModsSet.Values()
}
