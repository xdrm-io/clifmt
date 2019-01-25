package clifmt

import (
	"fmt"
	"git.xdrm.io/go/clifmt/internal/color"
	colorTransform "git.xdrm.io/go/clifmt/internal/transform/color"
	mdTransform "git.xdrm.io/go/clifmt/internal/transform/markdown"
)

var theme = color.DefaultTheme()

// Sprintf returns a terminal-colorized output following the coloring format
func Sprintf(format string, a ...interface{}) (string, error) {
	// 1. Pre-process format with 'fmt'
	formatted := fmt.Sprintf(format, a...)

	// 2. Colorize
	colorized, err := colorTransform.Transform(formatted, theme)
	if err != nil {
		return "", err
	}

	// 3. Markdown format
	markdown, err := mdTransform.Transform(colorized)
	if err != nil {
		return "", err
	}

	// 3. return final output
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
