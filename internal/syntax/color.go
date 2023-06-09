package syntax

import (
	"fmt"
	"regexp"

	"github.com/xdrm-io/clifmt/internal/color"
)

// Color implements transform.Transformer for the color syntax. It wraps its
// color theme
type Color color.Theme

// Regex implements transform.Transformer
func (Color) Regex() *regexp.Regexp {
	return regexp.MustCompile(`(?m)\${([^$]+)}\(((?:[a-z]+|#(?:[0-9a-f]{3}|[0-9a-f]{6})))?(?:\:((?:[a-z]+|#(?:[0-9a-f]{3}|[0-9a-f]{6}))))?\)`)
}

// Transform implements transform.Transformer
func (c Color) Transform(args ...string) (string, error) {
	// ignore no arg or empty
	if len(args) < 3 {
		return "", fmt.Errorf("invalid format")
	}

	// extract colors
	var (
		fg *color.T = nil
		bg *color.T = nil
	)

	if len(args[1]) > 0 {
		tmp, err := color.Parse(color.Theme(c), args[1])
		if err != nil {
			return "", err
		}
		fg = &tmp
	}
	if len(args[2]) > 0 {
		tmp, err := color.Parse(color.Theme(c), args[2])
		if err != nil {
			return "", err
		}
		bg = &tmp
	}

	return colorize(args[0], fg, bg), nil

}

// colorize returns the terminal-formatted @text colorized with the @fg and @bg colors
func colorize(t string, fg *color.T, bg *color.T) string {
	// no coloring
	if fg == nil && bg == nil {
		return t
	}

	// only foreground
	if bg == nil {
		return fmt.Sprintf("\x1b[38;2;%d;%d;%dm%s\x1b[0m", fg.Red(), fg.Green(), fg.Blue(), t)
	}
	// only background
	if fg == nil {
		return fmt.Sprintf("\x1b[48;2;%d;%d;%dm%s\x1b[0m", bg.Red(), bg.Green(), bg.Blue(), t)
	}

	// both colors
	return fmt.Sprintf("\x1b[38;2;%d;%d;%d;48;2;%d;%d;%dm%s\x1b[0m", fg.Red(), fg.Green(), fg.Blue(), bg.Red(), bg.Green(), bg.Blue(), t)
}
