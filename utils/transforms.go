package utils

import (
	. "axioms/mappings"
	. "axioms/node"
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"
)

func transformCapitalize(value []string) []string {
	return Map(value, func(part string) string {
		runes := []rune(part)
		runes[0] = unicode.ToUpper(runes[0])

		return string(runes)
	})
}

func transformLower(value []string) []string {
	return Map(value, func(part string) string {
		return strings.ToLower(part)
	})
}

func transformPascal(value []string) string {
	mappedValue := transformCapitalize(value)

	return strings.Join(mappedValue, "")
}

func transformCamel(value []string) string {
	mappedValue := []string{value[0]}

	capitalizedPart := transformCapitalize(value[1:])

	mappedValue = append(mappedValue, capitalizedPart[:]...)

	return strings.Join(mappedValue, "")
}

func transformSnake(value []string) string {
	mappedValue := transformLower(value)

	return strings.Join(mappedValue, "_")
}

func transformKebab(value []string) string {
	mappedValue := transformLower(value)

	return strings.Join(mappedValue, "-")
}

func transformTag(tag string, modificator string) string {
	if modificator == "" {
		return tag
	}

	dividedValue := strings.Split(tag, "_")

	switch modificator {
	case "pascalcase":
		return transformPascal(dividedValue)
	case "camelcase":
		return transformCamel(dividedValue)
	case "snakecase":
		return transformSnake(dividedValue)
	case "kebabcase":
		return transformKebab(dividedValue)
	default:
		return strings.Join(dividedValue, "")
	}
}

func GetTagsAndModsRegexp() *regexp.Regexp {
	return regexp.MustCompile(`\[\s*(?:[a-zA-Z0-9]{1,})\s*(?:\|\s*[a-zA-Z]{1,}\s*)?\s*]`)
}

func TransformString(value string, mappings TagMappings) string {
	lineReg := GetTagsAndModsRegexp()

	transformedValue := lineReg.ReplaceAllStringFunc(value, func(match string) string {

		tagName, modificator := getTagAndMod(match)

		matchedValue, tagMapExists := mappings.MappedTags[tagName]

		if !tagMapExists {
			fmt.Printf("Tag mapping not found: %s\n", tagName)
			os.Exit(1)
		}

		return transformTag(matchedValue, modificator)
	})

	return transformedValue
}

func TransformNodeName(node *Node, mappings TagMappings) string {
	return TransformString(node.Info.Name(), mappings)
}

func getTagAndMod(value string) (string, string) {
	replaceReg := regexp.MustCompile(`\w+(?:\s*\|\s*\w+)?`)

	nameWithMods := replaceReg.FindString(value)

	nameAndMods := strings.Split(nameWithMods, "|")

	if len(nameAndMods) == 0 {
		fmt.Printf("Error while parsing tag %s\n", value)
		os.Exit(1)
	}

	tagName := ""
	if len(nameAndMods) > 0 {
		tagName = strings.TrimSpace(nameAndMods[0])
	}

	modificator := ""
	if len(nameAndMods) > 1 {
		modificator = strings.TrimSpace(nameAndMods[1])
	}

	return tagName, modificator
}
