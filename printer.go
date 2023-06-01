package clifmt

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/xdrm-io/clifmt/internal/color"
	"github.com/xdrm-io/clifmt/internal/syntax"
	"github.com/xdrm-io/clifmt/internal/transform"
)

// Support can be set to false when the terminal does not support vt100 colors
// and formatting
var Support = true

// Theme is used to determine colors from their names ; feel free to replace it
// with yours
var Theme = color.DefaultTheme()

var (
	dollarToken        = `e4d097183ab04e49f25cb7b0956fb9eb25b90c0316a32cb5afcbcdd9a669`
	asteriskToken      = `253c3cd0a904d28abc3e601e3557d59ea69da2616079ceef4987d58d55c9`
	underscoreToken    = `2b08e24b7833e90c74ed8e6c27b7b3cd5fe949e0f18b28af813d5f2df863`
	squareBracketToken = `51b06edd58f36003844941916cd3b313979fece55824d89ba02af052a229`
)

// print formats a markdown-like syntax to the terminal-formatted output
func print(in string) (string, error) {
	formatted := in

	// protect escaped characters with tokens
	formatted = strings.Replace(formatted, "\\$", dollarToken, -1)
	formatted = strings.Replace(formatted, "\\*", asteriskToken, -1)
	formatted = strings.Replace(formatted, "\\_", underscoreToken, -1)
	formatted = strings.Replace(formatted, "\\[", squareBracketToken, -1)

	// create transformation registry
	reg := transform.Registry{Transformers: make([]transform.Transformer, 0, 10)}
	reg.Transformers = append(reg.Transformers, syntax.Color(Theme))
	reg.Transformers = append(reg.Transformers, syntax.Bold{})
	reg.Transformers = append(reg.Transformers, syntax.Italic{})
	reg.Transformers = append(reg.Transformers, syntax.Underline{})
	reg.Transformers = append(reg.Transformers, syntax.Hyperlink{})

	transformed, err := reg.Transform(formatted)
	if err != nil {
		return "", err
	}

	// restore token-protected characters
	transformed = strings.Replace(transformed, dollarToken, "$", -1)
	transformed = strings.Replace(transformed, asteriskToken, "*", -1)
	transformed = strings.Replace(transformed, underscoreToken, "_", -1)
	transformed = strings.Replace(transformed, squareBracketToken, "[", -1)

	if !Support {
		transformed = Escape(transformed)
	}

	return transformed, nil
}

// Sprintf mimics fmt.Sprintf with color formatting
func Sprintf(format string, a ...interface{}) string {
	// pre-process format with 'fmt'
	preformatted := fmt.Sprintf(format, a...)
	if strings.Contains(preformatted, "%!") { // fmt error
		return preformatted
	}

	formatted, err := print(preformatted)
	if err != nil {
		return fmt.Sprintf("%%!{clifmt: %s}", err)
	}
	return formatted
}

// Fprintf mimics fmt.Sprintf with color formatting
func Fprintf(w io.Writer, format string, a ...interface{}) (int, error) {
	return fmt.Fprint(w, Sprintf(format, a...))
}

// Printf mimics fmt.Sprintf with color formatting
func Printf(format string, a ...interface{}) (int, error) {
	return Fprintf(os.Stdout, format, a...)
}

var escapeSequence = regexp.MustCompile(`\x1b(?:\[|\]8;;)[^\\m]+[\\m]`)

// Escape remove escape sequences when terminal formatting/colors is not wanted
func Escape(format string) string {
	return escapeSequence.ReplaceAllString(format, "")
}
