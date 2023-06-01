package clifmt

import (
	"fmt"
	"strings"

	"github.com/xdrm-io/clifmt/internal/color"
	tbold "github.com/xdrm-io/clifmt/internal/syntax/bold"
	tcolor "github.com/xdrm-io/clifmt/internal/syntax/color"
	thyperlink "github.com/xdrm-io/clifmt/internal/syntax/hyperlink"
	titalic "github.com/xdrm-io/clifmt/internal/syntax/italic"
	tunderline "github.com/xdrm-io/clifmt/internal/syntax/underline"
	"github.com/xdrm-io/clifmt/internal/transform"
)

// ErrInvalidFormat raised when the format is invalid
var ErrInvalidFormat = fmt.Errorf("invalid format")

var theme = color.DefaultTheme()

var (
	dollarToken        = `e4d097183ab04e49f25cb7b0956fb9eb25b90c0316a32cb5afcbcdd9a669`
	asteriskToken      = `253c3cd0a904d28abc3e601e3557d59ea69da2616079ceef4987d58d55c9`
	underscoreToken    = `2b08e24b7833e90c74ed8e6c27b7b3cd5fe949e0f18b28af813d5f2df863`
	squareBracketToken = `51b06edd58f36003844941916cd3b313979fece55824d89ba02af052a229`
)

// Sprintf returns a terminal-colorized output following the coloring format
func Sprintf(format string, a ...interface{}) (string, error) {
	// Pre-process format with 'fmt'
	formatted := fmt.Sprintf(format, a...)
	if strings.Contains(formatted, "%!") { // error
		return "", ErrInvalidFormat
	}

	// protect escaped characters with tokens
	formatted = strings.Replace(formatted, "\\$", dollarToken, -1)
	formatted = strings.Replace(formatted, "\\*", asteriskToken, -1)
	formatted = strings.Replace(formatted, "\\_", underscoreToken, -1)
	formatted = strings.Replace(formatted, "\\[", squareBracketToken, -1)

	// create transformation registry
	reg := transform.Registry{Transformers: make([]transform.Transformer, 0, 10)}
	reg.Transformers = append(reg.Transformers, tcolor.Export)
	reg.Transformers = append(reg.Transformers, tbold.Export)
	reg.Transformers = append(reg.Transformers, titalic.Export)
	reg.Transformers = append(reg.Transformers, tunderline.Export)
	reg.Transformers = append(reg.Transformers, thyperlink.Export)

	transformed, err := reg.Transform(formatted)
	if err != nil {
		return "", err
	}

	// restore token-protected characters
	transformed = strings.Replace(transformed, dollarToken, "$", -1)
	transformed = strings.Replace(transformed, asteriskToken, "*", -1)
	transformed = strings.Replace(transformed, underscoreToken, "_", -1)
	transformed = strings.Replace(transformed, squareBracketToken, "[", -1)

	return transformed, nil
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
