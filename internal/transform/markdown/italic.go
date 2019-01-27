package markdown

import (
	"fmt"
	"regexp"
	"strings"
)

var italicRe = regexp.MustCompile(`(?m)\*([^\*]+)\*`)

// italic returns the terminal-formatted italic text @t
func italic(t string) string {
	return fmt.Sprintf("\x1b[3m%s\x1b[23m", strings.Replace(t, "\x1b[0m", "\x1b[0m\x1b[3m", -1))
}

// italicTransform the @input text using markdown-like syntax :
// - "normal *italic* normal"
func italicTransform(input string) (string, error) {
	output := ""
	cursor := int(0)

	// 1. Replace for each match
	for _, match := range italicRe.FindAllStringSubmatchIndex(input, -1) {

		// (1) add gap between input start OR previous match
		output += input[cursor:match[0]]
		cursor = match[1]

		// (2) extract features
		text := ""

		if match[3]-match[2] > 0 {
			text = input[match[2]:match[3]]
		}

		// (3) replace text with bold text
		output += italic(text)
	}

	// 2. Add end of input
	if cursor < len(input) {
		output += input[cursor:]
	}

	// 3. print final output
	return output, nil

}
