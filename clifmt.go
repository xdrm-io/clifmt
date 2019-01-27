package clifmt

import (
	"fmt"
	"git.xdrm.io/go/clifmt/internal/color"
	colorTransform "git.xdrm.io/go/clifmt/internal/transform/color"
	mdTransform "git.xdrm.io/go/clifmt/internal/transform/markdown"
	"strings"
)

var theme = color.DefaultTheme()

var (
	dollarToken        = `e4d097183ab04e49f25cb7b0956fb9eb25b90c0316a32cb5afcbcdd9a6692e8d2974919035789d5632b10d799db5b3e5bf8539592c904497f5c356f117ef37382`
	asteriskToken      = `253c3cd0a904d28abc3e601e3557d59ea69da2616079ceef4987d58d55c9820c83026be92a917ee19a298e613ea0b393cc70d4e55dc614a9afc6a020d8f08f37`
	underscoreToken    = `2b08e24b7833e90c74ed8e6c27b7b3cd5fe949e0f18b28af813d5f2df863d55f97b0ed7f8fbb26a152eda55ac073331ce11ac10702caca5b3ea4a29f722840b9`
	squareBracketToken = `51b06edd58f36003844941916cd3b313979fece55824d89ba02af052a229b2673aafffa541b703472c1a21d8e6a1bb3e844d236fb0e8bf5d62902b24042f4fb5`
)

// Sprintf returns a terminal-colorized output following the coloring format
func Sprintf(format string, a ...interface{}) (string, error) {
	// 1. Pre-process format with 'fmt'
	formatted := fmt.Sprintf(format, a...)

	// 2. Protect escaped characters with tokens
	formatted = strings.Replace(formatted, "\\$", dollarToken, -1)
	formatted = strings.Replace(formatted, "\\*", asteriskToken, -1)
	formatted = strings.Replace(formatted, "\\_", underscoreToken, -1)
	formatted = strings.Replace(formatted, "\\[", squareBracketToken, -1)

	// 3. Colorize
	colorized, err := colorTransform.Transform(formatted, theme)
	if err != nil {
		return "", err
	}

	// 4. Markdown format
	markdown, err := mdTransform.Transform(colorized)
	if err != nil {
		return "", err
	}

	// 5. Restore token-protected characters
	markdown = strings.Replace(markdown, dollarToken, "$", -1)
	markdown = strings.Replace(markdown, asteriskToken, "*", -1)
	markdown = strings.Replace(markdown, underscoreToken, "_", -1)
	markdown = strings.Replace(markdown, squareBracketToken, "[", -1)

	// 6. return final output
	return markdown, nil
}

// Printf prints a terminal-colorized output following the coloring format
func Printf(format string, a ...interface{}) error {
	s, err := Sprintf(format, a...)
	if err != nil {
		return err
	}

	fmt.Print(s)
	return nil
}
