package markdown

import (
	"fmt"
	"regexp"
	"strings"
)

var boldRe = regexp.MustCompile(`(?m)\*\*((?:[^\*]+\*?)+)\*\*`)

// boldify returns the terminal-formatted bold text @t
func boldify(t string) string {
	return fmt.Sprintf("\x1b[1m%s\x1b[22m", strings.Replace(t, "\x1b[0m", "\x1b[0m\x1b[1m", -1))
}

// boldTransform the @input text using markdown-like syntax :
// - "normal **bold** normal"
func boldTransform(input string) (string, error) {
	output := ""
	cursor := int(0)

	// 1. Replace for each match
	for _, match := range boldRe.FindAllStringSubmatchIndex(input, -1) {

		// (1) add gap between input start OR previous match
		output += input[cursor:match[0]]
		cursor = match[1]

		// (2) extract features
		text := ""

		if match[3]-match[2] > 0 {
			text = input[match[2]:match[3]]
		}

		// (3) replace text with bold text
		output += boldify(text)
	}

	// 2. Add end of input
	if cursor < len(input)-1 {
		output += input[cursor:]
	}

	// 3. print final output
	return output, nil

}
