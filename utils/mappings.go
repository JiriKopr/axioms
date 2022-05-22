package utils

import (
	"fmt"
	"os"
	"strings"
)

type TagMappings struct {
	MappedTags map[string]string
}

const DEFAULT_TAG = "axiom"

func (mappings *TagMappings) Init() {
    argsIndexes := GetArgsIndexesInstance()
	arguments := os.Args[argsIndexes.Tags:]

	mappings.MappedTags = map[string]string{}

	for _, argument := range arguments {
		assignment := strings.Split(argument, "=")

		if len(assignment) == 1 {
			assignment = []string{DEFAULT_TAG, assignment[0]}
		}

		key := assignment[0]
		value := assignment[1]

		existingTagMapping, exists := mappings.MappedTags[key]

		if exists {
			fmt.Printf("Duplicate mapping for '%s', tag already defined as %s", key, existingTagMapping)
			os.Exit(1)
		}

		mappings.MappedTags[key] = value
	}
}
